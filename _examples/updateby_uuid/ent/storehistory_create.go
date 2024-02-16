// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/storehistory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// StoreHistoryCreate is the builder for creating a StoreHistory entity.
type StoreHistoryCreate struct {
	config
	mutation *StoreHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (shc *StoreHistoryCreate) SetHistoryTime(t time.Time) *StoreHistoryCreate {
	shc.mutation.SetHistoryTime(t)
	return shc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (shc *StoreHistoryCreate) SetNillableHistoryTime(t *time.Time) *StoreHistoryCreate {
	if t != nil {
		shc.SetHistoryTime(*t)
	}
	return shc
}

// SetOperation sets the "operation" field.
func (shc *StoreHistoryCreate) SetOperation(et enthistory.OpType) *StoreHistoryCreate {
	shc.mutation.SetOperation(et)
	return shc
}

// SetRef sets the "ref" field.
func (shc *StoreHistoryCreate) SetRef(u uuid.UUID) *StoreHistoryCreate {
	shc.mutation.SetRef(u)
	return shc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (shc *StoreHistoryCreate) SetNillableRef(u *uuid.UUID) *StoreHistoryCreate {
	if u != nil {
		shc.SetRef(*u)
	}
	return shc
}

// SetUpdatedBy sets the "updated_by" field.
func (shc *StoreHistoryCreate) SetUpdatedBy(u uuid.UUID) *StoreHistoryCreate {
	shc.mutation.SetUpdatedBy(u)
	return shc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (shc *StoreHistoryCreate) SetNillableUpdatedBy(u *uuid.UUID) *StoreHistoryCreate {
	if u != nil {
		shc.SetUpdatedBy(*u)
	}
	return shc
}

// SetCreatedAt sets the "created_at" field.
func (shc *StoreHistoryCreate) SetCreatedAt(t time.Time) *StoreHistoryCreate {
	shc.mutation.SetCreatedAt(t)
	return shc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (shc *StoreHistoryCreate) SetNillableCreatedAt(t *time.Time) *StoreHistoryCreate {
	if t != nil {
		shc.SetCreatedAt(*t)
	}
	return shc
}

// SetUpdatedAt sets the "updated_at" field.
func (shc *StoreHistoryCreate) SetUpdatedAt(t time.Time) *StoreHistoryCreate {
	shc.mutation.SetUpdatedAt(t)
	return shc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (shc *StoreHistoryCreate) SetNillableUpdatedAt(t *time.Time) *StoreHistoryCreate {
	if t != nil {
		shc.SetUpdatedAt(*t)
	}
	return shc
}

// SetName sets the "name" field.
func (shc *StoreHistoryCreate) SetName(s string) *StoreHistoryCreate {
	shc.mutation.SetName(s)
	return shc
}

// SetRegion sets the "region" field.
func (shc *StoreHistoryCreate) SetRegion(s string) *StoreHistoryCreate {
	shc.mutation.SetRegion(s)
	return shc
}

// SetOrganizationID sets the "organization_id" field.
func (shc *StoreHistoryCreate) SetOrganizationID(u uuid.UUID) *StoreHistoryCreate {
	shc.mutation.SetOrganizationID(u)
	return shc
}

// SetID sets the "id" field.
func (shc *StoreHistoryCreate) SetID(i int) *StoreHistoryCreate {
	shc.mutation.SetID(i)
	return shc
}

// Mutation returns the StoreHistoryMutation object of the builder.
func (shc *StoreHistoryCreate) Mutation() *StoreHistoryMutation {
	return shc.mutation
}

// Save creates the StoreHistory in the database.
func (shc *StoreHistoryCreate) Save(ctx context.Context) (*StoreHistory, error) {
	shc.defaults()
	return withHooks(ctx, shc.sqlSave, shc.mutation, shc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (shc *StoreHistoryCreate) SaveX(ctx context.Context) *StoreHistory {
	v, err := shc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (shc *StoreHistoryCreate) Exec(ctx context.Context) error {
	_, err := shc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shc *StoreHistoryCreate) ExecX(ctx context.Context) {
	if err := shc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shc *StoreHistoryCreate) defaults() {
	if _, ok := shc.mutation.HistoryTime(); !ok {
		v := storehistory.DefaultHistoryTime()
		shc.mutation.SetHistoryTime(v)
	}
	if _, ok := shc.mutation.CreatedAt(); !ok {
		v := storehistory.DefaultCreatedAt()
		shc.mutation.SetCreatedAt(v)
	}
	if _, ok := shc.mutation.UpdatedAt(); !ok {
		v := storehistory.DefaultUpdatedAt()
		shc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (shc *StoreHistoryCreate) check() error {
	if _, ok := shc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "StoreHistory.history_time"`)}
	}
	if _, ok := shc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "StoreHistory.operation"`)}
	}
	if v, ok := shc.mutation.Operation(); ok {
		if err := storehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "StoreHistory.operation": %w`, err)}
		}
	}
	if _, ok := shc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StoreHistory.created_at"`)}
	}
	if _, ok := shc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StoreHistory.updated_at"`)}
	}
	if _, ok := shc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "StoreHistory.name"`)}
	}
	if _, ok := shc.mutation.Region(); !ok {
		return &ValidationError{Name: "region", err: errors.New(`ent: missing required field "StoreHistory.region"`)}
	}
	if _, ok := shc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "StoreHistory.organization_id"`)}
	}
	return nil
}

func (shc *StoreHistoryCreate) sqlSave(ctx context.Context) (*StoreHistory, error) {
	if err := shc.check(); err != nil {
		return nil, err
	}
	_node, _spec := shc.createSpec()
	if err := sqlgraph.CreateNode(ctx, shc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	shc.mutation.id = &_node.ID
	shc.mutation.done = true
	return _node, nil
}

func (shc *StoreHistoryCreate) createSpec() (*StoreHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &StoreHistory{config: shc.config}
		_spec = sqlgraph.NewCreateSpec(storehistory.Table, sqlgraph.NewFieldSpec(storehistory.FieldID, field.TypeInt))
	)
	if id, ok := shc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := shc.mutation.HistoryTime(); ok {
		_spec.SetField(storehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := shc.mutation.Operation(); ok {
		_spec.SetField(storehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := shc.mutation.Ref(); ok {
		_spec.SetField(storehistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := shc.mutation.UpdatedBy(); ok {
		_spec.SetField(storehistory.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = &value
	}
	if value, ok := shc.mutation.CreatedAt(); ok {
		_spec.SetField(storehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := shc.mutation.UpdatedAt(); ok {
		_spec.SetField(storehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := shc.mutation.Name(); ok {
		_spec.SetField(storehistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := shc.mutation.Region(); ok {
		_spec.SetField(storehistory.FieldRegion, field.TypeString, value)
		_node.Region = value
	}
	if value, ok := shc.mutation.OrganizationID(); ok {
		_spec.SetField(storehistory.FieldOrganizationID, field.TypeUUID, value)
		_node.OrganizationID = value
	}
	return _node, _spec
}

// StoreHistoryCreateBulk is the builder for creating many StoreHistory entities in bulk.
type StoreHistoryCreateBulk struct {
	config
	err      error
	builders []*StoreHistoryCreate
}

// Save creates the StoreHistory entities in the database.
func (shcb *StoreHistoryCreateBulk) Save(ctx context.Context) ([]*StoreHistory, error) {
	if shcb.err != nil {
		return nil, shcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(shcb.builders))
	nodes := make([]*StoreHistory, len(shcb.builders))
	mutators := make([]Mutator, len(shcb.builders))
	for i := range shcb.builders {
		func(i int, root context.Context) {
			builder := shcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StoreHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, shcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, shcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, shcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (shcb *StoreHistoryCreateBulk) SaveX(ctx context.Context) []*StoreHistory {
	v, err := shcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (shcb *StoreHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := shcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shcb *StoreHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := shcb.Exec(ctx); err != nil {
		panic(err)
	}
}
