// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/without_updatedby/ent/character"
	"_examples/without_updatedby/ent/friendship"
	"_examples/without_updatedby/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipUpdate is the builder for updating Friendship entities.
type FriendshipUpdate struct {
	config
	hooks    []Hook
	mutation *FriendshipMutation
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fu *FriendshipUpdate) Where(ps ...predicate.Friendship) *FriendshipUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FriendshipUpdate) SetUpdatedAt(t time.Time) *FriendshipUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetCharacterID sets the "character_id" field.
func (fu *FriendshipUpdate) SetCharacterID(i int) *FriendshipUpdate {
	fu.mutation.SetCharacterID(i)
	return fu
}

// SetNillableCharacterID sets the "character_id" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableCharacterID(i *int) *FriendshipUpdate {
	if i != nil {
		fu.SetCharacterID(*i)
	}
	return fu
}

// SetFriendID sets the "friend_id" field.
func (fu *FriendshipUpdate) SetFriendID(i int) *FriendshipUpdate {
	fu.mutation.SetFriendID(i)
	return fu
}

// SetNillableFriendID sets the "friend_id" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableFriendID(i *int) *FriendshipUpdate {
	if i != nil {
		fu.SetFriendID(*i)
	}
	return fu
}

// SetCharacter sets the "character" edge to the Character entity.
func (fu *FriendshipUpdate) SetCharacter(c *Character) *FriendshipUpdate {
	return fu.SetCharacterID(c.ID)
}

// SetFriend sets the "friend" edge to the Character entity.
func (fu *FriendshipUpdate) SetFriend(c *Character) *FriendshipUpdate {
	return fu.SetFriendID(c.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fu *FriendshipUpdate) Mutation() *FriendshipMutation {
	return fu.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (fu *FriendshipUpdate) ClearCharacter() *FriendshipUpdate {
	fu.mutation.ClearCharacter()
	return fu
}

// ClearFriend clears the "friend" edge to the Character entity.
func (fu *FriendshipUpdate) ClearFriend() *FriendshipUpdate {
	fu.mutation.ClearFriend()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendshipUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendshipUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendshipUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendshipUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FriendshipUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := friendship.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FriendshipUpdate) check() error {
	if _, ok := fu.mutation.CharacterID(); fu.mutation.CharacterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.character"`)
	}
	if _, ok := fu.mutation.FriendID(); fu.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

func (fu *FriendshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeString))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(friendship.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.CharacterTable,
			Columns: []string{friendship.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.CharacterTable,
			Columns: []string{friendship.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FriendshipUpdateOne is the builder for updating a single Friendship entity.
type FriendshipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendshipMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FriendshipUpdateOne) SetUpdatedAt(t time.Time) *FriendshipUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetCharacterID sets the "character_id" field.
func (fuo *FriendshipUpdateOne) SetCharacterID(i int) *FriendshipUpdateOne {
	fuo.mutation.SetCharacterID(i)
	return fuo
}

// SetNillableCharacterID sets the "character_id" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableCharacterID(i *int) *FriendshipUpdateOne {
	if i != nil {
		fuo.SetCharacterID(*i)
	}
	return fuo
}

// SetFriendID sets the "friend_id" field.
func (fuo *FriendshipUpdateOne) SetFriendID(i int) *FriendshipUpdateOne {
	fuo.mutation.SetFriendID(i)
	return fuo
}

// SetNillableFriendID sets the "friend_id" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableFriendID(i *int) *FriendshipUpdateOne {
	if i != nil {
		fuo.SetFriendID(*i)
	}
	return fuo
}

// SetCharacter sets the "character" edge to the Character entity.
func (fuo *FriendshipUpdateOne) SetCharacter(c *Character) *FriendshipUpdateOne {
	return fuo.SetCharacterID(c.ID)
}

// SetFriend sets the "friend" edge to the Character entity.
func (fuo *FriendshipUpdateOne) SetFriend(c *Character) *FriendshipUpdateOne {
	return fuo.SetFriendID(c.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fuo *FriendshipUpdateOne) Mutation() *FriendshipMutation {
	return fuo.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (fuo *FriendshipUpdateOne) ClearCharacter() *FriendshipUpdateOne {
	fuo.mutation.ClearCharacter()
	return fuo
}

// ClearFriend clears the "friend" edge to the Character entity.
func (fuo *FriendshipUpdateOne) ClearFriend() *FriendshipUpdateOne {
	fuo.mutation.ClearFriend()
	return fuo
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fuo *FriendshipUpdateOne) Where(ps ...predicate.Friendship) *FriendshipUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendshipUpdateOne) Select(field string, fields ...string) *FriendshipUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friendship entity.
func (fuo *FriendshipUpdateOne) Save(ctx context.Context) (*Friendship, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) SaveX(ctx context.Context) *Friendship {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendshipUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FriendshipUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := friendship.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FriendshipUpdateOne) check() error {
	if _, ok := fuo.mutation.CharacterID(); fuo.mutation.CharacterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.character"`)
	}
	if _, ok := fuo.mutation.FriendID(); fuo.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

func (fuo *FriendshipUpdateOne) sqlSave(ctx context.Context) (_node *Friendship, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeString))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friendship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for _, f := range fields {
			if !friendship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(friendship.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.CharacterTable,
			Columns: []string{friendship.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.CharacterTable,
			Columns: []string{friendship.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Friendship{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
