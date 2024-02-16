// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/basic/ent/character"
	"_examples/basic/ent/friendship"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Friendship is the model entity for the Friendship schema.
type Friendship struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CharacterID holds the value of the "character_id" field.
	CharacterID int `json:"character_id,omitempty"`
	// FriendID holds the value of the "friend_id" field.
	FriendID int `json:"friend_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FriendshipQuery when eager-loading is set.
	Edges        FriendshipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FriendshipEdges holds the relations/edges for other nodes in the graph.
type FriendshipEdges struct {
	// Character holds the value of the character edge.
	Character *Character `json:"character,omitempty"`
	// Friend holds the value of the friend edge.
	Friend *Character `json:"friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CharacterOrErr returns the Character value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) CharacterOrErr() (*Character, error) {
	if e.loadedTypes[0] {
		if e.Character == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: character.Label}
		}
		return e.Character, nil
	}
	return nil, &NotLoadedError{edge: "character"}
}

// FriendOrErr returns the Friend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) FriendOrErr() (*Character, error) {
	if e.loadedTypes[1] {
		if e.Friend == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: character.Label}
		}
		return e.Friend, nil
	}
	return nil, &NotLoadedError{edge: "friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Friendship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friendship.FieldCharacterID, friendship.FieldFriendID:
			values[i] = new(sql.NullInt64)
		case friendship.FieldID:
			values[i] = new(sql.NullString)
		case friendship.FieldCreatedAt, friendship.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Friendship fields.
func (f *Friendship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friendship.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case friendship.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case friendship.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Time
			}
		case friendship.FieldCharacterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field character_id", values[i])
			} else if value.Valid {
				f.CharacterID = int(value.Int64)
			}
		case friendship.FieldFriendID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field friend_id", values[i])
			} else if value.Valid {
				f.FriendID = int(value.Int64)
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Friendship.
// This includes values selected through modifiers, order, etc.
func (f *Friendship) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryCharacter queries the "character" edge of the Friendship entity.
func (f *Friendship) QueryCharacter() *CharacterQuery {
	return NewFriendshipClient(f.config).QueryCharacter(f)
}

// QueryFriend queries the "friend" edge of the Friendship entity.
func (f *Friendship) QueryFriend() *CharacterQuery {
	return NewFriendshipClient(f.config).QueryFriend(f)
}

// Update returns a builder for updating this Friendship.
// Note that you need to call Friendship.Unwrap() before calling this method if this Friendship
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Friendship) Update() *FriendshipUpdateOne {
	return NewFriendshipClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Friendship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Friendship) Unwrap() *Friendship {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Friendship is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Friendship) String() string {
	var builder strings.Builder
	builder.WriteString("Friendship(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(f.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("character_id=")
	builder.WriteString(fmt.Sprintf("%v", f.CharacterID))
	builder.WriteString(", ")
	builder.WriteString("friend_id=")
	builder.WriteString(fmt.Sprintf("%v", f.FriendID))
	builder.WriteByte(')')
	return builder.String()
}

// Friendships is a parsable slice of Friendship.
type Friendships []*Friendship
