package enthistory

import (
	"strings"
	"text/template"

	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/gen"
)

func extractUpdatedByKey(val any) string {
	ub, ok := val.(*UpdatedBy)
	if !ok || ub == nil {
		return ""
	}
	return ub.key
}

func extractUpdatedByValueType(val any) string {
	ub, ok := val.(*UpdatedBy)
	if !ok || ub == nil {
		return ""
	}

	switch ub.valueType {
	case ValueTypeInt:
		return "int"
	case ValueTypeString:
		return "string"
	case ValueTypeUUID:
		return "uuid.UUID"
	default:
		return ""
	}
}

func isSlice(typeString string) bool {
	return strings.HasPrefix(typeString, "[]")
}

func in(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

func isIdTypeUUID(node any) bool {
	return node.(*gen.Type).IDType.Type == field.TypeUUID
}

func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"extractUpdatedByKey":       extractUpdatedByKey,
		"extractUpdatedByValueType": extractUpdatedByValueType,
		"isSlice":                   isSlice,
		"in":                        in,
		"isIdTypeUUID":              isIdTypeUUID,
	})
	return gen.MustParse(t.ParseFS(_templates, path))
}
