//go:build ignore

package main

import (
	"log"

	"github.com/frisbm/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	if err := entc.Generate("./schema",
		entc.Extensions(enthistory.NewHistoryExtension("userId")),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
