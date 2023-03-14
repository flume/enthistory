# enthistory
enthistory is an extension to generate history tables with ent

## Install
Install enthistory via `go get`
```shell
go get github.com/flume/enthistory
```
and add it the extension to ent by creating two files in your `ent` directory `entc.go` and `generate.go`

your entc.go should contain:
```go
//go:build ignore

package main

import (
	"log"
	"github.com/flume/enthistory"
	"entgo.io/ent/entc"
)

func main() {
	if err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				// Let enthistory know that the updatedBy field can derive 
				// the value (a string) from the context with this key
				enthistory.WithUpdatedByKey("userId"),
			),
		),
	); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

and your generate.go should contain:
```go
package ent

//go:generate go run -mod=mod entc.go
```

###

Then you can generate your history tables from your schema by running 
```shell
go generate ./ent
```

If you manage migrations on manually, you will want to create/generate new migrations for the newly created history tables.

## Usage

Your newly generated code creates the history tables for you for every single table you have. It also hooks up the hooks to the ent client so that you can start tracking history right away.
You can query the history tables directly, just like any other ent table, or you can query the history of a specific row using the `History()` method.

enthistory will also track the user updating the row if you provide it a key when initializing. Store a user's id, email, IP address, etc. in context with the key you provide for it to be tracked in history. 

For example, let's say we have a Character table, and we got a character from the table just now. We can also pull the history for that character directly via enthistory.

```go
// Create
client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
// Activate the history hooks on the client
client.WithHistory()
character, _ := client.Character.Create().SetName("Marceline").Save(ctx)
characterHistory, _ := character.History().All(ctx)
fmt.Println(len(characterHistory)) // 1

// Update
character, _ = character.Update().SetName("Marshall Lee").Save(ctx)
characterHistory, _ = character.History().All(ctx)
fmt.Println(len(characterHistory)) // 2

// Delete
client.Character.DeleteOne(character)
characterHistory, _ = character.History().All(ctx)
fmt.Println(len(characterHistory)) // 3
```

## Caveats

A few caveats to keep in mind when using enthistory

### Edges
To track edges with history, you will need to manage your own through tables. There's no way to hook into the ent generated through tables following the ent guide, through tables are fairly easy to manage yourself.
Note: You will not be able to track the history on these through tables if you use the setters for edges on the main schema tables. You must directly update the through tables with the information required.

Instead of `.AddFriends()`
```go
finn, _ := client.Character.Create().SetName("Finn the Human").Save(ctx)
jake, _ := client.Character.Create().SetName("Jake the Dog").Save(ctx)
finn, _ = finn.Update().AddFriends(jake).Save(ctx)
```
Use the Friendship through table
```go
finn, _ := client.Character.Create().SetName("Finn the Human").Save(ctx)
jake, _ := client.Character.Create().SetName("Jake the Dog").Save(ctx)
friendship, _ := client.Friendship.Create().SetCharacterID(finn.ID).SetFriendID(jake.ID).Save(ctx)
```

See the [ent docs](https://entgo.io/docs/schema-edges#edge-schema) for more information on through tables and edges

### Enums
If your ent schemas contain enum fields, you should be creating "enums" with Go and setting the `GoType` on the enum field.
This is because ent will generate a unique enum type for both your schema and the history table schema that won't play well together.

Instead of `.Values()`
```go
field.Enum("action").
    Values("PUSH", "PULL")
```
use `.GoType()`
```go
field.Enum("action").
    GoType(types.Action(""))
```

See the [ent docs](https://entgo.io/docs/schema-fields#enum-fields) for more information on Enums


## Contributing

Please see our [contributing](.github/CONTRIBUTING.md) and [code of conduct](.github/CODE_OF_CONDUCT.md) documentation
