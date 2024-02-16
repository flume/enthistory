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

// Residence holds the schema definition for the Residence entity.
type Residence struct {
	ent.Schema
}

// Annotations of the Residence.
func (Residence) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "residence",
		},
	}
}

// Fields of the Residence.
func (Residence) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
	}
}

// Edges of the Residence.
func (Residence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("occupants", Character.Type),
	}
}

// Mixin of the Residence.
func (Residence) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}
