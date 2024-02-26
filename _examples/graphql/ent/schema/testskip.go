// Test Schema courtesy of: github.com/nixxxon/entdemo

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestSkip holds the schema definition for the TestSkip entity.
type TestSkip struct {
	ent.Schema
}

// Fields of the TestSkip.
func (TestSkip) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New),
		field.UUID("other_id", uuid.New()).Optional().Annotations(
			entgql.Annotation{
				Type: "ID",
			},
		),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

func (TestSkip) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(),
		entsql.Annotation{
			Table: "testskip",
		},
	}
}
