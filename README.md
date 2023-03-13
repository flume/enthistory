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
		// userId key in context with the user's information, value should be a string
		entc.Extensions(enthistory.NewHistoryExtension("userId")),
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

For example, let's say we have a User table, and we got a user from the table just now. We can also pull the history for that user directly via enthistory.

```go
// Create
client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
// Activate the history hooks on the client
client.WithHistory()
user, _ := client.User.Create().SetName("BMS").Save(ctx)
userHistory, _ := user.History().All(ctx)
fmt.Println(len(userHistory)) // 1

// Update
user, _ = user.Update().SetName("Marceline").Save(ctx)
userHistory, _ = user.History().All(ctx)
fmt.Println(len(userHistory)) // 2

// Delete
client.User.DeleteOne(user)
userHistory, _ = user.History().All(ctx)
fmt.Println(len(userHistory)) // 3
```
