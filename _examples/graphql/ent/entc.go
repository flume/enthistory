//go:build ignore

package main

import (
	"_examples/graphql/ent/schema"
	"log"

	"entgo.io/ent"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	err := enthistory.Generate("./ent/schema", []ent.Interface{
		schema.TestSkip{},
		schema.Todo{},
	},
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeUUID),
		enthistory.WithInheritIdType(),
	)
	if err != nil {
		log.Fatalf("failed to run enthistory codegen: %v", err)
	}

	gqlExtension, err := entgql.NewExtension(
		// Generate a GraphQL schema for the Ent schema
		// and save it as "query.graphql".
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaPath("graphql/query.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
	)
	if err != nil {
		log.Fatalf("failed to create entgql extension: %v", err)
	}

	historyExtension := enthistory.NewHistoryExtension()

	opts := []entc.Option{
		entc.Extensions(gqlExtension, historyExtension),
	}

	if err = entc.Generate("./ent/schema", &gen.Config{
		Target:  "ent",
		Package: "_examples/graphql/ent",
		Features: []gen.Feature{
			gen.FeatureEntQL,
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
		},
	}, opts...); err != nil {
		log.Fatalf("failed to run ent codegen: %v", err)
	}
}
