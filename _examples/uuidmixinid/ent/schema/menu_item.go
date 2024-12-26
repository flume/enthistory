package schema

import (
	mixins2 "_examples/uuidmixinid/ent/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MenuItem holds the schema definition for the MenuItem entity.
type MenuItem struct {
	ent.Schema
}

// Annotations of the MenuItem.
func (MenuItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "menu_item",
		},
	}
}

// Fields of the MenuItem.
func (MenuItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Float("price").
			Positive(),
		field.String("description").
			Optional(),
	}
}

// Edges of the MenuItem.
func (MenuItem) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the MenuItem.
func (MenuItem) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins2.IDMixin{},
		mixins2.TimeMixin{},
	}
}
