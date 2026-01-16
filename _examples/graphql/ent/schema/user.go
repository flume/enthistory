// User schema demonstrates anonymous function Default for UUID fields.
// This is the exact reproduction case from the GitHub issue:
// "field.UUID with Default is incompatible with enthistory.WithInheritIdType"

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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// Using anonymous function Default - this is the exact reproduction case
		// from the GitHub issue that was previously failing with:
		// "schemast: only selector exprs are supported for default func"
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID {
			return uuid.Must(uuid.NewV7())
		}).Annotations(entgql.Annotation{Type: "ID"}),
		field.String("name").NotEmpty(),
		field.String("email").Optional(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entsql.Annotation{Table: "users"},
		enthistory.Annotations{
			Annotations: []schema.Annotation{
				entgql.RelayConnection(),
				entgql.QueryField(),
				entsql.Annotation{Table: "user_history"},
			},
		},
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
