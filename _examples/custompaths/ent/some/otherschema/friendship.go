package otherschema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Friendship holds the schema definition for the Friendship entity.
type Friendship struct {
	ent.Schema
}

// Annotations of the Friendship.
func (Friendship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "friendship",
		},
	}
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("character_id"),
		field.Int("friend_id"),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Required().
			Unique().
			Field("character_id"),
		edge.To("friend", Character.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
