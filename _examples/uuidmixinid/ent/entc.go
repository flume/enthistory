//go:build ignore

package main

import (
	"_examples/uuidmixinid/ent/schema"
	"fmt"
	"log"

	"entgo.io/ent"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory"
)

func main() {
	if err := enthistory.Generate("./schema", []ent.Interface{
		// Add all the schemas you want history tracking on here
		schema.MenuItem{},
	},
		enthistory.WithInheritIdType(),
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeUUID),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithImmutableFields(),
		// Without this line, all triggers will be used as the default
		enthistory.WithTriggers(enthistory.OpTypeInsert),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate("./schema", &gen.Config{
		IDType: &field.TypeInfo{Type: field.TypeString},
		Features: []gen.Feature{
			gen.FeatureSnapshot,
		},
	},
		entc.Extensions(
			enthistory.NewHistoryExtension(enthistory.WithAuditing()),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
