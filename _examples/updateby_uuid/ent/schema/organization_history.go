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
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

type OrganizationHistory struct {
	ent.Schema
}

func (OrganizationHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Immutable(),
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
		field.String("name"),
		field.JSON("info", map[string]any{}).
			Optional(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now)}
}
func (OrganizationHistory) Edges() []ent.Edge {
	return nil
}
func (OrganizationHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "organization_history"}}
}
func (OrganizationHistory) Indexes() []ent.Index {
	return []ent.Index{index.Fields("history_time")}
}
