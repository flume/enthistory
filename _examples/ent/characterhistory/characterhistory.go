// Code generated by ent, DO NOT EDIT.

package characterhistory

import (
	"fmt"
	"time"

	"github.com/flume/enthistory"
)

const (
	// Label holds the string label denoting the characterhistory type in the database.
	Label = "character_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the characterhistory in the database.
	Table = "character_history"
)

// Columns holds all SQL columns for characterhistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldRef,
	FieldOperation,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldAge,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultHistoryTime holds the default value on creation for the "history_time" field.
	DefaultHistoryTime func() time.Time
	// UpdateDefaultHistoryTime holds the default value on update for the "history_time" field.
	UpdateDefaultHistoryTime func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("characterhistory: invalid enum value for operation field: %q", o)
	}
}
