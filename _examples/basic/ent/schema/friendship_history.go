// Code generated by enthistory, DO NOT EDIT.

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/flume/enthistory"
)

type FriendshipHistory struct {
	ent.Schema
}

func (FriendshipHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Immutable(),
		field.Time("history_time").
			Immutable().
			Default(time.Now),
		field.Enum("operation").
			Immutable().
			GoType(enthistory.OpType("")),
		field.String("ref").
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.Int("updated_by").
			Nillable().
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.Int("character_id").
			Immutable(),
		field.Int("friend_id").
			Immutable(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now)}
}
func (FriendshipHistory) Edges() []ent.Edge {
	return nil
}
func (FriendshipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "friendship_history"}}
}
func (FriendshipHistory) Indexes() []ent.Index {
	return []ent.Index{index.Fields("history_time")}
}
