// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/flume/enthistory"
	"github.com/flume/enthistory/_examples/basic/ent/friendshiphistory"
)

// FriendshipHistoryCreate is the builder for creating a FriendshipHistory entity.
type FriendshipHistoryCreate struct {
	config
	mutation *FriendshipHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (fhc *FriendshipHistoryCreate) SetHistoryTime(t time.Time) *FriendshipHistoryCreate {
	fhc.mutation.SetHistoryTime(t)
	return fhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (fhc *FriendshipHistoryCreate) SetNillableHistoryTime(t *time.Time) *FriendshipHistoryCreate {
	if t != nil {
		fhc.SetHistoryTime(*t)
	}
	return fhc
}

// SetRef sets the "ref" field.
func (fhc *FriendshipHistoryCreate) SetRef(i int) *FriendshipHistoryCreate {
	fhc.mutation.SetRef(i)
	return fhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (fhc *FriendshipHistoryCreate) SetNillableRef(i *int) *FriendshipHistoryCreate {
	if i != nil {
		fhc.SetRef(*i)
	}
	return fhc
}

// SetOperation sets the "operation" field.
func (fhc *FriendshipHistoryCreate) SetOperation(et enthistory.OpType) *FriendshipHistoryCreate {
	fhc.mutation.SetOperation(et)
	return fhc
}

// SetUpdatedBy sets the "updated_by" field.
func (fhc *FriendshipHistoryCreate) SetUpdatedBy(i int) *FriendshipHistoryCreate {
	fhc.mutation.SetUpdatedBy(i)
	return fhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fhc *FriendshipHistoryCreate) SetNillableUpdatedBy(i *int) *FriendshipHistoryCreate {
	if i != nil {
		fhc.SetUpdatedBy(*i)
	}
	return fhc
}

// SetCreatedAt sets the "created_at" field.
func (fhc *FriendshipHistoryCreate) SetCreatedAt(t time.Time) *FriendshipHistoryCreate {
	fhc.mutation.SetCreatedAt(t)
	return fhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fhc *FriendshipHistoryCreate) SetNillableCreatedAt(t *time.Time) *FriendshipHistoryCreate {
	if t != nil {
		fhc.SetCreatedAt(*t)
	}
	return fhc
}

// SetUpdatedAt sets the "updated_at" field.
func (fhc *FriendshipHistoryCreate) SetUpdatedAt(t time.Time) *FriendshipHistoryCreate {
	fhc.mutation.SetUpdatedAt(t)
	return fhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fhc *FriendshipHistoryCreate) SetNillableUpdatedAt(t *time.Time) *FriendshipHistoryCreate {
	if t != nil {
		fhc.SetUpdatedAt(*t)
	}
	return fhc
}

// SetCharacterID sets the "character_id" field.
func (fhc *FriendshipHistoryCreate) SetCharacterID(i int) *FriendshipHistoryCreate {
	fhc.mutation.SetCharacterID(i)
	return fhc
}

// SetFriendID sets the "friend_id" field.
func (fhc *FriendshipHistoryCreate) SetFriendID(i int) *FriendshipHistoryCreate {
	fhc.mutation.SetFriendID(i)
	return fhc
}

// Mutation returns the FriendshipHistoryMutation object of the builder.
func (fhc *FriendshipHistoryCreate) Mutation() *FriendshipHistoryMutation {
	return fhc.mutation
}

// Save creates the FriendshipHistory in the database.
func (fhc *FriendshipHistoryCreate) Save(ctx context.Context) (*FriendshipHistory, error) {
	fhc.defaults()
	return withHooks(ctx, fhc.sqlSave, fhc.mutation, fhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fhc *FriendshipHistoryCreate) SaveX(ctx context.Context) *FriendshipHistory {
	v, err := fhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fhc *FriendshipHistoryCreate) Exec(ctx context.Context) error {
	_, err := fhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhc *FriendshipHistoryCreate) ExecX(ctx context.Context) {
	if err := fhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fhc *FriendshipHistoryCreate) defaults() {
	if _, ok := fhc.mutation.HistoryTime(); !ok {
		v := friendshiphistory.DefaultHistoryTime()
		fhc.mutation.SetHistoryTime(v)
	}
	if _, ok := fhc.mutation.CreatedAt(); !ok {
		v := friendshiphistory.DefaultCreatedAt()
		fhc.mutation.SetCreatedAt(v)
	}
	if _, ok := fhc.mutation.UpdatedAt(); !ok {
		v := friendshiphistory.DefaultUpdatedAt()
		fhc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fhc *FriendshipHistoryCreate) check() error {
	if _, ok := fhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "FriendshipHistory.history_time"`)}
	}
	if _, ok := fhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "FriendshipHistory.operation"`)}
	}
	if v, ok := fhc.mutation.Operation(); ok {
		if err := friendshiphistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "FriendshipHistory.operation": %w`, err)}
		}
	}
	if _, ok := fhc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FriendshipHistory.created_at"`)}
	}
	if _, ok := fhc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FriendshipHistory.updated_at"`)}
	}
	if _, ok := fhc.mutation.CharacterID(); !ok {
		return &ValidationError{Name: "character_id", err: errors.New(`ent: missing required field "FriendshipHistory.character_id"`)}
	}
	if _, ok := fhc.mutation.FriendID(); !ok {
		return &ValidationError{Name: "friend_id", err: errors.New(`ent: missing required field "FriendshipHistory.friend_id"`)}
	}
	return nil
}

