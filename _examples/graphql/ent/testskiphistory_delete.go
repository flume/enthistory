// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testskiphistory"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TestSkipHistoryDelete is the builder for deleting a TestSkipHistory entity.
type TestSkipHistoryDelete struct {
	config
	hooks    []Hook
	mutation *TestSkipHistoryMutation
}

// Where appends a list predicates to the TestSkipHistoryDelete builder.
func (tshd *TestSkipHistoryDelete) Where(ps ...predicate.TestSkipHistory) *TestSkipHistoryDelete {
	tshd.mutation.Where(ps...)
	return tshd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tshd *TestSkipHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tshd.sqlExec, tshd.mutation, tshd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tshd *TestSkipHistoryDelete) ExecX(ctx context.Context) int {
	n, err := tshd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tshd *TestSkipHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(testskiphistory.Table, sqlgraph.NewFieldSpec(testskiphistory.FieldID, field.TypeUUID))
	if ps := tshd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tshd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tshd.mutation.done = true
	return affected, err
}

// TestSkipHistoryDeleteOne is the builder for deleting a single TestSkipHistory entity.
type TestSkipHistoryDeleteOne struct {
	tshd *TestSkipHistoryDelete
}

// Where appends a list predicates to the TestSkipHistoryDelete builder.
func (tshdo *TestSkipHistoryDeleteOne) Where(ps ...predicate.TestSkipHistory) *TestSkipHistoryDeleteOne {
	tshdo.tshd.mutation.Where(ps...)
	return tshdo
}

// Exec executes the deletion query.
func (tshdo *TestSkipHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := tshdo.tshd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{testskiphistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tshdo *TestSkipHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := tshdo.Exec(ctx); err != nil {
		panic(err)
	}
}