package enthistory

import (
	"strings"
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
	case ValueTypeUUID:
		return "uuid.UUID"
	default:
		return ""
	}
}

// func fieldPropertiesNillable(config Config) bool {
// 	return config.FieldProperties != nil && config.FieldProperties.Nillable
// }

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

// func parseTemplate(name, path string) *gen.Template {
// 	t := gen.NewTemplate(name)
// 	t.Funcs(template.FuncMap{
// 		"extractUpdatedByKey":       extractUpdatedByKey,
// 		"extractUpdatedByValueType": extractUpdatedByValueType,
// 		"fieldPropertiesNillable":   fieldPropertiesNillable,
// 		"isSlice":                   isSlice,
// 		"in":                        in,
// 	})
// 	return gen.MustParse(t.ParseFS(_templates, path))
// }
