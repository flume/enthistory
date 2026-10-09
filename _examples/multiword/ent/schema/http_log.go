package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HTTPLog has an acronym followed by a word.
type HTTPLog struct {
	ent.Schema
}

// Fields of the HTTPLog.
func (HTTPLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("request_path"),
		field.Int("status_code").
			Default(200),
	}
}
