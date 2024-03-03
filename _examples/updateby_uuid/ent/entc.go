//go:build ignore

package main

import (
	"_examples/updateby_uuid/ent/schema"
	"fmt"
	"log"

	"entgo.io/ent"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	if err := enthistory.Generate("./schema", []ent.Interface{
		schema.Organization{},
		schema.Store{},
	},
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeUUID),
		enthistory.WithHistoryTimeIndex(),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(
			enthistory.NewHistoryExtension(enthistory.WithAuditing()),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
