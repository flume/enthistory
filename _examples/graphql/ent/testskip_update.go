// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testskip"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestSkipUpdate is the builder for updating TestSkip entities.
type TestSkipUpdate struct {
	config
	hooks    []Hook
	mutation *TestSkipMutation
}

// Where appends a list predicates to the TestSkipUpdate builder.
func (tsu *TestSkipUpdate) Where(ps ...predicate.TestSkip) *TestSkipUpdate {
	tsu.mutation.Where(ps...)
	return tsu
}

// SetOtherID sets the "other_id" field.
func (tsu *TestSkipUpdate) SetOtherID(u uuid.UUID) *TestSkipUpdate {
	tsu.mutation.SetOtherID(u)
	return tsu
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (tsu *TestSkipUpdate) SetNillableOtherID(u *uuid.UUID) *TestSkipUpdate {
	if u != nil {
		tsu.SetOtherID(*u)
	}
	return tsu
}

// ClearOtherID clears the value of the "other_id" field.
func (tsu *TestSkipUpdate) ClearOtherID() *TestSkipUpdate {
	tsu.mutation.ClearOtherID()
	return tsu
}

// SetName sets the "name" field.
func (tsu *TestSkipUpdate) SetName(s string) *TestSkipUpdate {
	tsu.mutation.SetName(s)
	return tsu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tsu *TestSkipUpdate) SetNillableName(s *string) *TestSkipUpdate {
	if s != nil {
		tsu.SetName(*s)
	}
	return tsu
}

// Mutation returns the TestSkipMutation object of the builder.
func (tsu *TestSkipUpdate) Mutation() *TestSkipMutation {
	return tsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsu *TestSkipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tsu.sqlSave, tsu.mutation, tsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsu *TestSkipUpdate) SaveX(ctx context.Context) int {
	affected, err := tsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsu *TestSkipUpdate) Exec(ctx context.Context) error {
	_, err := tsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsu *TestSkipUpdate) ExecX(ctx context.Context) {
	if err := tsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsu *TestSkipUpdate) check() error {
	if v, ok := tsu.mutation.Name(); ok {
		if err := testskip.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestSkip.name": %w`, err)}
		}
	}
	return nil
}

func (tsu *TestSkipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(testskip.Table, testskip.Columns, sqlgraph.NewFieldSpec(testskip.FieldID, field.TypeUUID))
	if ps := tsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsu.mutation.OtherID(); ok {
		_spec.SetField(testskip.FieldOtherID, field.TypeUUID, value)
	}
	if tsu.mutation.OtherIDCleared() {
		_spec.ClearField(testskip.FieldOtherID, field.TypeUUID)
	}
	if value, ok := tsu.mutation.Name(); ok {
		_spec.SetField(testskip.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testskip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsu.mutation.done = true
	return n, nil
}

// TestSkipUpdateOne is the builder for updating a single TestSkip entity.
type TestSkipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestSkipMutation
}

// SetOtherID sets the "other_id" field.
func (tsuo *TestSkipUpdateOne) SetOtherID(u uuid.UUID) *TestSkipUpdateOne {
	tsuo.mutation.SetOtherID(u)
	return tsuo
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (tsuo *TestSkipUpdateOne) SetNillableOtherID(u *uuid.UUID) *TestSkipUpdateOne {
	if u != nil {
		tsuo.SetOtherID(*u)
	}
	return tsuo
}

// ClearOtherID clears the value of the "other_id" field.
func (tsuo *TestSkipUpdateOne) ClearOtherID() *TestSkipUpdateOne {
	tsuo.mutation.ClearOtherID()
	return tsuo
}

// SetName sets the "name" field.
func (tsuo *TestSkipUpdateOne) SetName(s string) *TestSkipUpdateOne {
	tsuo.mutation.SetName(s)
	return tsuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tsuo *TestSkipUpdateOne) SetNillableName(s *string) *TestSkipUpdateOne {
	if s != nil {
		tsuo.SetName(*s)
	}
	return tsuo
}

// Mutation returns the TestSkipMutation object of the builder.
func (tsuo *TestSkipUpdateOne) Mutation() *TestSkipMutation {
	return tsuo.mutation
}

// Where appends a list predicates to the TestSkipUpdate builder.
func (tsuo *TestSkipUpdateOne) Where(ps ...predicate.TestSkip) *TestSkipUpdateOne {
	tsuo.mutation.Where(ps...)
	return tsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsuo *TestSkipUpdateOne) Select(field string, fields ...string) *TestSkipUpdateOne {
	tsuo.fields = append([]string{field}, fields...)
	return tsuo
}

// Save executes the query and returns the updated TestSkip entity.
func (tsuo *TestSkipUpdateOne) Save(ctx context.Context) (*TestSkip, error) {
	return withHooks(ctx, tsuo.sqlSave, tsuo.mutation, tsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsuo *TestSkipUpdateOne) SaveX(ctx context.Context) *TestSkip {
	node, err := tsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsuo *TestSkipUpdateOne) Exec(ctx context.Context) error {
	_, err := tsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsuo *TestSkipUpdateOne) ExecX(ctx context.Context) {
	if err := tsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tsuo *TestSkipUpdateOne) check() error {
	if v, ok := tsuo.mutation.Name(); ok {
		if err := testskip.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestSkip.name": %w`, err)}
		}
	}
	return nil
}

func (tsuo *TestSkipUpdateOne) sqlSave(ctx context.Context) (_node *TestSkip, err error) {
	if err := tsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(testskip.Table, testskip.Columns, sqlgraph.NewFieldSpec(testskip.FieldID, field.TypeUUID))
	id, ok := tsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TestSkip.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testskip.FieldID)
		for _, f := range fields {
			if !testskip.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testskip.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsuo.mutation.OtherID(); ok {
		_spec.SetField(testskip.FieldOtherID, field.TypeUUID, value)
	}
	if tsuo.mutation.OtherIDCleared() {
		_spec.ClearField(testskip.FieldOtherID, field.TypeUUID)
	}
	if value, ok := tsuo.mutation.Name(); ok {
		_spec.SetField(testskip.FieldName, field.TypeString, value)
	}
	_node = &TestSkip{config: tsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testskip.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsuo.mutation.done = true
	return _node, nil
}
