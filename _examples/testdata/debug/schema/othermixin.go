package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OtherMixin struct {
	mixin.Schema
}

func (OtherMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("other"),
	}
}
