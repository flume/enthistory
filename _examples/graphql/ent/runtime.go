// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/schema"
	"_examples/graphql/ent/testexclude"
	"_examples/graphql/ent/testskip"
	"_examples/graphql/ent/testskiphistory"
	"_examples/graphql/ent/todo"
	"_examples/graphql/ent/todohistory"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	testexcludeFields := schema.TestExclude{}.Fields()
	_ = testexcludeFields
	// testexcludeDescName is the schema descriptor for name field.
	testexcludeDescName := testexcludeFields[2].Descriptor()
	// testexclude.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testexclude.NameValidator = testexcludeDescName.Validators[0].(func(string) error)
	// testexcludeDescID is the schema descriptor for id field.
	testexcludeDescID := testexcludeFields[0].Descriptor()
	// testexclude.DefaultID holds the default value on creation for the id field.
	testexclude.DefaultID = testexcludeDescID.Default.(func() uuid.UUID)
	testskipFields := schema.TestSkip{}.Fields()
	_ = testskipFields
	// testskipDescName is the schema descriptor for name field.
	testskipDescName := testskipFields[2].Descriptor()
	// testskip.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testskip.NameValidator = testskipDescName.Validators[0].(func(string) error)
	// testskipDescID is the schema descriptor for id field.
	testskipDescID := testskipFields[0].Descriptor()
	// testskip.DefaultID holds the default value on creation for the id field.
	testskip.DefaultID = testskipDescID.Default.(func() uuid.UUID)
	testskiphistoryFields := schema.TestSkipHistory{}.Fields()
	_ = testskiphistoryFields
	// testskiphistoryDescHistoryTime is the schema descriptor for history_time field.
	testskiphistoryDescHistoryTime := testskiphistoryFields[0].Descriptor()
	// testskiphistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	testskiphistory.DefaultHistoryTime = testskiphistoryDescHistoryTime.Default.(func() time.Time)
	// testskiphistoryDescName is the schema descriptor for name field.
	testskiphistoryDescName := testskiphistoryFields[6].Descriptor()
	// testskiphistory.NameValidator is a validator for the "name" field. It is called by the builders before save.
	testskiphistory.NameValidator = testskiphistoryDescName.Validators[0].(func(string) error)
	// testskiphistoryDescID is the schema descriptor for id field.
	testskiphistoryDescID := testskiphistoryFields[4].Descriptor()
	// testskiphistory.DefaultID holds the default value on creation for the id field.
	testskiphistory.DefaultID = testskiphistoryDescID.Default.(func() uuid.UUID)
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescName is the schema descriptor for name field.
	todoDescName := todoFields[2].Descriptor()
	// todo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todo.NameValidator = todoDescName.Validators[0].(func(string) error)
	// todoDescID is the schema descriptor for id field.
	todoDescID := todoFields[0].Descriptor()
	// todo.DefaultID holds the default value on creation for the id field.
	todo.DefaultID = todoDescID.Default.(func() uuid.UUID)
	todohistoryFields := schema.TodoHistory{}.Fields()
	_ = todohistoryFields
	// todohistoryDescHistoryTime is the schema descriptor for history_time field.
	todohistoryDescHistoryTime := todohistoryFields[0].Descriptor()
	// todohistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	todohistory.DefaultHistoryTime = todohistoryDescHistoryTime.Default.(func() time.Time)
	// todohistoryDescName is the schema descriptor for name field.
	todohistoryDescName := todohistoryFields[6].Descriptor()
	// todohistory.NameValidator is a validator for the "name" field. It is called by the builders before save.
	todohistory.NameValidator = todohistoryDescName.Validators[0].(func(string) error)
	// todohistoryDescID is the schema descriptor for id field.
	todohistoryDescID := todohistoryFields[4].Descriptor()
	// todohistory.DefaultID holds the default value on creation for the id field.
	todohistory.DefaultID = todohistoryDescID.Default.(func() uuid.UUID)
}
