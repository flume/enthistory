// Code generated by enthistory, DO NOT EDIT.

package schema

import (
	"_examples/testdata/debug/schema/mixins"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type CharacterHistory struct {
	ent.Schema
}

func (CharacterHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Annotations(entgql.Annotation{Type: "ID"}).
			Default(uuid.New),
		field.Time("history_time").
			Immutable().
			Default(time.Now),
		field.Enum("operation").
			Immutable().
			GoType(enthistory.OpType("")),
		field.UUID("ref", uuid.UUID{}).
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.UUID("updated_by", uuid.UUID{}).
			Nillable().
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.Int("age").
			Nillable().
			Optional().
			Immutable(),
		field.String("name").
			Nillable().
			Optional().
			Immutable(),
		field.JSON("nicknames", []string{}).
			Optional().
			Immutable(),
		field.JSON("info", map[string]any{}).
			Optional().
			Immutable()}
}
func (CharacterHistory) Edges() []ent.Edge {
	return nil
}
func (CharacterHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "character_history"}, enthistory.Annotations{IsHistory: true}}
}
func (CharacterHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{mixins.TimeMixin{}}
}
func (CharacterHistory) Indexes() []ent.Index {
	return []ent.Index{index.Fields("history_time")}
}
