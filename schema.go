package enthistory

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type history struct {
	ent.Schema
}

func (history) Fields() []ent.Field {
	return []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.Int("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(OpType("")).
			Immutable(),
	}
}
