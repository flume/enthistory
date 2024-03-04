// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/testdata/debug/internal/ent/friendshiphistory"
	"_examples/testdata/debug/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipHistoryUpdate is the builder for updating FriendshipHistory entities.
type FriendshipHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *FriendshipHistoryMutation
}

// Where appends a list predicates to the FriendshipHistoryUpdate builder.
func (fhu *FriendshipHistoryUpdate) Where(ps ...predicate.FriendshipHistory) *FriendshipHistoryUpdate {
	fhu.mutation.Where(ps...)
	return fhu
}

// Mutation returns the FriendshipHistoryMutation object of the builder.
func (fhu *FriendshipHistoryUpdate) Mutation() *FriendshipHistoryMutation {
	return fhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fhu *FriendshipHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fhu.sqlSave, fhu.mutation, fhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fhu *FriendshipHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := fhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fhu *FriendshipHistoryUpdate) Exec(ctx context.Context) error {
	_, err := fhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhu *FriendshipHistoryUpdate) ExecX(ctx context.Context) {
	if err := fhu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fhu *FriendshipHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendshiphistory.Table, friendshiphistory.Columns, sqlgraph.NewFieldSpec(friendshiphistory.FieldID, field.TypeUUID))
	if ps := fhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fhu.mutation.RefCleared() {
		_spec.ClearField(friendshiphistory.FieldRef, field.TypeUUID)
	}
	if fhu.mutation.UpdatedByCleared() {
		_spec.ClearField(friendshiphistory.FieldUpdatedBy, field.TypeUUID)
	}
	if fhu.mutation.CharacterIDCleared() {
		_spec.ClearField(friendshiphistory.FieldCharacterID, field.TypeUUID)
	}
	if fhu.mutation.FriendIDCleared() {
		_spec.ClearField(friendshiphistory.FieldFriendID, field.TypeUUID)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendshiphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fhu.mutation.done = true
	return n, nil
}

// FriendshipHistoryUpdateOne is the builder for updating a single FriendshipHistory entity.
type FriendshipHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendshipHistoryMutation
}

// Mutation returns the FriendshipHistoryMutation object of the builder.
func (fhuo *FriendshipHistoryUpdateOne) Mutation() *FriendshipHistoryMutation {
	return fhuo.mutation
}

// Where appends a list predicates to the FriendshipHistoryUpdate builder.
func (fhuo *FriendshipHistoryUpdateOne) Where(ps ...predicate.FriendshipHistory) *FriendshipHistoryUpdateOne {
	fhuo.mutation.Where(ps...)
	return fhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fhuo *FriendshipHistoryUpdateOne) Select(field string, fields ...string) *FriendshipHistoryUpdateOne {
	fhuo.fields = append([]string{field}, fields...)
	return fhuo
}

// Save executes the query and returns the updated FriendshipHistory entity.
func (fhuo *FriendshipHistoryUpdateOne) Save(ctx context.Context) (*FriendshipHistory, error) {
	return withHooks(ctx, fhuo.sqlSave, fhuo.mutation, fhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fhuo *FriendshipHistoryUpdateOne) SaveX(ctx context.Context) *FriendshipHistory {
	node, err := fhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fhuo *FriendshipHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := fhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhuo *FriendshipHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := fhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fhuo *FriendshipHistoryUpdateOne) sqlSave(ctx context.Context) (_node *FriendshipHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(friendshiphistory.Table, friendshiphistory.Columns, sqlgraph.NewFieldSpec(friendshiphistory.FieldID, field.TypeUUID))
	id, ok := fhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FriendshipHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendshiphistory.FieldID)
		for _, f := range fields {
			if !friendshiphistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendshiphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fhuo.mutation.RefCleared() {
		_spec.ClearField(friendshiphistory.FieldRef, field.TypeUUID)
	}
	if fhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(friendshiphistory.FieldUpdatedBy, field.TypeUUID)
	}
	if fhuo.mutation.CharacterIDCleared() {
		_spec.ClearField(friendshiphistory.FieldCharacterID, field.TypeUUID)
	}
	if fhuo.mutation.FriendIDCleared() {
		_spec.ClearField(friendshiphistory.FieldFriendID, field.TypeUUID)
	}
	_node = &FriendshipHistory{config: fhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendshiphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fhuo.mutation.done = true
	return _node, nil
}
