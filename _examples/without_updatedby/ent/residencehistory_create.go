// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/without_updatedby/ent/residencehistory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// ResidenceHistoryCreate is the builder for creating a ResidenceHistory entity.
type ResidenceHistoryCreate struct {
	config
	mutation *ResidenceHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (rhc *ResidenceHistoryCreate) SetHistoryTime(t time.Time) *ResidenceHistoryCreate {
	rhc.mutation.SetHistoryTime(t)
	return rhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (rhc *ResidenceHistoryCreate) SetNillableHistoryTime(t *time.Time) *ResidenceHistoryCreate {
	if t != nil {
		rhc.SetHistoryTime(*t)
	}
	return rhc
}

// SetOperation sets the "operation" field.
func (rhc *ResidenceHistoryCreate) SetOperation(et enthistory.OpType) *ResidenceHistoryCreate {
	rhc.mutation.SetOperation(et)
	return rhc
}

// SetRef sets the "ref" field.
func (rhc *ResidenceHistoryCreate) SetRef(u uuid.UUID) *ResidenceHistoryCreate {
	rhc.mutation.SetRef(u)
	return rhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (rhc *ResidenceHistoryCreate) SetNillableRef(u *uuid.UUID) *ResidenceHistoryCreate {
	if u != nil {
		rhc.SetRef(*u)
	}
	return rhc
}

// SetName sets the "name" field.
func (rhc *ResidenceHistoryCreate) SetName(s string) *ResidenceHistoryCreate {
	rhc.mutation.SetName(s)
	return rhc
}

// SetCreatedAt sets the "created_at" field.
func (rhc *ResidenceHistoryCreate) SetCreatedAt(t time.Time) *ResidenceHistoryCreate {
	rhc.mutation.SetCreatedAt(t)
	return rhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rhc *ResidenceHistoryCreate) SetNillableCreatedAt(t *time.Time) *ResidenceHistoryCreate {
	if t != nil {
		rhc.SetCreatedAt(*t)
	}
	return rhc
}

// SetUpdatedAt sets the "updated_at" field.
func (rhc *ResidenceHistoryCreate) SetUpdatedAt(t time.Time) *ResidenceHistoryCreate {
	rhc.mutation.SetUpdatedAt(t)
	return rhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rhc *ResidenceHistoryCreate) SetNillableUpdatedAt(t *time.Time) *ResidenceHistoryCreate {
	if t != nil {
		rhc.SetUpdatedAt(*t)
	}
	return rhc
}

// SetID sets the "id" field.
func (rhc *ResidenceHistoryCreate) SetID(i int) *ResidenceHistoryCreate {
	rhc.mutation.SetID(i)
	return rhc
}

// Mutation returns the ResidenceHistoryMutation object of the builder.
func (rhc *ResidenceHistoryCreate) Mutation() *ResidenceHistoryMutation {
	return rhc.mutation
}

// Save creates the ResidenceHistory in the database.
func (rhc *ResidenceHistoryCreate) Save(ctx context.Context) (*ResidenceHistory, error) {
	rhc.defaults()
	return withHooks(ctx, rhc.sqlSave, rhc.mutation, rhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rhc *ResidenceHistoryCreate) SaveX(ctx context.Context) *ResidenceHistory {
	v, err := rhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rhc *ResidenceHistoryCreate) Exec(ctx context.Context) error {
	_, err := rhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rhc *ResidenceHistoryCreate) ExecX(ctx context.Context) {
	if err := rhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rhc *ResidenceHistoryCreate) defaults() {
	if _, ok := rhc.mutation.HistoryTime(); !ok {
		v := residencehistory.DefaultHistoryTime()
		rhc.mutation.SetHistoryTime(v)
	}
	if _, ok := rhc.mutation.CreatedAt(); !ok {
		v := residencehistory.DefaultCreatedAt()
		rhc.mutation.SetCreatedAt(v)
	}
	if _, ok := rhc.mutation.UpdatedAt(); !ok {
		v := residencehistory.DefaultUpdatedAt()
		rhc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rhc *ResidenceHistoryCreate) check() error {
	if _, ok := rhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "ResidenceHistory.history_time"`)}
	}
	if _, ok := rhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "ResidenceHistory.operation"`)}
	}
	if v, ok := rhc.mutation.Operation(); ok {
		if err := residencehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "ResidenceHistory.operation": %w`, err)}
		}
	}
	if _, ok := rhc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ResidenceHistory.name"`)}
	}
	if _, ok := rhc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ResidenceHistory.created_at"`)}
	}
	if _, ok := rhc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ResidenceHistory.updated_at"`)}
	}
	return nil
}

func (rhc *ResidenceHistoryCreate) sqlSave(ctx context.Context) (*ResidenceHistory, error) {
	if err := rhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	rhc.mutation.id = &_node.ID
	rhc.mutation.done = true
	return _node, nil
}

func (rhc *ResidenceHistoryCreate) createSpec() (*ResidenceHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ResidenceHistory{config: rhc.config}
		_spec = sqlgraph.NewCreateSpec(residencehistory.Table, sqlgraph.NewFieldSpec(residencehistory.FieldID, field.TypeInt))
	)
	if id, ok := rhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rhc.mutation.HistoryTime(); ok {
		_spec.SetField(residencehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := rhc.mutation.Operation(); ok {
		_spec.SetField(residencehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := rhc.mutation.Ref(); ok {
		_spec.SetField(residencehistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := rhc.mutation.Name(); ok {
		_spec.SetField(residencehistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rhc.mutation.CreatedAt(); ok {
		_spec.SetField(residencehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rhc.mutation.UpdatedAt(); ok {
		_spec.SetField(residencehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// ResidenceHistoryCreateBulk is the builder for creating many ResidenceHistory entities in bulk.
type ResidenceHistoryCreateBulk struct {
	config
	err      error
	builders []*ResidenceHistoryCreate
}

// Save creates the ResidenceHistory entities in the database.
func (rhcb *ResidenceHistoryCreateBulk) Save(ctx context.Context) ([]*ResidenceHistory, error) {
	if rhcb.err != nil {
		return nil, rhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rhcb.builders))
	nodes := make([]*ResidenceHistory, len(rhcb.builders))
	mutators := make([]Mutator, len(rhcb.builders))
	for i := range rhcb.builders {
		func(i int, root context.Context) {
			builder := rhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResidenceHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, rhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rhcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, rhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rhcb *ResidenceHistoryCreateBulk) SaveX(ctx context.Context) []*ResidenceHistory {
	v, err := rhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rhcb *ResidenceHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := rhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rhcb *ResidenceHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := rhcb.Exec(ctx); err != nil {
		panic(err)
	}
}
