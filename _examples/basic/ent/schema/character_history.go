// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/flume/enthistory"

	"time"
)

// CharacterHistory holds the schema definition for the CharacterHistory entity.
type CharacterHistory struct {
	ent.Schema
}

// Annotations of the CharacterHistory.
func (CharacterHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "character_history",
		},
		enthistory.Annotations{
			IsHistory: true,
			Exclude:   true,
		},
	}
}

// Fields of the CharacterHistory.
func (CharacterHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.Int("ref").
			Immutable().
			Optional(),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
		field.Int("updated_by").
			Optional().
			Immutable().
			Nillable(),
	}

	original := Character{}
	for _, field := range original.Fields() {
		if field.Descriptor().Name != "id" {
			historyFields = append(historyFields, field)
		}
	}

	return historyFields
}

// Mixin of the CharacterHistory.
func (CharacterHistory) Mixin() []ent.Mixin {
	return Character{}.Mixin()
}
func (CharacterHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("history_time"),
	}
}
