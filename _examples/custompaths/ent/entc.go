//go:build ignore

package main

import (
	"_examples/custompaths/ent/some/otherschema"
	"fmt"
	"log"

	"entgo.io/ent"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

const (
	schemaPath = "./some/otherschema"
)

func main() {
	if err := enthistory.Generate(schemaPath, []ent.Interface{
		&otherschema.Character{},
		&otherschema.Friendship{},
	}); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}
	if err := entc.Generate(schemaPath,
		&gen.Config{
			Target:  "../internal/ent",
			Schema:  "_examples/custompaths/ent/some/path",
			Package: "_examples/custompaths/internal/ent",
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
