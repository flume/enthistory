package enthistory

import (
	"fmt"
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

func calculateHooks(original, history gen.Type) string {
	var triggers []OpType
	for n, f := range history.Annotations {
		if n == "History" {
			if h, ok := f.(map[string]any); ok {
				for k, v := range h {
					if k == "triggers" {
						for _, op := range v.([]any) {
							triggers = append(triggers, OpType(op.(string)))
						}
					}
				}
			}
		}
	}
	if len(triggers) == 0 {
		return "\n"
	}
	res := fmt.Sprintf("\n// %s hooks\n", original.Name)
	const hookPrefix = "enthistory.HistoryTrigger"
	const hookRegister = "c.%s.Use(%s)\n"
	for _, trigger := range triggers {
		switch {
		case trigger == OpTypeInsert:
			res += fmt.Sprintf(hookRegister, original.Name, fmt.Sprintf("%sInsert[*%sMutation]()", hookPrefix, original.Name))
		case trigger == OpTypeUpdate:
			res += fmt.Sprintf(hookRegister, original.Name, fmt.Sprintf("%sUpdate[*%sMutation]()", hookPrefix, original.Name))
		case trigger == OpTypeDelete:
			res += fmt.Sprintf(hookRegister, original.Name, fmt.Sprintf("%sDelete[*%sMutation]()", hookPrefix, original.Name))
		}
	}
	return res
}

func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"extractUpdatedByKey":       extractUpdatedByKey,
		"extractUpdatedByValueType": extractUpdatedByValueType,
		"isSlice":                   isSlice,
		"in":                        in,
		"isIdTypeUUID":              isIdTypeUUID,
		"calculateHooks":            calculateHooks,
	})
	return gen.MustParse(t.ParseFS(_templates, path))
}
