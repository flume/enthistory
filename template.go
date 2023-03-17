package enthistory

import (
	"entgo.io/ent/entc/gen"
)

func extractUpdatedByKey(val any) string {
	updatedBy, ok := val.(*UpdatedBy)
	if !ok || updatedBy == nil {
		return ""
	}
	return updatedBy.key
}

func extractUpdatedByValueType(val any) string {
	updatedBy, ok := val.(*UpdatedBy)
	if !ok || updatedBy == nil {
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
	funcMap := gen.Funcs
	funcMap["extractUpdatedByKey"] = extractUpdatedByKey
	funcMap["extractUpdatedByValueType"] = extractUpdatedByValueType
	t.Funcs(funcMap)
	return gen.MustParse(t.ParseFS(_templates, path))
}
