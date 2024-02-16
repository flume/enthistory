// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/organization"
	"_examples/updateby_uuid/ent/store"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Store is the model entity for the Store schema.
type Store struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StoreQuery when eager-loading is set.
	Edges        StoreEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StoreEdges holds the relations/edges for other nodes in the graph.
type StoreEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StoreEdges) OrganizationOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Organization == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Organization, nil
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Store) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case store.FieldName, store.FieldRegion:
			values[i] = new(sql.NullString)
		case store.FieldCreatedAt, store.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case store.FieldID, store.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Store fields.
func (s *Store) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case store.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case store.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case store.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case store.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case store.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				s.Region = value.String
			}
		case store.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				s.OrganizationID = *value
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Store.
// This includes values selected through modifiers, order, etc.
func (s *Store) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the Store entity.
func (s *Store) QueryOrganization() *OrganizationQuery {
	return NewStoreClient(s.config).QueryOrganization(s)
}

// Update returns a builder for updating this Store.
// Note that you need to call Store.Unwrap() before calling this method if this Store
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Store) Update() *StoreUpdateOne {
	return NewStoreClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Store entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Store) Unwrap() *Store {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Store is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Store) String() string {
	var builder strings.Builder
	builder.WriteString("Store(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(s.Region)
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", s.OrganizationID))
	builder.WriteByte(')')
	return builder.String()
}

// Stores is a parsable slice of Store.
type Stores []*Store
