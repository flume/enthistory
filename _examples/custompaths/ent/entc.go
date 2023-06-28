//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

const (
	schemaOriginPath  = "./some/otherschema"
	schemaHistoryPath = "./some/history"
)

func main() {
	if err := entc.Generate(schemaOriginPath,
		&gen.Config{
			Target:  "../internal/ent",
			Schema:  "github.com/flume/enthistory/_examples/custompaths/ent/some/path",
			Package: "github.com/flume/enthistory/_examples/custompaths/internal/ent",
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithOriginSchemaPath(schemaOriginPath),
				enthistory.WithHisotrySchemaPath(schemaHistoryPath),
				enthistory.WithOriginSchemaFullPkg("github.com/flume/enthistory/_examples/custompaths/ent/some/otherschema"),
			),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
