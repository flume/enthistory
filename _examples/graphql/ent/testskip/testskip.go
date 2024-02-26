// Code generated by ent, DO NOT EDIT.

package testskip

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the testskip type in the database.
	Label = "test_skip"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOtherID holds the string denoting the other_id field in the database.
	FieldOtherID = "other_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the testskip in the database.
	Table = "testskip"
)

// Columns holds all SQL columns for testskip fields.
var Columns = []string{
	FieldID,
	FieldOtherID,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the TestSkip queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOtherID orders the results by the other_id field.
func ByOtherID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}
