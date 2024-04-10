// Code generated by enthistory, DO NOT EDIT.

package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type TestSkipHistory struct {
	ent.Schema
}

func (TestSkipHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("history_time").
			Immutable().
			Default(time.Now),
		field.Enum("operation").
			Immutable().
			GoType(enthistory.OpType("")),
		field.UUID("ref", uuid.UUID{}).
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.UUID("updated_by", uuid.UUID{}).
			Nillable().
			Optional().
			Immutable().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.UUID("other_id", uuid.UUID{}).
			Optional().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.String("name").
			Annotations(entgql.Annotation{OrderField: "NAME"})}
}
func (TestSkipHistory) Edges() []ent.Edge {
	return nil
}
func (TestSkipHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entgql.Annotation{Skip: entgql.SkipAll}, entsql.Annotation{Table: "testskip_history"}, enthistory.Annotations{IsHistory: true, Triggers: []enthistory.OpType{enthistory.OpTypeInsert, enthistory.OpTypeUpdate, enthistory.OpTypeDelete}}}
}
