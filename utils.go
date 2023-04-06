package enthistory

import (
	"regexp"
	"strings"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/load"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
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

func createHistoryFields(schemaFields []*load.Field) []*load.Field {
	historyFields := make([]*load.Field, len(schemaFields))
	i := 4
	for j, field := range schemaFields {
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
		historyFields[j] = &newField
	}
	return historyFields
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
		return load.NewField(field.String("updated_by").Optional().Nillable().Immutable().Descriptor())
	}
	if updatedByValueType == "Int" {
		return load.NewField(field.Int("updated_by").Optional().Nillable().Immutable().Descriptor())
	}
	return nil, nil
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

func getSchemaTableName(schema *load.Schema) string {
	if entSqlMap, ok := schema.Annotations["EntSQL"].(map[string]any); ok {
		if table, ok := entSqlMap["table"].(string); ok {
			return table
		}
	}
	return toSnakeCase(schema.Name)
}
