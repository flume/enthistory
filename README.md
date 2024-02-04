# enthistory

enthistory is a powerful extension for generating history tables using ent.

## Installation

You can install enthistory by running the following command:

```shell
go get github.com/flume/enthistory@latest
```

In addition to installing enthistory, you need to create two files in your `ent` directory: `entc.go` and `generate.go`.
The `entc.go` file should contain the following code:

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

The `generate.go` file should contain the following code:

```go
package ent

//go:generate go run -mod=mod entc.go
```

> **Note:** Starting from enthistory v0.8.0, ent v0.12.x or greater is required. If you are using an older version of
> ent, install enthistory v0.7.0 instead by running `go get github.com/flume/enthistory@v0.7.0`.

## Usage

### Querying History

After generating your history tables from your schema, you can use the ent client to query the history tables. The
generated code automatically creates history tables for every table in your schema and hooks them up to the ent client.

You can query the history tables directly, just like any other ent table. You can also retrieve the history of a
specific row using the `History()` method.

enthistory tracks the user who updates a row if you provide a key during initialization. You can store a user's ID,
email, IP address, etc., in the context with the key you provide to track it in the history.

Here's an example that demonstrates these features:

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

In addition to regular queries, you can perform common history queries such as retrieving the earliest history, the
latest history, and the history of a row at a specific point in time. enthistory provides functions for these queries:

```go
character, _ := client.Character.Query().First(ctx)

// Get the earliest history for this character (i.e., when the character was created)
earliest, _ := character.History().Earliest(ctx)

// Get the latest history for this character (i.e., the current state of the actual character)
latest, _ := character.History().Latest(ctx)

// Get the history for this character as it was at a given point in time 
// (i.e., the state of the actual character at the given point in time)
historyNow, _ := character.History().AsOf(ctx, time.Now())
```

You can also use the `.Next()` and `.Prev()` methods to navigate to the next or previous history entries in time:

```go
character, _ := client.Character.Query().First(ctx)

// Get the earliest history for this character (i.e., when the character was created)
earliest, _ := character.History().Earliest(ctx)

// Get the next history after the earliest history
next, _ := earliest.Next(ctx)

// Get the previous history before the next history
prev, _ := next.Prev(ctx)

// prev would now be the earliest history once again
fmt.Println(prev.ID == earliest.ID) // true
```

### Restoring History

If you need to rollback a row in the database to a specific history entry, you can use the `.Restore()` function to
accomplish that. Here's an example:

```go
// Let's say we create this character
simon, _ := client.Character.Create().SetName("Simon Petrikov").Save(ctx)
// And we update the character's name
iceking, _ := simon.Update().SetName("Ice King").Save(ctx)
// We can find the exact point in history we want to restore, in this case, the oldest history entry
icekingHistory, _ := iceking.History().Order(ent.Asc(characterhistory.FieldHistoryTime)).First(ctx)
// And we can restore the value back to the original table
restored, _ = icekingHistory.Restore(ctx)

fmt.Println(simon.ID == restored.ID) // true
fmt.Println(simon.Name == restored.Name) // true
// The restoration is also tracked in history
simonHistory, _ := restored.History().All(ctx)
fmt.Println(len(simonHistory)) // 3
```

### Auditing

enthistory includes tools for auditing history tables. You can enable auditing by using the `enthistory.WithAuditing()`
option when initializing the extension. The main tool for auditing is the `Audit()` method, which builds an audit log of
the history tables that you can export as a file, upload to S3, or inspect.

Here's an example of how to use the `Audit()` method to export an audit log as a CSV file:

```go
auditTable, _ := client.Audit(ctx)
```

The audit log contains six columns when user tracking is enabled. Here's an example of how the audit log might look:

| Table            | Ref Id | History Time             | Operation | Changes                              | Updated By |
|------------------|--------|--------------------------|-----------|--------------------------------------|------------|
| CharacterHistory | 1      | Sat Mar 18 16:31:31 2023 | INSERT    | age: 47 name: "Simon Petrikov"       | 75         |
| CharacterHistory | 1      | Sat Mar 18 16:31:31 2023 | UPDATE    | name: "Simon Petrikov" -> "Ice King" | 75         |
| CharacterHistory | 1      | Sat Mar 18 16:31:31 2023 | DELETE    | age: 47 name: "Ice King"             | 75         |

You can also build your own custom audit log using the `.Diff()` method on history models. The `Diff()` method returns
the older history, the newer history, and the changes to fields when comparing the newer history to the older history.

