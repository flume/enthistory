package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory/_examples/ent/schema/mixins"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Annotations of the Character.
func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "character",
		},
	}
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name"),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Character.Type).
			Through("friendships", Friendship.Type),
	}
}

// Mixin of the Character.
func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}
