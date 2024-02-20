package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"_examples/basic/ent/schema/mixins"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Annotations of the Store.
func (Store) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Store",
		},
	}
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("region"),
		// FK to `Organization` table
		field.UUID("organization_id", uuid.UUID{}),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).
			Ref("organization_stores").
			Unique().
			Required().
			Field("organization_id"),
	}
}

// Mixin of the Store.
func (Store) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}
