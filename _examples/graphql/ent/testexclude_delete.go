// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testexclude"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestExcludeDelete is the builder for deleting a TestExclude entity.
type TestExcludeDelete struct {
	config
	hooks    []Hook
	mutation *TestExcludeMutation
}

// Where appends a list predicates to the TestExcludeDelete builder.
func (ted *TestExcludeDelete) Where(ps ...predicate.TestExclude) *TestExcludeDelete {
	ted.mutation.Where(ps...)
	return ted
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ted *TestExcludeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ted.sqlExec, ted.mutation, ted.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ted *TestExcludeDelete) ExecX(ctx context.Context) int {
	n, err := ted.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ted *TestExcludeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testexclude.Table, sqlgraph.NewFieldSpec(testexclude.FieldID, field.TypeUUID))
	if ps := ted.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ted.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ted.mutation.done = true
	return affected, err
}

// TestExcludeDeleteOne is the builder for deleting a single TestExclude entity.
type TestExcludeDeleteOne struct {
	ted *TestExcludeDelete
}

// Where appends a list predicates to the TestExcludeDelete builder.
func (tedo *TestExcludeDeleteOne) Where(ps ...predicate.TestExclude) *TestExcludeDeleteOne {
	tedo.ted.mutation.Where(ps...)
	return tedo
}

// Exec executes the deletion query.
func (tedo *TestExcludeDeleteOne) Exec(ctx context.Context) error {
	n, err := tedo.ted.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testexclude.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tedo *TestExcludeDeleteOne) ExecX(ctx context.Context) {
	if err := tedo.Exec(ctx); err != nil {
		panic(err)
	}
}
