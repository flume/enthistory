// Code generated by enthistory, DO NOT EDIT.
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"

	"time"
)

// TestSkipHistory holds the schema definition for the TestSkipHistory entity.
type TestSkipHistory struct {
	ent.Schema
}

// Annotations of the TestSkipHistory.
func (TestSkipHistory) Annotations() []schema.Annotation {
	tablename := "testskip_history"
	annotations := append(TestSkip{}.Annotations(), entsql.Annotation{}, enthistory.Annotations{})
	for i, a := range annotations {
		switch ant := a.(type) {
		case entsql.Annotation:
			ant.Table = tablename
			annotations[i] = ant
		case enthistory.Annotations:
			ant.IsHistory = true
			ant.Exclude = true
			annotations[i] = ant
		}
	}
	return annotations
}

// Fields of the TestSkipHistory.
func (TestSkipHistory) Fields() []ent.Field {
	historyFields := []ent.Field{
		field.Time("history_time").
			Default(time.Now).
			Immutable(),
		field.UUID("ref", uuid.UUID{}).
			Immutable().
			Optional().
			Annotations(entgql.Annotation{Type: "ID"}),
		field.Enum("operation").
			GoType(enthistory.OpType("")).
			Immutable(),
		field.UUID("updated_by", uuid.UUID{}).
			Optional().
			Immutable().
			Nillable().
			Annotations(entgql.Annotation{Type: "ID"}),
	}

	original := TestSkip{}
	historyFields = append(historyFields, original.Fields()...)

	return historyFields
}

// Mixin of the TestSkipHistory.
func (TestSkipHistory) Mixin() []ent.Mixin {
	return TestSkip{}.Mixin()
}