## Configuration Options

enthistory provides several configuration options to customize its behavior.

### Setting All Tracked Fields as Nillable and/or Immutable

By default, enthistory does not modify the columns in the history tables that are being tracked from your original
tables; it simply copies their state from ent when loading them.

However, you may want to set all tracked fields in the history tables as either `Nillable` or `Immutable` for various
reasons. You can use the `enthistory.WithNillableFields()` option to set them all as `Nillable`,
or `enthistory.WithImmutableFields()` to set them all as `Immutable`.

**Note:** Setting `enthistory.WithNillableFields()` will remove the ability to call the `Restore()` function on a
history object. Setting all fields to `Nillable` causes the history tables to diverge from the original tables, and the
unpredictability of that means the `Restore()` function cannot be generated.

### History Time Indexing

By default, an index is not placed on the `history_time` field. If you want to enable indexing on the `history_time`
field, you can use the `enthistory.WithHistoryTimeIndex()` configuration option. This option gives you more control over
indexing based on your specific needs.

### Updated By

To track which users are making changes to your tables, you can use the `enthistory.WithUpdatedBy()` option when
initializing the extension. You need to provide a key name (string) and specify the type of
value (`enthistory.ValueTypeInt` for integers, `enthistory.ValueTypeUUID` for UUID or `enthistory.ValueTypeString` for strings). The value corresponding to
the key should be stored in the context using `context.WithValue()`. If you don't plan to use this feature, you can omit
it.

```go
// Example for tracking user ID
enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt)

// Example for tracking user as UUID
enthistory.WithUpdatedBy("userId", enthistory.ValueTypeUUID)

// Example for tracking user email
enthistory.WithUpdatedBy("userEmail", enthistory.ValueTypeString)
```

### Auditing

As mentioned earlier, you can enable auditing by using the `enthistory.WithAuditing()` configuration option when
initializing the extension.

### Excluding History on a Schema

enthistory is designed to always track history, but in cases where you don't want to generate history tables for a
particular schema, you can apply annotations to the schema to exclude it. Here's an example:

```go
func (Character) Annotations() []schema.Annotation {
    return []schema.Annotation{
        enthistory.Annotations{
            // Exclude history tables for this schema
            Exclude: true,
        },
    }
}
```

### Setting a Schema Path

If you want to set an alternative schema location other than `ent/schema`, you can use the `enthistory.WithSchemaPath()`
configuration option. The schema path should be the same as the one set in the `entc.Generate` function. If you don't
plan to set an alternative schema location, you can omit this option.

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

For a complete example of using a custom schema path, refer to the [custompaths](./_examples/custompaths/ent/entc.go)
example.

## Caveats

Here are a few caveats to keep in mind when using enthistory:

### Edges

To track edges with history, you need to manage your own through tables. enthistory does not hook into the ent-generated
through tables automatically, but managing through tables manually is straightforward. Note that if you use the setters
for edges on the main schema tables, the history on the through tables won't be tracked. To track history on through
tables, you must update the through tables directly with the required information.

Instead of using `.AddFriends()` like this:

```go
finn, _ := client.Character.Create().SetName("Finn the Human").Save(ctx)
jake, _ := client.Character.Create().SetName("Jake the Dog").Save(ctx)
finn, _ = finn.Update().AddFriends(jake).Save(ctx)
```

You should use the Friendship through table:

```go
finn, _ := client.Character.Create().SetName("Finn the Human").Save(ctx)
jake, _ := client.Character.Create().SetName("Jake the Dog").Save(ctx)
friendship, _ := client.Friendship.Create().SetCharacterID(finn.ID).SetFriendID(jake.ID).Save(ctx)
```

For more information on through tables and edges, refer to
the [ent documentation](https://entgo.io/docs/schema-edges#edge-schema).

### Enums

If your ent schemas contain enum fields, it is recommended to create Go enums and set the `GoType` on the enum field.
This is because ent generates a unique enum type for both your schema and the history table schema, which may not work
well together.

Instead of using `.Values()` like this:

```go
field.Enum("action").
    Values("PUSH", "PULL")
```

Use `.GoType()` like this:

```go
field.Enum("action").
    GoType(types.Action(""))
```

For more information on enums, refer to the [ent documentation](https://entgo.io/docs/schema-fields#enum-fields).

## Contributing

Please refer to our [contributing guidelines](.github/CONTRIBUTING.md) and [code of conduct](.github/CODE_OF_CONDUCT.md)
for information on how to contribute to enthistory.
