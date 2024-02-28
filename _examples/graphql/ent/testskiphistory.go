// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/testskiphistory"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// TestSkipHistory is the model entity for the TestSkipHistory schema.
type TestSkipHistory struct {
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
	// OtherID holds the value of the "other_id" field.
	OtherID uuid.UUID `json:"other_id,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TestSkipHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case testskiphistory.FieldUpdatedBy:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case testskiphistory.FieldOperation, testskiphistory.FieldName:
			values[i] = new(sql.NullString)
		case testskiphistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case testskiphistory.FieldID, testskiphistory.FieldRef, testskiphistory.FieldOtherID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TestSkipHistory fields.
func (tsh *TestSkipHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case testskiphistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tsh.ID = *value
			}
		case testskiphistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				tsh.HistoryTime = value.Time
			}
		case testskiphistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				tsh.Operation = enthistory.OpType(value.String)
			}
		case testskiphistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				tsh.Ref = *value
			}
		case testskiphistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				tsh.UpdatedBy = new(uuid.UUID)
				*tsh.UpdatedBy = *value.S.(*uuid.UUID)
			}
		case testskiphistory.FieldOtherID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field other_id", values[i])
			} else if value != nil {
				tsh.OtherID = *value
			}
		case testskiphistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tsh.Name = value.String
			}
		default:
			tsh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TestSkipHistory.
// This includes values selected through modifiers, order, etc.
func (tsh *TestSkipHistory) Value(name string) (ent.Value, error) {
	return tsh.selectValues.Get(name)
}

// Update returns a builder for updating this TestSkipHistory.
// Note that you need to call TestSkipHistory.Unwrap() before calling this method if this TestSkipHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (tsh *TestSkipHistory) Update() *TestSkipHistoryUpdateOne {
	return NewTestSkipHistoryClient(tsh.config).UpdateOne(tsh)
}

// Unwrap unwraps the TestSkipHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tsh *TestSkipHistory) Unwrap() *TestSkipHistory {
	_tx, ok := tsh.config.driver.(*txDriver)
	if !ok {
		panic("ent: TestSkipHistory is not a transactional entity")
	}
	tsh.config.driver = _tx.drv
	return tsh
}

// String implements the fmt.Stringer.
func (tsh *TestSkipHistory) String() string {
	var builder strings.Builder
	builder.WriteString("TestSkipHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tsh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(tsh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", tsh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", tsh.Ref))
	builder.WriteString(", ")
	if v := tsh.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("other_id=")
	builder.WriteString(fmt.Sprintf("%v", tsh.OtherID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tsh.Name)
	builder.WriteByte(')')
	return builder.String()
}

// TestSkipHistories is a parsable slice of TestSkipHistory.
type TestSkipHistories []*TestSkipHistory
