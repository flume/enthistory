// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/uuidmixinid/ent/menuitem"
	"_examples/uuidmixinid/ent/menuitemhistory"
	"_examples/uuidmixinid/ent/schema"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	menuitemMixin := schema.MenuItem{}.Mixin()
	menuitemMixinFields0 := menuitemMixin[0].Fields()
	_ = menuitemMixinFields0
	menuitemMixinFields1 := menuitemMixin[1].Fields()
	_ = menuitemMixinFields1
	menuitemFields := schema.MenuItem{}.Fields()
	_ = menuitemFields
	// menuitemDescCreatedAt is the schema descriptor for created_at field.
	menuitemDescCreatedAt := menuitemMixinFields1[0].Descriptor()
	// menuitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	menuitem.DefaultCreatedAt = menuitemDescCreatedAt.Default.(func() time.Time)
	// menuitemDescUpdatedAt is the schema descriptor for updated_at field.
	menuitemDescUpdatedAt := menuitemMixinFields1[1].Descriptor()
	// menuitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menuitem.DefaultUpdatedAt = menuitemDescUpdatedAt.Default.(func() time.Time)
	// menuitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menuitem.UpdateDefaultUpdatedAt = menuitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// menuitemDescName is the schema descriptor for name field.
	menuitemDescName := menuitemFields[0].Descriptor()
	// menuitem.NameValidator is a validator for the "name" field. It is called by the builders before save.
	menuitem.NameValidator = menuitemDescName.Validators[0].(func(string) error)
	// menuitemDescPrice is the schema descriptor for price field.
	menuitemDescPrice := menuitemFields[1].Descriptor()
	// menuitem.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	menuitem.PriceValidator = menuitemDescPrice.Validators[0].(func(float64) error)
	// menuitemDescID is the schema descriptor for id field.
	menuitemDescID := menuitemMixinFields0[0].Descriptor()
	// menuitem.DefaultID holds the default value on creation for the id field.
	menuitem.DefaultID = menuitemDescID.Default.(func() uuid.UUID)
	menuitemhistoryMixin := schema.MenuItemHistory{}.Mixin()
	menuitemhistoryMixinFields1 := menuitemhistoryMixin[1].Fields()
	_ = menuitemhistoryMixinFields1
	menuitemhistoryFields := schema.MenuItemHistory{}.Fields()
	_ = menuitemhistoryFields
	// menuitemhistoryDescCreatedAt is the schema descriptor for created_at field.
	menuitemhistoryDescCreatedAt := menuitemhistoryMixinFields1[0].Descriptor()
	// menuitemhistory.DefaultCreatedAt holds the default value on creation for the created_at field.
	menuitemhistory.DefaultCreatedAt = menuitemhistoryDescCreatedAt.Default.(func() time.Time)
	// menuitemhistoryDescUpdatedAt is the schema descriptor for updated_at field.
	menuitemhistoryDescUpdatedAt := menuitemhistoryMixinFields1[1].Descriptor()
	// menuitemhistory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	menuitemhistory.DefaultUpdatedAt = menuitemhistoryDescUpdatedAt.Default.(func() time.Time)
	// menuitemhistory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	menuitemhistory.UpdateDefaultUpdatedAt = menuitemhistoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// menuitemhistoryDescHistoryTime is the schema descriptor for history_time field.
	menuitemhistoryDescHistoryTime := menuitemhistoryFields[1].Descriptor()
	// menuitemhistory.DefaultHistoryTime holds the default value on creation for the history_time field.
	menuitemhistory.DefaultHistoryTime = menuitemhistoryDescHistoryTime.Default.(func() time.Time)
	// menuitemhistoryDescID is the schema descriptor for id field.
	menuitemhistoryDescID := menuitemhistoryFields[0].Descriptor()
	// menuitemhistory.DefaultID holds the default value on creation for the id field.
	menuitemhistory.DefaultID = menuitemhistoryDescID.Default.(func() uuid.UUID)
}
