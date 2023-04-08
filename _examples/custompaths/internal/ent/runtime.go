// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/flume/enthistory/_examples/custompaths/ent/some/otherschema"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/character"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/characterhistory"
	"github.com/flume/enthistory/_examples/custompaths/internal/ent/friendshiphistory"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	characterFields := otherschema.Character{}.Fields()
	_ = characterFields
	// characterDescAge is the schema descriptor for age field.
	characterDescAge := characterFields[0].Descriptor()
	// character.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	character.AgeValidator = characterDescAge.Validators[0].(func(int) error)
	characterhistoryFields := otherschema.CharacterHistory{}.Fields()
	_ = characterhistoryFields
	// characterhistoryDescHistoryTime is the schema descriptor for history_time field.
	characterhistoryDescHistoryTime := characterhistoryFields[0].Descriptor()
	// characterhistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	characterhistory.DefaultHistoryTime = characterhistoryDescHistoryTime.Default.(func() time.Time)
	// characterhistoryDescAge is the schema descriptor for age field.
	characterhistoryDescAge := characterhistoryFields[4].Descriptor()
	// characterhistory.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	characterhistory.AgeValidator = characterhistoryDescAge.Validators[0].(func(int) error)
	friendshiphistoryFields := otherschema.FriendshipHistory{}.Fields()
	_ = friendshiphistoryFields
	// friendshiphistoryDescHistoryTime is the schema descriptor for history_time field.
	friendshiphistoryDescHistoryTime := friendshiphistoryFields[0].Descriptor()
	// friendshiphistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	friendshiphistory.DefaultHistoryTime = friendshiphistoryDescHistoryTime.Default.(func() time.Time)
}
