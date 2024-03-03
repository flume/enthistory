package main

import (
	"_examples/testdata/debug/schema"
	"fmt"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"entgo.io/ent"

	"github.com/flume/enthistory"
)

const (
	schemaPath = "./testdata/debug/schema"
)

func main() {
	if err := enthistory.Generate(schemaPath, []ent.Interface{
		&schema.Character{},
		&schema.Friendship{},
	},
		enthistory.WithInheritIdType(),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithUpdatedBy("userid", enthistory.ValueTypeUUID),
		enthistory.WithImmutableFields(),
		enthistory.WithNillableFields(),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate(schemaPath,
		&gen.Config{
			Target:  "./testdata/debug/internal/ent",
			Schema:  "_examples/testdata/debug/schema",
			Package: "_examples/testdata/debug/internal/ent",
		},
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
