// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testexclude"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestExcludeUpdate is the builder for updating TestExclude entities.
type TestExcludeUpdate struct {
	config
	hooks    []Hook
	mutation *TestExcludeMutation
}

// Where appends a list predicates to the TestExcludeUpdate builder.
func (teu *TestExcludeUpdate) Where(ps ...predicate.TestExclude) *TestExcludeUpdate {
	teu.mutation.Where(ps...)
	return teu
}

// SetOtherID sets the "other_id" field.
func (teu *TestExcludeUpdate) SetOtherID(u uuid.UUID) *TestExcludeUpdate {
	teu.mutation.SetOtherID(u)
	return teu
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (teu *TestExcludeUpdate) SetNillableOtherID(u *uuid.UUID) *TestExcludeUpdate {
	if u != nil {
		teu.SetOtherID(*u)
	}
	return teu
}

// ClearOtherID clears the value of the "other_id" field.
func (teu *TestExcludeUpdate) ClearOtherID() *TestExcludeUpdate {
	teu.mutation.ClearOtherID()
	return teu
}

// SetName sets the "name" field.
func (teu *TestExcludeUpdate) SetName(s string) *TestExcludeUpdate {
	teu.mutation.SetName(s)
	return teu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (teu *TestExcludeUpdate) SetNillableName(s *string) *TestExcludeUpdate {
	if s != nil {
		teu.SetName(*s)
	}
	return teu
}

// Mutation returns the TestExcludeMutation object of the builder.
func (teu *TestExcludeUpdate) Mutation() *TestExcludeMutation {
	return teu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (teu *TestExcludeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, teu.sqlSave, teu.mutation, teu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teu *TestExcludeUpdate) SaveX(ctx context.Context) int {
	affected, err := teu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (teu *TestExcludeUpdate) Exec(ctx context.Context) error {
	_, err := teu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teu *TestExcludeUpdate) ExecX(ctx context.Context) {
	if err := teu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (teu *TestExcludeUpdate) check() error {
	if v, ok := teu.mutation.Name(); ok {
		if err := testexclude.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestExclude.name": %w`, err)}
		}
	}
	return nil
}

func (teu *TestExcludeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := teu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(testexclude.Table, testexclude.Columns, sqlgraph.NewFieldSpec(testexclude.FieldID, field.TypeUUID))
	if ps := teu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teu.mutation.OtherID(); ok {
		_spec.SetField(testexclude.FieldOtherID, field.TypeUUID, value)
	}
	if teu.mutation.OtherIDCleared() {
		_spec.ClearField(testexclude.FieldOtherID, field.TypeUUID)
	}
	if value, ok := teu.mutation.Name(); ok {
		_spec.SetField(testexclude.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, teu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testexclude.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	teu.mutation.done = true
	return n, nil
}

// TestExcludeUpdateOne is the builder for updating a single TestExclude entity.
type TestExcludeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TestExcludeMutation
}

// SetOtherID sets the "other_id" field.
func (teuo *TestExcludeUpdateOne) SetOtherID(u uuid.UUID) *TestExcludeUpdateOne {
	teuo.mutation.SetOtherID(u)
	return teuo
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (teuo *TestExcludeUpdateOne) SetNillableOtherID(u *uuid.UUID) *TestExcludeUpdateOne {
	if u != nil {
		teuo.SetOtherID(*u)
	}
	return teuo
}

// ClearOtherID clears the value of the "other_id" field.
func (teuo *TestExcludeUpdateOne) ClearOtherID() *TestExcludeUpdateOne {
	teuo.mutation.ClearOtherID()
	return teuo
}

// SetName sets the "name" field.
func (teuo *TestExcludeUpdateOne) SetName(s string) *TestExcludeUpdateOne {
	teuo.mutation.SetName(s)
	return teuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (teuo *TestExcludeUpdateOne) SetNillableName(s *string) *TestExcludeUpdateOne {
	if s != nil {
		teuo.SetName(*s)
	}
	return teuo
}

// Mutation returns the TestExcludeMutation object of the builder.
func (teuo *TestExcludeUpdateOne) Mutation() *TestExcludeMutation {
	return teuo.mutation
}

// Where appends a list predicates to the TestExcludeUpdate builder.
func (teuo *TestExcludeUpdateOne) Where(ps ...predicate.TestExclude) *TestExcludeUpdateOne {
	teuo.mutation.Where(ps...)
	return teuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (teuo *TestExcludeUpdateOne) Select(field string, fields ...string) *TestExcludeUpdateOne {
	teuo.fields = append([]string{field}, fields...)
	return teuo
}

// Save executes the query and returns the updated TestExclude entity.
func (teuo *TestExcludeUpdateOne) Save(ctx context.Context) (*TestExclude, error) {
	return withHooks(ctx, teuo.sqlSave, teuo.mutation, teuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teuo *TestExcludeUpdateOne) SaveX(ctx context.Context) *TestExclude {
	node, err := teuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (teuo *TestExcludeUpdateOne) Exec(ctx context.Context) error {
	_, err := teuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teuo *TestExcludeUpdateOne) ExecX(ctx context.Context) {
	if err := teuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (teuo *TestExcludeUpdateOne) check() error {
	if v, ok := teuo.mutation.Name(); ok {
		if err := testexclude.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestExclude.name": %w`, err)}
		}
	}
	return nil
}

func (teuo *TestExcludeUpdateOne) sqlSave(ctx context.Context) (_node *TestExclude, err error) {
	if err := teuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(testexclude.Table, testexclude.Columns, sqlgraph.NewFieldSpec(testexclude.FieldID, field.TypeUUID))
	id, ok := teuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TestExclude.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := teuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testexclude.FieldID)
		for _, f := range fields {
			if !testexclude.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != testexclude.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := teuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teuo.mutation.OtherID(); ok {
		_spec.SetField(testexclude.FieldOtherID, field.TypeUUID, value)
	}
	if teuo.mutation.OtherIDCleared() {
		_spec.ClearField(testexclude.FieldOtherID, field.TypeUUID)
	}
	if value, ok := teuo.mutation.Name(); ok {
		_spec.SetField(testexclude.FieldName, field.TypeString, value)
	}
	_node = &TestExclude{config: teuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, teuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{testexclude.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	teuo.mutation.done = true
	return _node, nil
}
