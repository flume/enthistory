# enthistory

enthistory is a powerful extension for generating history tables using ent.

## Installation

You can install enthistory by running the following command:

```shell
go get github.com/flume/enthistory@latest
```

In addition to installing enthistory, you need to create two files in your `ent` directory: `entc.go` and `generate.go`.
The `entc.go` file should contain code similar to below:

```go
//go:build ignore

package main

import (
	"_examples/basic/ent/schema"
	"fmt"
	"log"

	"entgo.io/ent"

	"entgo.io/ent/entc/gen"

	"github.com/flume/enthistory"

	"entgo.io/ent/entc"
)

func main() {
	if err := enthistory.Generate("./schema", []ent.Interface{
		// Add all the schemas you want history tracking on here
		schema.Character{},
		schema.Friendship{},
		schema.Residence{},
	},
		enthistory.WithUpdatedBy("userId", enthistory.ValueTypeInt),
		enthistory.WithHistoryTimeIndex(),
		enthistory.WithImmutableFields(),
		// Without this line, all triggers will be used as the default
		enthistory.WithTriggers(enthistory.OpTypeInsert),
		// Enable reverse edges from history to original entities
		enthistory.WithReverseEdge(),
	); err != nil {
		log.Fatal(fmt.Sprintf("running enthistory codegen: %v", err))
	}

	if err := entc.Generate("./schema",
		&gen.Config{
			Features: []gen.Feature{gen.FeatureSnapshot},
		},
		entc.Extensions(
			enthistory.NewHistoryExtension(
				enthistory.WithAuditing(),
				// Required when using WithReverseEdge() above
				enthistory.WithReverseEdgeExtension(),
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

> **Note:** Breaking change introduced in enthistory version v0.12.0, if you are upgrading please see the example file above
> for the new way to use the enthistory extension.

## Usage

### Customizing History Schemas

You can customize the history tables to an extent by using `enthistory.Annotations{}` annotation on the original schema.
You can set custom annotations for the history schemas or add custom mixins. Adding custom annotations or mixins 
will overwrite the default annotations and mixins inherited from the original schema, so you need to include them if you 
want to keep the original annotations and mixins. Custom Mixins add the ability to add history specific 
policies, hooks, interceptors, etc. to the history table as well. You can also set custom triggers for tracking history
on a schema level. The triggers can be any combination of `OpTypeInsert`, `OpTypeUpdate`, `OpTypeDelete`. If you want to
set a global trigger set for all schemas you can set the triggers in the `enthistory.Generate()` function.
```go
type Annotations struct {
	// If you would like to add custom annotations to the history table,
	// otherwise it will default to the same annotations as the original table
	Annotations []schema.Annotation
	// if you would like to add custom mixins to the history table,
	// otherwise it will default to the same mixins as the original table
	Mixins []ent.Mixin
    // Global triggers for tracking history, can be any combination of OpTypeInsert, OpTypeUpdate, OpTypeDelete,
    // nil value will default to all triggers, to exclude all triggers set to an empty slice,
    // schema specific triggers will override these global triggers
    Triggers []OpType `json:"triggers,omitempty"`
}
```

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

## Schema Generation Configuration Options

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

### History Model ID Types

By default, history models do not inherit the ID type from the original models. If you want to use the same ID type for
history models as the original models, you can use the `enthistory.WithInheritIdType()` configuration option.

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

### Reverse Edge

By default, history tables do not have an edge back to the original entity. If you want to traverse from a history
record back to its original entity, you can enable reverse edges using the `enthistory.WithReverseEdge()` option in the
`Generate()` function.

```go
enthistory.Generate("./schema", []ent.Interface{
    schema.Character{},
},
    enthistory.WithReverseEdge(),
)
```

When enabled, history schemas will include an edge pointing back to the original entity. For example, `CharacterHistory`
will have a `character` edge that allows queries like:

```go
// Get a history record
historyRecord, _ := character.History().Earliest(ctx)

// Traverse from history back to the original character
originalCharacter, _ := historyRecord.QueryCharacter().Only(ctx)

// Or use eager loading
historyRecords, _ := client.CharacterHistory.Query().
    Where(characterhistory.RefEQ(characterID)).
    WithCharacter().
    All(ctx)
```

**Note:** This option must be used together with `enthistory.WithReverseEdgeExtension()` in the extension configuration
to fully enable the feature. The reverse edge uses `ON DELETE SET NULL` semantics, so querying the edge for a deleted
entity will return a "not found" error.

## Extension Configuration Options

### Auditing

As mentioned earlier, you can enable auditing by using the `enthistory.WithAuditing()` configuration option when
initializing the extension.

For a complete example of using a custom schema path, refer to the [custompaths](./_examples/custompaths/ent/entc.go)
example.

### Reverse Edge Extension

To enable reverse edges from history tables back to their original entities, you must use
`enthistory.WithReverseEdgeExtension()` when initializing the extension. This works together with
`enthistory.WithReverseEdge()` in the schema generation to populate the foreign key when creating history records.

```go
entc.Extensions(
    enthistory.NewHistoryExtension(
        enthistory.WithReverseEdgeExtension(),
    ),
)
```

When both options are enabled together, history records will automatically have their reverse edge foreign key populated
during create, update, and delete operations. This allows you to traverse from any history record back to the original
entity using the generated query methods (e.g., `QueryCharacter()`, `WithCharacter()`).

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
