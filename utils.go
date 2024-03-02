package enthistory

import (
	"errors"
	"maps"
	"reflect"
	"regexp"
	"strings"

	"entgo.io/contrib/entgql"

	"github.com/google/uuid"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/load"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

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

func loadHistorySchema(IdType *field.TypeInfo, entqlEnabled bool) (*load.Schema, error) {
	schema := history{}

	switch IdType.String() {
	case "int":
		f := field.Int("ref").Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		schema.ref = f
	case "string":
		f := field.String("ref").Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		schema.ref = f
	case "uuid.UUID":
		equal := IdType.RType.TypeEqual(reflect.TypeOf(uuid.UUID{}))
		if !equal {
			return nil, errors.New("unsupported uuid type")
		}
		f := field.UUID("ref", uuid.UUID{}).Immutable().Optional()
		if entqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		schema.ref = f
	default:
		return nil, errors.New("only id and string are supported id types right now")
	}

	bytes, err := load.MarshalSchema(schema)
	if err != nil {
		return nil, err
	}

	historySchema, err := load.UnmarshalSchema(bytes)
	if err != nil {
		return nil, err
	}

	return historySchema, nil
}

func deref[T any](t *T) T {
	var zero T
	if t == nil {
		return zero
	}
	return *t
}

func getUpdatedByField(updatedByValueType string, entgqlEnabled bool) (*load.Field, error) {
	if updatedByValueType == "String" {
		f := field.String("updated_by").Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return load.NewField(f.Descriptor())
	}
	if updatedByValueType == "Int" {
		f := field.Int("updated_by").Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return load.NewField(f.Descriptor())
	}
	if updatedByValueType == "UUID" {
		f := field.UUID("updated_by", uuid.UUID{}).Optional().Nillable().Immutable()
		if entgqlEnabled {
			f = f.Annotations(entgql.Annotation{Type: "ID"})
		}
		return load.NewField(f.Descriptor())
	}
	return nil, nil
}

func mergeAnnotations(dest, src map[string]any) map[string]any {
	merged := maps.Clone(src)
	for k, v := range dest {
		if _, ok := merged[k]; !ok {
			merged[k] = v
		} else {
			destMap, destOk := v.(map[string]any)
			srcMap, srcOk := merged[k].(map[string]any)
			if destOk && srcOk {
				merged[k] = mergeAnnotations(destMap, srcMap)
			}
		}
	}
	return merged
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

func getPkgFromSchemaPath(schemaPath string) (string, error) {
	parts := strings.Split(schemaPath, "/")
	lastPart := parts[len(parts)-1]
	if len(lastPart) == 0 {
		return "", errors.New("invalid schema path, unable to find package name in path")
	}
	return lastPart, nil
}
