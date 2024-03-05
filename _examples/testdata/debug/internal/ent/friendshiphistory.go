// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/testdata/debug/internal/ent/friendshiphistory"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// FriendshipHistory is the model entity for the FriendshipHistory schema.
type FriendshipHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// CharacterID holds the value of the "character_id" field.
	CharacterID *uuid.UUID `json:"character_id,omitempty"`
	// FriendID holds the value of the "friend_id" field.
	FriendID     *uuid.UUID `json:"friend_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FriendshipHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friendshiphistory.FieldUpdatedBy, friendshiphistory.FieldCharacterID, friendshiphistory.FieldFriendID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case friendshiphistory.FieldOperation:
			values[i] = new(sql.NullString)
		case friendshiphistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case friendshiphistory.FieldID, friendshiphistory.FieldRef:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FriendshipHistory fields.
func (fh *FriendshipHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friendshiphistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fh.ID = *value
			}
		case friendshiphistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				fh.HistoryTime = value.Time
			}
		case friendshiphistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				fh.Operation = enthistory.OpType(value.String)
			}
		case friendshiphistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				fh.Ref = *value
			}
		case friendshiphistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				fh.UpdatedBy = new(uuid.UUID)
				*fh.UpdatedBy = *value.S.(*uuid.UUID)
			}
		case friendshiphistory.FieldCharacterID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field character_id", values[i])
			} else if value.Valid {
				fh.CharacterID = new(uuid.UUID)
				*fh.CharacterID = *value.S.(*uuid.UUID)
			}
		case friendshiphistory.FieldFriendID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field friend_id", values[i])
			} else if value.Valid {
				fh.FriendID = new(uuid.UUID)
				*fh.FriendID = *value.S.(*uuid.UUID)
			}
		default:
			fh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FriendshipHistory.
// This includes values selected through modifiers, order, etc.
func (fh *FriendshipHistory) Value(name string) (ent.Value, error) {
	return fh.selectValues.Get(name)
}

// Update returns a builder for updating this FriendshipHistory.
// Note that you need to call FriendshipHistory.Unwrap() before calling this method if this FriendshipHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (fh *FriendshipHistory) Update() *FriendshipHistoryUpdateOne {
	return NewFriendshipHistoryClient(fh.config).UpdateOne(fh)
}

// Unwrap unwraps the FriendshipHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fh *FriendshipHistory) Unwrap() *FriendshipHistory {
	_tx, ok := fh.config.driver.(*txDriver)
	if !ok {
		panic("ent: FriendshipHistory is not a transactional entity")
	}
	fh.config.driver = _tx.drv
	return fh
}

// String implements the fmt.Stringer.
func (fh *FriendshipHistory) String() string {
	var builder strings.Builder
	builder.WriteString("FriendshipHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(fh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", fh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", fh.Ref))
	builder.WriteString(", ")
	if v := fh.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := fh.CharacterID; v != nil {
		builder.WriteString("character_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := fh.FriendID; v != nil {
		builder.WriteString("friend_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// FriendshipHistories is a parsable slice of FriendshipHistory.
type FriendshipHistories []*FriendshipHistory
