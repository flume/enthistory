// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/basic/ent/character"
	"_examples/basic/ent/residence"
	"_examples/basic/ent/schema/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Character is the model entity for the Character schema.
type Character struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// TypedAge holds the value of the "typed_age" field.
	TypedAge models.Uint64 `json:"typed_age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nicknames holds the value of the "nicknames" field.
	Nicknames []string `json:"nicknames,omitempty"`
	// Info holds the value of the "info" field.
	Info map[string]interface{} `json:"info,omitempty"`
	// InfoStruct holds the value of the "info_struct" field.
	InfoStruct models.InfoStruct `json:"info_struct,omitempty"`
	// Level holds the value of the "level" field.
	Level *int `json:"level,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CharacterQuery when eager-loading is set.
	Edges               CharacterEdges `json:"edges"`
	residence_occupants *uuid.UUID
	selectValues        sql.SelectValues
}

// CharacterEdges holds the relations/edges for other nodes in the graph.
type CharacterEdges struct {
	// Friends holds the value of the friends edge.
	Friends []*Character `json:"friends,omitempty"`
	// Residence holds the value of the residence edge.
	Residence *Residence `json:"residence,omitempty"`
	// Friendships holds the value of the friendships edge.
	Friendships []*Friendship `json:"friendships,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) FriendsOrErr() ([]*Character, error) {
	if e.loadedTypes[0] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// ResidenceOrErr returns the Residence value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterEdges) ResidenceOrErr() (*Residence, error) {
	if e.Residence != nil {
		return e.Residence, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: residence.Label}
	}
	return nil, &NotLoadedError{edge: "residence"}
}

// FriendshipsOrErr returns the Friendships value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) FriendshipsOrErr() ([]*Friendship, error) {
	if e.loadedTypes[2] {
		return e.Friendships, nil
	}
	return nil, &NotLoadedError{edge: "friendships"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Character) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case character.FieldNicknames, character.FieldInfo, character.FieldInfoStruct:
			values[i] = new([]byte)
		case character.FieldID, character.FieldAge, character.FieldTypedAge, character.FieldLevel:
			values[i] = new(sql.NullInt64)
		case character.FieldName:
			values[i] = new(sql.NullString)
		case character.FieldCreatedAt, character.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case character.ForeignKeys[0]: // residence_occupants
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Character fields.
func (c *Character) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case character.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case character.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case character.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				c.Age = int(value.Int64)
			}
		case character.FieldTypedAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field typed_age", values[i])
			} else if value.Valid {
				c.TypedAge = models.Uint64(value.Int64)
			}
		case character.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case character.FieldNicknames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field nicknames", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Nicknames); err != nil {
					return fmt.Errorf("unmarshal field nicknames: %w", err)
				}
			}
		case character.FieldInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Info); err != nil {
					return fmt.Errorf("unmarshal field info: %w", err)
				}
			}
		case character.FieldInfoStruct:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field info_struct", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.InfoStruct); err != nil {
					return fmt.Errorf("unmarshal field info_struct: %w", err)
				}
			}
		case character.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				c.Level = new(int)
				*c.Level = int(value.Int64)
			}
		case character.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field residence_occupants", values[i])
			} else if value.Valid {
				c.residence_occupants = new(uuid.UUID)
				*c.residence_occupants = *value.S.(*uuid.UUID)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Character.
// This includes values selected through modifiers, order, etc.
func (c *Character) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryFriends queries the "friends" edge of the Character entity.
func (c *Character) QueryFriends() *CharacterQuery {
	return NewCharacterClient(c.config).QueryFriends(c)
}

// QueryResidence queries the "residence" edge of the Character entity.
func (c *Character) QueryResidence() *ResidenceQuery {
	return NewCharacterClient(c.config).QueryResidence(c)
}

// QueryFriendships queries the "friendships" edge of the Character entity.
func (c *Character) QueryFriendships() *FriendshipQuery {
	return NewCharacterClient(c.config).QueryFriendships(c)
}

// Update returns a builder for updating this Character.
// Note that you need to call Character.Unwrap() before calling this method if this Character
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Character) Update() *CharacterUpdateOne {
	return NewCharacterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Character entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Character) Unwrap() *Character {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Character is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Character) String() string {
	var builder strings.Builder
	builder.WriteString("Character(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", c.Age))
	builder.WriteString(", ")
	builder.WriteString("typed_age=")
	builder.WriteString(fmt.Sprintf("%v", c.TypedAge))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("nicknames=")
	builder.WriteString(fmt.Sprintf("%v", c.Nicknames))
	builder.WriteString(", ")
	builder.WriteString("info=")
	builder.WriteString(fmt.Sprintf("%v", c.Info))
	builder.WriteString(", ")
	builder.WriteString("info_struct=")
	builder.WriteString(fmt.Sprintf("%v", c.InfoStruct))
	builder.WriteString(", ")
	if v := c.Level; v != nil {
		builder.WriteString("level=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Characters is a parsable slice of Character.
type Characters []*Character
