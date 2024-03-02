package enthistory

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type history struct {
	ent.Schema
	ref ent.Field
}

func (h history) Fields() []ent.Field {
	return []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.Enum("operation").
			GoType(OpType("")).
			Immutable(),
		h.ref,
	}
}
