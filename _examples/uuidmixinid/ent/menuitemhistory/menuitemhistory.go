// Code generated by ent, DO NOT EDIT.

package menuitemhistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

const (
	// Label holds the string label denoting the menuitemhistory type in the database.
	Label = "menu_item_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// Table holds the table name of the menuitemhistory in the database.
	Table = "menu_item_history"
)

// Columns holds all SQL columns for menuitemhistory fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldHistoryTime,
	FieldOperation,
	FieldRef,
	FieldUpdatedBy,
	FieldName,
	FieldPrice,
	FieldDescription,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultHistoryTime holds the default value on creation for the "history_time" field.
	DefaultHistoryTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("menuitemhistory: invalid enum value for operation field: %q", o)
	}
}

// OrderOption defines the ordering options for the MenuItemHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByHistoryTime orders the results by the history_time field.
func ByHistoryTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHistoryTime, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByRef orders the results by the ref field.
func ByRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRef, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}