func (fhc *FriendshipHistoryCreate) sqlSave(ctx context.Context) (*FriendshipHistory, error) {
	if err := fhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fhc.mutation.id = &_node.ID
	fhc.mutation.done = true
	return _node, nil
}

func (fhc *FriendshipHistoryCreate) createSpec() (*FriendshipHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &FriendshipHistory{config: fhc.config}
		_spec = sqlgraph.NewCreateSpec(friendshiphistory.Table, sqlgraph.NewFieldSpec(friendshiphistory.FieldID, field.TypeInt))
	)
	if value, ok := fhc.mutation.HistoryTime(); ok {
		_spec.SetField(friendshiphistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := fhc.mutation.Ref(); ok {
		_spec.SetField(friendshiphistory.FieldRef, field.TypeInt, value)
		_node.Ref = value
	}
	if value, ok := fhc.mutation.Operation(); ok {
		_spec.SetField(friendshiphistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := fhc.mutation.UpdatedBy(); ok {
		_spec.SetField(friendshiphistory.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = &value
	}
	if value, ok := fhc.mutation.CreatedAt(); ok {
		_spec.SetField(friendshiphistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fhc.mutation.UpdatedAt(); ok {
		_spec.SetField(friendshiphistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fhc.mutation.CharacterID(); ok {
		_spec.SetField(friendshiphistory.FieldCharacterID, field.TypeInt, value)
		_node.CharacterID = value
	}
	if value, ok := fhc.mutation.FriendID(); ok {
		_spec.SetField(friendshiphistory.FieldFriendID, field.TypeInt, value)
		_node.FriendID = value
	}
	return _node, _spec
}

// FriendshipHistoryCreateBulk is the builder for creating many FriendshipHistory entities in bulk.
type FriendshipHistoryCreateBulk struct {
	config
	builders []*FriendshipHistoryCreate
}

// Save creates the FriendshipHistory entities in the database.
func (fhcb *FriendshipHistoryCreateBulk) Save(ctx context.Context) ([]*FriendshipHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fhcb.builders))
	nodes := make([]*FriendshipHistory, len(fhcb.builders))
	mutators := make([]Mutator, len(fhcb.builders))
	for i := range fhcb.builders {
		func(i int, root context.Context) {
			builder := fhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FriendshipHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fhcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fhcb *FriendshipHistoryCreateBulk) SaveX(ctx context.Context) []*FriendshipHistory {
	v, err := fhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fhcb *FriendshipHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := fhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhcb *FriendshipHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := fhcb.Exec(ctx); err != nil {
		panic(err)
	}
}
