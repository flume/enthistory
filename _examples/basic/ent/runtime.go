// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/basic/ent/character"
	"_examples/basic/ent/characterhistory"
	"_examples/basic/ent/friendship"
	"_examples/basic/ent/friendshiphistory"
	"_examples/basic/ent/residence"
	"_examples/basic/ent/residencehistory"
	"_examples/basic/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	characterMixin := schema.Character{}.Mixin()
	characterMixinFields0 := characterMixin[0].Fields()
	_ = characterMixinFields0
	characterFields := schema.Character{}.Fields()
	_ = characterFields
	// characterDescCreatedAt is the schema descriptor for created_at field.
	characterDescCreatedAt := characterMixinFields0[0].Descriptor()
	// character.DefaultCreatedAt holds the default value on creation for the created_at field.
	character.DefaultCreatedAt = characterDescCreatedAt.Default.(func() time.Time)
	// characterDescUpdatedAt is the schema descriptor for updated_at field.
	characterDescUpdatedAt := characterMixinFields0[1].Descriptor()
	// character.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	character.UpdateDefaultUpdatedAt = characterDescUpdatedAt.UpdateDefault.(func() time.Time)
	// characterDescAge is the schema descriptor for age field.
	characterDescAge := characterFields[0].Descriptor()
	// character.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	character.AgeValidator = characterDescAge.Validators[0].(func(int) error)
	characterhistoryMixin := schema.CharacterHistory{}.Mixin()
	characterhistoryMixinFields0 := characterhistoryMixin[0].Fields()
	_ = characterhistoryMixinFields0
	characterhistoryFields := schema.CharacterHistory{}.Fields()
	_ = characterhistoryFields
	// characterhistoryDescCreatedAt is the schema descriptor for created_at field.
	characterhistoryDescCreatedAt := characterhistoryMixinFields0[0].Descriptor()
	// characterhistory.DefaultCreatedAt holds the default value on creation for the created_at field.
	characterhistory.DefaultCreatedAt = characterhistoryDescCreatedAt.Default.(func() time.Time)
	// characterhistoryDescUpdatedAt is the schema descriptor for updated_at field.
	characterhistoryDescUpdatedAt := characterhistoryMixinFields0[1].Descriptor()
	// characterhistory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	characterhistory.UpdateDefaultUpdatedAt = characterhistoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// characterhistoryDescHistoryTime is the schema descriptor for history_time field.
	characterhistoryDescHistoryTime := characterhistoryFields[0].Descriptor()
	// characterhistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	characterhistory.DefaultHistoryTime = characterhistoryDescHistoryTime.Default.(func() time.Time)
	friendshipMixin := schema.Friendship{}.Mixin()
	friendshipMixinFields0 := friendshipMixin[0].Fields()
	_ = friendshipMixinFields0
	friendshipFields := schema.Friendship{}.Fields()
	_ = friendshipFields
	// friendshipDescCreatedAt is the schema descriptor for created_at field.
	friendshipDescCreatedAt := friendshipMixinFields0[0].Descriptor()
	// friendship.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendship.DefaultCreatedAt = friendshipDescCreatedAt.Default.(func() time.Time)
	// friendshipDescUpdatedAt is the schema descriptor for updated_at field.
	friendshipDescUpdatedAt := friendshipMixinFields0[1].Descriptor()
	// friendship.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	friendship.UpdateDefaultUpdatedAt = friendshipDescUpdatedAt.UpdateDefault.(func() time.Time)
	friendshiphistoryMixin := schema.FriendshipHistory{}.Mixin()
	friendshiphistoryMixinFields0 := friendshiphistoryMixin[0].Fields()
	_ = friendshiphistoryMixinFields0
	friendshiphistoryFields := schema.FriendshipHistory{}.Fields()
	_ = friendshiphistoryFields
	// friendshiphistoryDescCreatedAt is the schema descriptor for created_at field.
	friendshiphistoryDescCreatedAt := friendshiphistoryMixinFields0[0].Descriptor()
	// friendshiphistory.DefaultCreatedAt holds the default value on creation for the created_at field.
	friendshiphistory.DefaultCreatedAt = friendshiphistoryDescCreatedAt.Default.(func() time.Time)
	// friendshiphistoryDescUpdatedAt is the schema descriptor for updated_at field.
	friendshiphistoryDescUpdatedAt := friendshiphistoryMixinFields0[1].Descriptor()
	// friendshiphistory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	friendshiphistory.UpdateDefaultUpdatedAt = friendshiphistoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// friendshiphistoryDescHistoryTime is the schema descriptor for history_time field.
	friendshiphistoryDescHistoryTime := friendshiphistoryFields[1].Descriptor()
	// friendshiphistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	friendshiphistory.DefaultHistoryTime = friendshiphistoryDescHistoryTime.Default.(func() time.Time)
	residenceMixin := schema.Residence{}.Mixin()
	residenceMixinFields0 := residenceMixin[0].Fields()
	_ = residenceMixinFields0
	residenceFields := schema.Residence{}.Fields()
	_ = residenceFields
	// residenceDescCreatedAt is the schema descriptor for created_at field.
	residenceDescCreatedAt := residenceMixinFields0[0].Descriptor()
	// residence.DefaultCreatedAt holds the default value on creation for the created_at field.
	residence.DefaultCreatedAt = residenceDescCreatedAt.Default.(func() time.Time)
	// residenceDescUpdatedAt is the schema descriptor for updated_at field.
	residenceDescUpdatedAt := residenceMixinFields0[1].Descriptor()
	// residence.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	residence.UpdateDefaultUpdatedAt = residenceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// residenceDescID is the schema descriptor for id field.
	residenceDescID := residenceFields[0].Descriptor()
	// residence.DefaultID holds the default value on creation for the id field.
	residence.DefaultID = residenceDescID.Default.(func() uuid.UUID)
	residencehistoryMixin := schema.ResidenceHistory{}.Mixin()
	residencehistoryMixinFields0 := residencehistoryMixin[0].Fields()
	_ = residencehistoryMixinFields0
	residencehistoryFields := schema.ResidenceHistory{}.Fields()
	_ = residencehistoryFields
	// residencehistoryDescCreatedAt is the schema descriptor for created_at field.
	residencehistoryDescCreatedAt := residencehistoryMixinFields0[0].Descriptor()
	// residencehistory.DefaultCreatedAt holds the default value on creation for the created_at field.
	residencehistory.DefaultCreatedAt = residencehistoryDescCreatedAt.Default.(func() time.Time)
	// residencehistoryDescUpdatedAt is the schema descriptor for updated_at field.
	residencehistoryDescUpdatedAt := residencehistoryMixinFields0[1].Descriptor()
	// residencehistory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	residencehistory.UpdateDefaultUpdatedAt = residencehistoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// residencehistoryDescHistoryTime is the schema descriptor for history_time field.
	residencehistoryDescHistoryTime := residencehistoryFields[1].Descriptor()
	// residencehistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	residencehistory.DefaultHistoryTime = residencehistoryDescHistoryTime.Default.(func() time.Time)
}
