// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/custompaths/internal/ent/friendshiphistory"
	"_examples/custompaths/internal/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipHistoryDelete is the builder for deleting a FriendshipHistory entity.
type FriendshipHistoryDelete struct {
	config
	hooks    []Hook
	mutation *FriendshipHistoryMutation
}

// Where appends a list predicates to the FriendshipHistoryDelete builder.
func (fhd *FriendshipHistoryDelete) Where(ps ...predicate.FriendshipHistory) *FriendshipHistoryDelete {
	fhd.mutation.Where(ps...)
	return fhd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fhd *FriendshipHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fhd.sqlExec, fhd.mutation, fhd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fhd *FriendshipHistoryDelete) ExecX(ctx context.Context) int {
	n, err := fhd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fhd *FriendshipHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(friendshiphistory.Table, sqlgraph.NewFieldSpec(friendshiphistory.FieldID, field.TypeInt))
	if ps := fhd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fhd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fhd.mutation.done = true
	return affected, err
}

// FriendshipHistoryDeleteOne is the builder for deleting a single FriendshipHistory entity.
type FriendshipHistoryDeleteOne struct {
	fhd *FriendshipHistoryDelete
}

// Where appends a list predicates to the FriendshipHistoryDelete builder.
func (fhdo *FriendshipHistoryDeleteOne) Where(ps ...predicate.FriendshipHistory) *FriendshipHistoryDeleteOne {
	fhdo.fhd.mutation.Where(ps...)
	return fhdo
}

// Exec executes the deletion query.
func (fhdo *FriendshipHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := fhdo.fhd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{friendshiphistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fhdo *FriendshipHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := fhdo.Exec(ctx); err != nil {
		panic(err)
	}
}
