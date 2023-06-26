// Code generated by ent, DO NOT EDIT.

package friendshiphistory

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"

	"github.com/flume/enthistory"
)

const (
	// Label holds the string label denoting the friendshiphistory type in the database.
	Label = "friendship_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHistoryTime holds the string denoting the history_time field in the database.
	FieldHistoryTime = "history_time"
	// FieldRef holds the string denoting the ref field in the database.
	FieldRef = "ref"
	// FieldOperation holds the string denoting the operation field in the database.
	FieldOperation = "operation"
	// FieldCharacterID holds the string denoting the character_id field in the database.
	FieldCharacterID = "character_id"
	// FieldFriendID holds the string denoting the friend_id field in the database.
	FieldFriendID = "friend_id"
	// Table holds the table name of the friendshiphistory in the database.
	Table = "friendship_history"
)

// Columns holds all SQL columns for friendshiphistory fields.
var Columns = []string{
	FieldID,
	FieldHistoryTime,
	FieldRef,
	FieldOperation,
	FieldCharacterID,
	FieldFriendID,
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
)

// OperationValidator is a validator for the "operation" field enum values. It is called by the builders before save.
func OperationValidator(o enthistory.OpType) error {
	switch o.String() {
	case "INSERT", "UPDATE", "DELETE":
		return nil
	default:
		return fmt.Errorf("friendshiphistory: invalid enum value for operation field: %q", o)
	}
}

// OrderOption defines the ordering options for the FriendshipHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByHistoryTime orders the results by the history_time field.
func ByHistoryTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHistoryTime, opts...).ToFunc()
}

// ByRef orders the results by the ref field.
func ByRef(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRef, opts...).ToFunc()
}

// ByOperation orders the results by the operation field.
func ByOperation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperation, opts...).ToFunc()
}

// ByCharacterID orders the results by the character_id field.
func ByCharacterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCharacterID, opts...).ToFunc()
}

// ByFriendID orders the results by the friend_id field.
func ByFriendID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFriendID, opts...).ToFunc()
}
