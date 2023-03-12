//go:build ignore

package main

import (
	"log"

	"github.com/frisbm/enthistory"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(enthistory.NewHistoryExtension("userId")),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
