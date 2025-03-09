package schema

import (
	"_examples/basic/ent/schema/models"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"_examples/basic/ent/schema/mixins"
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
		field.Uint64("typed_age").
			Positive().
			GoType(models.Uint64(0)).
			DefaultFunc(models.DefaultUint64),
		field.String("name"),
		field.JSON("nicknames", []string{}).
			Optional(),
		field.JSON("info", map[string]any{}).
			Optional(),
		field.JSON("info_struct", models.InfoStruct{}).
			Optional().
			Default(models.DefaultInfoStruct),
		field.Int("level").
			Optional().
			Nillable(),
		field.String("species").
			Optional().
			GoType(models.SpeciesType("")).
			DefaultFunc(models.DefaultSpeciesType),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Character.Type).
			Through("friendships", Friendship.Type),
		edge.From("residence", Residence.Type).
			Ref("occupants").
			Unique(),
	}
}

// Mixin of the Character.
func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}
