// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/testdata/debug/internal/ent/characterhistory"
	"_examples/testdata/debug/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// CharacterHistory is the model entity for the CharacterHistory schema.
type CharacterHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// HistoryTime holds the value of the "history_time" field.
	HistoryTime time.Time `json:"history_time,omitempty"`
	// Operation holds the value of the "operation" field.
	Operation enthistory.OpType `json:"operation,omitempty"`
	// Ref holds the value of the "ref" field.
	Ref uuid.UUID `json:"ref,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty"`
	// Age holds the value of the "age" field.
	Age *int `json:"age,omitempty"`
	// TypedAge holds the value of the "typed_age" field.
	TypedAge *models.Uint64 `json:"typed_age,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Nicknames holds the value of the "nicknames" field.
	Nicknames []string `json:"nicknames,omitempty"`
	// Info holds the value of the "info" field.
	Info map[string]interface{} `json:"info,omitempty"`
	// InfoStruct holds the value of the "info_struct" field.
	InfoStruct models.InfoStruct `json:"info_struct,omitempty"`
	// Species holds the value of the "species" field.
	Species      *models.SpeciesType `json:"species,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CharacterHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case characterhistory.FieldUpdatedBy:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case characterhistory.FieldNicknames, characterhistory.FieldInfo, characterhistory.FieldInfoStruct:
			values[i] = new([]byte)
		case characterhistory.FieldAge, characterhistory.FieldTypedAge:
			values[i] = new(sql.NullInt64)
		case characterhistory.FieldOperation, characterhistory.FieldName, characterhistory.FieldSpecies:
			values[i] = new(sql.NullString)
		case characterhistory.FieldCreatedAt, characterhistory.FieldUpdatedAt, characterhistory.FieldHistoryTime:
			values[i] = new(sql.NullTime)
		case characterhistory.FieldID, characterhistory.FieldRef:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CharacterHistory fields.
func (ch *CharacterHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case characterhistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ch.ID = *value
			}
		case characterhistory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ch.CreatedAt = value.Time
			}
		case characterhistory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ch.UpdatedAt = value.Time
			}
		case characterhistory.FieldHistoryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field history_time", values[i])
			} else if value.Valid {
				ch.HistoryTime = value.Time
			}
		case characterhistory.FieldOperation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation", values[i])
			} else if value.Valid {
				ch.Operation = enthistory.OpType(value.String)
			}
		case characterhistory.FieldRef:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ref", values[i])
			} else if value != nil {
				ch.Ref = *value
			}
		case characterhistory.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ch.UpdatedBy = new(uuid.UUID)
				*ch.UpdatedBy = *value.S.(*uuid.UUID)
			}
		case characterhistory.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				ch.Age = new(int)
				*ch.Age = int(value.Int64)
			}
		case characterhistory.FieldTypedAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field typed_age", values[i])
			} else if value.Valid {
				ch.TypedAge = new(models.Uint64)
				*ch.TypedAge = models.Uint64(value.Int64)
			}
		case characterhistory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ch.Name = new(string)
				*ch.Name = value.String
			}
		case characterhistory.FieldNicknames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field nicknames", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ch.Nicknames); err != nil {
					return fmt.Errorf("unmarshal field nicknames: %w", err)
				}
			}
		case characterhistory.FieldInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ch.Info); err != nil {
					return fmt.Errorf("unmarshal field info: %w", err)
				}
			}
		case characterhistory.FieldInfoStruct:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info_struct", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ch.InfoStruct); err != nil {
					return fmt.Errorf("unmarshal field info_struct: %w", err)
				}
			}
		case characterhistory.FieldSpecies:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field species", values[i])
			} else if value.Valid {
				ch.Species = new(models.SpeciesType)
				*ch.Species = models.SpeciesType(value.String)
			}
		default:
			ch.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CharacterHistory.
// This includes values selected through modifiers, order, etc.
func (ch *CharacterHistory) Value(name string) (ent.Value, error) {
	return ch.selectValues.Get(name)
}

// Update returns a builder for updating this CharacterHistory.
// Note that you need to call CharacterHistory.Unwrap() before calling this method if this CharacterHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ch *CharacterHistory) Update() *CharacterHistoryUpdateOne {
	return NewCharacterHistoryClient(ch.config).UpdateOne(ch)
}

// Unwrap unwraps the CharacterHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ch *CharacterHistory) Unwrap() *CharacterHistory {
	_tx, ok := ch.config.driver.(*txDriver)
	if !ok {
		panic("ent: CharacterHistory is not a transactional entity")
	}
	ch.config.driver = _tx.drv
	return ch
}

// String implements the fmt.Stringer.
func (ch *CharacterHistory) String() string {
	var builder strings.Builder
	builder.WriteString("CharacterHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ch.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ch.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ch.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("history_time=")
	builder.WriteString(ch.HistoryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("operation=")
	builder.WriteString(fmt.Sprintf("%v", ch.Operation))
	builder.WriteString(", ")
	builder.WriteString("ref=")
	builder.WriteString(fmt.Sprintf("%v", ch.Ref))
	builder.WriteString(", ")
	if v := ch.UpdatedBy; v != nil {
		builder.WriteString("updated_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ch.Age; v != nil {
		builder.WriteString("age=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ch.TypedAge; v != nil {
		builder.WriteString("typed_age=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ch.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("nicknames=")
	builder.WriteString(fmt.Sprintf("%v", ch.Nicknames))
	builder.WriteString(", ")
	builder.WriteString("info=")
	builder.WriteString(fmt.Sprintf("%v", ch.Info))
	builder.WriteString(", ")
	builder.WriteString("info_struct=")
	builder.WriteString(fmt.Sprintf("%v", ch.InfoStruct))
	builder.WriteString(", ")
	if v := ch.Species; v != nil {
		builder.WriteString("species=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// CharacterHistories is a parsable slice of CharacterHistory.
type CharacterHistories []*CharacterHistory
