package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/flume/enthistory/_examples/basic/ent/schema/mixins"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Annotations of the Organization.
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Organization",
		},
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),		
		field.JSON("info", map[string]any{}).
			Optional(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organization_stores", Store.Type).			
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}
