// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/storehistory"
	"github.com/google/uuid"
)

// StoreHistory is the model entity for the StoreHistory schema.
type StoreHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organization_id,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StoreHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case storehistory.FieldUpdatedBy:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case storehistory.FieldID:
			values[i] = new(sql.NullInt64)
		case storehistory.FieldOperation, storehistory.FieldName, storehistory.FieldRegion:
			values[i] = new(sql.NullString)
		case storehistory.FieldHistoryTime, storehistory.FieldCreatedAt, storehistory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case storehistory.FieldRef, storehistory.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StoreHistory fields.
func (sh *StoreHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storehistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sh.ID = int(value.Int64)
		case storehistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				sh.HistoryTime = value.Time
			}
		case storehistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				sh.Operation = enthistory.OpType(value.String)
			}
		case storehistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				sh.Ref = *value
			}
		case storehistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				sh.UpdatedBy = new(uuid.UUID)
				*sh.UpdatedBy = *value.S.(*uuid.UUID)
			}
		case storehistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sh.CreatedAt = value.Time
			}
		case storehistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sh.UpdatedAt = value.Time
			}
		case storehistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sh.Name = value.String
			}
		case storehistory.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				sh.Region = value.String
			}
		case storehistory.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				sh.OrganizationID = *value
			}
		default:
			sh.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StoreHistory.
// This includes values selected through modifiers, order, etc.
func (sh *StoreHistory) Value(name string) (ent.Value, error) {
	return sh.selectValues.Get(name)
}

// Update returns a builder for updating this StoreHistory.
// Note that you need to call StoreHistory.Unwrap() before calling this method if this StoreHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sh *StoreHistory) Update() *StoreHistoryUpdateOne {
	return NewStoreHistoryClient(sh.config).UpdateOne(sh)
}

// Unwrap unwraps the StoreHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sh *StoreHistory) Unwrap() *StoreHistory {
	_tx, ok := sh.config.driver.(*txDriver)
	if !ok {
		panic("ent: StoreHistory is not a transactional entity")
	}
	sh.config.driver = _tx.drv
	return sh
}

// String implements the fmt.Stringer.
func (sh *StoreHistory) String() string {
	var builder strings.Builder
	builder.WriteString("StoreHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sh.ID))
	builder.WriteString("history_time=")
	builder.WriteString(sh.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", sh.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", sh.Ref))
	builder.WriteString(", ")
	if v := sh.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sh.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sh.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sh.Name)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(sh.Region)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", sh.OrganizationID))
	builder.WriteByte(')')
	return builder.String()
}

// StoreHistories is a parsable slice of StoreHistory.
type StoreHistories []*StoreHistory
