//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	if err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureSnapshot},
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt),
				enthistory.WithAuditing(),
				enthistory.WithHistoryTimeIndex(),
			),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
