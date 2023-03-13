// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory"

	"time"
)

// FriendshipHistory holds the schema definition for the FriendshipHistory entity.
type FriendshipHistory struct {
	ent.Schema
}

// Annotations of the FriendshipHistory.
func (FriendshipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "friendship_history",
		},
	}
}

// Fields of the FriendshipHistory.
func (FriendshipHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("ref").
			Optional(),
		field.String("updated_by").
			Optional().
			Nillable(),
		field.Enum("operation").GoType(enthistory.OpType("")),
	}

	return append(historyFields, Friendship{}.Fields()...)
}

// Mixin of the FriendshipHistory.
func (FriendshipHistory) Mixin() []ent.Mixin {
	return Friendship{}.Mixin()
}
