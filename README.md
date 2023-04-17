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
				enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt),
				enthistory.WithAuditing(),
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

### Querying History
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

A couple common history queries include the earliest history, the latest history, and the history of a row at a given point in time.
Since these queries come up often, we added in functions for them directly.

```go
character, _ := client.Character.Query().First(ctx)

// Get the earliest history for this character (i.e. when the character was created)
earliest, _ := character.History().Earliest(ctx)

// Get the latest history for this character (i.e. the current state of the actual character)
latest, _ := character.History().Latest(ctx)

// Get the history for this character as it was at a given point in time 
// (i.e. the state of the actual character at the given point in time)
historyNow, _ := character.History().AsOf(ctx, time.Now())
```

Once you have a history model, you can also use `.Next()` and `.Prev()` to pull the next/previous history in time.
```go
character, _ := client.Character.Query().First(ctx)

// Get the earliest history for this character (i.e. when the character was created)
earliest, _ := character.History().Earliest(ctx)

// Get the next history after the earliest history
next, _ := earliest.Next(ctx)

// Get the previous history before the next history
prev, _ := next.Prev(ctx)

// prev would now be the earliest history once again
fmt.Println(prev.ID == earliest.ID) // true
```

### Restoring History
In the event you want to rollback a row in the database to a particular history row, you can use the `.Restore()` function
to do accomplish that.

```go
// Let's say we create this character
simon, _ := client.Character.Create().SetName("Simon Petrikov").Save(ctx)
// And we update the character's name
iceking, _ := simon.Update().SetName("Ice King").Save(ctx)
// We can find the exact point in history we want to restore, in this case the oldest history row
icekingHistory, _ := iceking.History().Order(ent.Asc(characterhistory.FieldHistoryTime)).First(ctx)
// And we can restore value back to the original table
restored, _ = icekingHistory.Restore(ctx)

fmt.Println(simon.ID == restored.ID) // true
fmt.Println(simon.Name == restored.Name) // true
// The restoration is also tracked in history
simonHistory, _ := restored.History().All(ctx)
fmt.Println(len(simonHistory)) // 3
```

### Auditing

Another common use for history tables is for auditing, so `enthistory` has some helpful tools for maintaining/reviewing 
audits of the history tables. The largest of which is the `Audit()` method, which builds an audit log of the history tables
for you to export as a file, upload to S3, or simply inspect on your own. 

Auditing can be turned on by passing in the `enthistory.WithAuditing()` option to the `NewExtension()` Method

```go
// returns the audit log as a .csv file encoded as []byte
auditTable, _ = client.Audit(ctx)
```

The audit log has 6 columns when user tracking is turned on. An example audit log might look like this:

|Table            | Ref Id |History Time            |Operation| Changes                               |Updated By|
|-----------------|--------|------------------------|---------|---------------------------------------|----------|
|CharacterHistory | 1      |Sat Mar 18 16:31:31 2023|INSERT   | age: 47 name: "Simon Petrikov"        |75        |
|CharacterHistory | 1      |Sat Mar 18 16:31:31 2023|UPDATE   | name: "Simon Petrikov" -> "Ice King"  |75        |
|CharacterHistory | 1      |Sat Mar 18 16:31:31 2023|DELETE   | age: 47 name: "Ice King"              |75        |


The provided Audit log is obviously opinionated and if you want/need something different you can also build your own using
the `.Diff()` method on history models.

The History diff returned has the older history, the newer history, and the changes to fields when comparing the newer history
to the older history.
```go
...
nextHistory, _ := prismoHistory.Next(ctx)
diff, _ := prismoHistory.Diff(nextHistory)
```

## Config Options

### Setting all tracked fields as Nillable and/or Immutable
By default `enthistory` does not modify the columns in the history tables that are being tracked from your
original tables, it simply copies their state from `ent` when loading them. However, may want to set all tracked fields 
in the history tables as either `Nillable` or `Immutable` for a variety of reasons. If that is the case you can use the
`enthistory.WithNillableFields()` to set them all as `Nillable` or `enthistory.WithImmutableFields()` to set them all as `Immutable`.

**WARNING:** Setting `enthistory.WithNillableFields()` will remove the ability to call the `Restore()` function on a history object. 
Setting all fields to `Nillable` causes the history tables to diverge from the original tables and the unpredictability 
of that means the `Restore()` function cannot be generated.


### Updated By
To track which users are making which changes to your tables, you can supply the `enthistory.NewExtension()` function with 
the `enthistory.WithUpdatedBy()` Option. You choose your key name (string) and you can set either `enthistory.ValueTypeInt` (int) 
or `enthistory.ValueTypeString` (string) for the type of the value. This value would need to get populated in the context using 
`context.WithValue()`. You can leave it out entirely if you don't plan on using this feature.

```go
// context.WithValue(ctx, "userId", 5)
enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt)

// context.WithValue(ctx, "userEmail", "test@test.com")
enthistory.WithUpdatedBy("userEmail", enthistory.ValueTypeString)
```

### Auditing
Mentioned above, but you can turn on the ability to export an "audit" of the enthistory tables by setting the `enthistory.WithAudit()` 
config option when initializing the extension.

```go
enthistory.NewHistoryExtension((enthistory.WithAuditing())
```


### Excluding History on a Schema
`enthistory` has an always on philosophy but in instances you would like to not generate the history tables for a schema
you can apply annotations to the schema like so:

```go
func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		enthistory.Annotations{
			// Tells enthistory to exclude history tables for this schema
			Exclude: true,
		},
	}
}
```


### Setting a schema path
To allow setting an alternative schema location other than `ent/schema` you can supply the `enthistory.NewExtension()` function with 
the `enthistory.WithSchemaPath()` Option. You choose your schema path (string) which should be the same as the schema path set in the `entc.Generate` function
otherwise the extension will be unable to locate the schemas. You can omit the `enthistory.WithSchemaPath()` function if you don't plan on setting an alternative
schema location than `ent/schema`.

```go
func main() {
	entc.Generate("./schema2",
		&gen.Config{},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithSchemaPath("./schema2")
			),
		),
	)
}
```

See the [custompaths](./_examples/custompaths/ent/entc.go) example for a full example of how you can use this functional option.

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
