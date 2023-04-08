//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

const (
	schemaPath = "./some/otherschema"
)

func main() {
	if err := entc.Generate(schemaPath,
		&gen.Config{
			Target:  "../internal/ent",
			Schema:  "github.com/flume/enthistory/_examples/custompaths/ent/some/path",
			Package: "github.com/flume/enthistory/_examples/custompaths/internal/ent",
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithSchemaPath(schemaPath),
			),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
