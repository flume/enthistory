package enthistory

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

func parseTemplate(name, path string) *gen.Template {
	return gen.MustParse(gen.NewTemplate(name).ParseFS(_templates, path))
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func copyRef[T any](ref *T) *T {
	if ref == nil {
		return nil
	}
	val := *ref
	return &val
}

func newFieldsFromField(fields []*load.Field) []*load.Field {
	newFields := make([]*load.Field, len(fields))
	i := 4
	for j, field := range fields {
		newField := load.Field{
			Name:          field.Name,
			Info:          copyRef(field.Info),
			Tag:           field.Tag,
			Size:          copyRef(field.Size),
			Enums:         field.Enums,
			Unique:        field.Unique,
			Nillable:      field.Nillable,
			Optional:      field.Optional,
			Default:       field.Default,
			DefaultValue:  field.DefaultValue,
			DefaultKind:   field.DefaultKind,
			UpdateDefault: field.UpdateDefault,
			Immutable:     field.Immutable,
			Validators:    field.Validators,
			StorageKey:    field.StorageKey,
			Position:      copyRef(field.Position),
			Sensitive:     field.Sensitive,
			SchemaType:    field.SchemaType,
			Annotations:   field.Annotations,
			Comment:       field.Comment,
		}
		if !field.Position.MixedIn {
			newField.Position = &load.Position{
				Index:      i,
				MixedIn:    false,
				MixinIndex: 0,
			}
			i += 1
		}
		newFields[j] = &newField
	}
	return newFields
}

func getHistorySchemaPath(schema *load.Schema) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%v/schema/%v.go", dir, fmt.Sprintf("%s_history", strings.ToLower(schema.Name)))
	return path, nil
}

func removeOldGenerated(schemas []*load.Schema) error {
	for _, schema := range schemas {
		path, err := getHistorySchemaPath(schema)
		if err != nil {
			return err
		}

		err = os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadHistorySchema() (*load.Schema, error) {
	bytes, err := load.MarshalSchema(history{})
	if err != nil {
		return nil, err
	}

	historySchema, err := load.UnmarshalSchema(bytes)
	if err != nil {
		return nil, err
	}
	return historySchema, nil
}

func mergeSchemaAndHistorySchema(historySchema, schema *load.Schema) string {
	historySchema.Name = fmt.Sprintf("%vHistory", schema.Name)

	historySchema.Fields = append(historySchema.Fields, newFieldsFromField(schema.Fields)...)

	tableName := fmt.Sprintf("%v", ToSnakeCase(historySchema.Name))

	if entSqlMap, ok := schema.Annotations["EntSQL"].(map[string]any); ok {
		if table, ok := entSqlMap["table"].(string); ok {
			tableName = fmt.Sprintf("%v_history", table)
		}
	}

	historySchema.Annotations = map[string]any{
		"EntSQL": map[string]any{
			"table": tableName,
		},
	}

	return tableName
}
