// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"

	"time"
)

// TodoHistory holds the schema definition for the TodoHistory entity.
type TodoHistory struct {
	ent.Schema
}

// Annotations of the TodoHistory.
func (TodoHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "todo_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
	}
}

// Fields of the TodoHistory.
func (TodoHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.UUID("ref", uuid.UUID{}).
			Immutable().
			Optional().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
		field.UUID("updated_by", uuid.UUID{}).
			Optional().
			Immutable().
			Nillable().
			Annotations(entgql.Annotation{Type: "ID"}),
	}

	original := Todo{}
	historyFields = append(historyFields, original.Fields()...)

	return historyFields
}

// Mixin of the TodoHistory.
func (TodoHistory) Mixin() []ent.Mixin {
	return Todo{}.Mixin()
}
