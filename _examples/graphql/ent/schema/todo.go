// Test Schema courtesy of: github.com/nixxxon/entdemo

package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Annotations(entgql.Annotation{Type: "ID"}),
		field.UUID("other_id", uuid.New()).Optional().Annotations(entgql.Annotation{Type: "ID"}),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		enthistory.Annotations{
			Annotations: []schema.Annotation{
				// no mutations on TodoHistory
				entgql.RelayConnection(),
				entgql.QueryField(),
				entsql.Annotation{Table: "todo_history"},
			},
		},
	}
}
