// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/testdata/debug/internal/ent/character"
	"_examples/testdata/debug/internal/ent/characterhistory"
	"_examples/testdata/debug/internal/ent/friendship"
	"_examples/testdata/debug/internal/ent/friendshiphistory"
	"_examples/testdata/debug/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	characterFields := schema.Character{}.Fields()
	_ = characterFields
	// characterDescAge is the schema descriptor for age field.
	characterDescAge := characterFields[1].Descriptor()
	// character.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	character.AgeValidator = characterDescAge.Validators[0].(func(int) error)
	// characterDescID is the schema descriptor for id field.
	characterDescID := characterFields[0].Descriptor()
	// character.DefaultID holds the default value on creation for the id field.
	character.DefaultID = characterDescID.Default.(func() uuid.UUID)
	characterhistoryFields := schema.CharacterHistory{}.Fields()
	_ = characterhistoryFields
	// characterhistoryDescHistoryTime is the schema descriptor for history_time field.
	characterhistoryDescHistoryTime := characterhistoryFields[1].Descriptor()
	// characterhistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	characterhistory.DefaultHistoryTime = characterhistoryDescHistoryTime.Default.(func() time.Time)
	// characterhistoryDescID is the schema descriptor for id field.
	characterhistoryDescID := characterhistoryFields[0].Descriptor()
	// characterhistory.DefaultID holds the default value on creation for the id field.
	characterhistory.DefaultID = characterhistoryDescID.Default.(func() uuid.UUID)
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescID is the schema descriptor for id field.
	friendshipDescID := friendshipFields[0].Descriptor()
	// friendship.DefaultID holds the default value on creation for the id field.
	friendship.DefaultID = friendshipDescID.Default.(func() uuid.UUID)
	friendshiphistoryFields := schema.FriendshipHistory{}.Fields()
	_ = friendshiphistoryFields
	// friendshiphistoryDescHistoryTime is the schema descriptor for history_time field.
	friendshiphistoryDescHistoryTime := friendshiphistoryFields[1].Descriptor()
	// friendshiphistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	friendshiphistory.DefaultHistoryTime = friendshiphistoryDescHistoryTime.Default.(func() time.Time)
	// friendshiphistoryDescID is the schema descriptor for id field.
	friendshiphistoryDescID := friendshiphistoryFields[0].Descriptor()
	// friendshiphistory.DefaultID holds the default value on creation for the id field.
	friendshiphistory.DefaultID = friendshiphistoryDescID.Default.(func() uuid.UUID)
}