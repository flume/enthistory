//go:build ignore

package main

import (
	"_examples/multiword/ent/schema"
	"fmt"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"
)

func main() {
	if err := enthistory.Generate("./schema", []ent.Interface{
		&schema.UserProfile{},
		&schema.APIKey{},
		&schema.HTTPLog{},
	},
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithImmutableFields(),
		enthistory.WithReverseEdge(),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureSnapshot},
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithAuditing(),
				enthistory.WithReverseEdgeExtension(),
			),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
