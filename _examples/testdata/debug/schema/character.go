package schema

import (
	"_examples/testdata/debug/models"
	"_examples/testdata/debug/schema/mixins"

	"github.com/flume/enthistory"

	"entgo.io/contrib/entgql"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		enthistory.Annotations{
			Mixins: []ent.Mixin{mixins.TimeMixin{}},
			Annotations: []schema.Annotation{
				entsql.Annotation{Table: "character_history"},
			},
			Triggers: []enthistory.OpType{enthistory.OpTypeUpdate},
		},
	}
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Annotations(entgql.Annotation{Type: "ID"}),
		field.Int("age").
			Positive(),
		field.Uint64("typed_age").
			Positive().
			GoType(models.Uint64(0)),
		field.String("name"),
		field.Strings("nicknames").
			Optional(),
		field.JSON("info", map[string]any{}).
			Optional(),
		field.JSON("info_struct", models.InfoStruct{}).
			Optional(),
		field.String("species").
			Optional().
			GoType(models.SpeciesType("")),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Character.Type).
			Through("friendships", Friendship.Type),
	}
}

func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
		OtherMixin{},
	}
}
