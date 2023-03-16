package enthistory

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"regexp"
	"strings"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
)

func extractUpdatedByKey(val any) string {
	updatedBy, ok := val.(UpdatedBy)
	if !ok {
		return ""
	}
	return updatedBy.key
}

func extractUpdatedByValueType(val any) string {
	updatedBy, ok := val.(UpdatedBy)
	if !ok {
		return ""
	}

	switch updatedBy.valueType {
	case ValueTypeInt:
		return "int"
	case ValueTypeString:
		return "string"
	default:
		return ""
	}
}

func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"extractUpdatedByKey":       extractUpdatedByKey,
		"extractUpdatedByValueType": extractUpdatedByValueType,
	})
	return gen.MustParse(t.ParseFS(_templates, path))
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

func getUpdatedByField(updatedByValueType string) (*load.Field, error) {
	if updatedByValueType == "String" {
		return load.NewField(field.String("updated_by").Optional().Nillable().Descriptor())
	}
	if updatedByValueType == "Int" {
		return load.NewField(field.Int("updated_by").Optional().Nillable().Descriptor())
	}
	return nil, errors.New("improper value type must be 'String' or 'Int'")
}

func getHistoryAnnotations(schema *load.Schema) Annotations {
	annotations := Annotations{}
	if historyAnnotations, ok := schema.Annotations["History"].(map[string]any); ok {
		if exclude, ok := historyAnnotations["exclude"].(bool); ok {
			annotations.Exclude = exclude
		}
		if isHistory, ok := historyAnnotations["isHistory"].(bool); ok {
			annotations.IsHistory = isHistory
		}
	}
	return annotations
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
