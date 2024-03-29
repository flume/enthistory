// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/organizationhistory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/flume/enthistory"
)

// OrganizationHistoryCreate is the builder for creating a OrganizationHistory entity.
type OrganizationHistoryCreate struct {
	config
	mutation *OrganizationHistoryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ohc *OrganizationHistoryCreate) SetCreatedAt(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetCreatedAt(t)
	return ohc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableCreatedAt(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetCreatedAt(*t)
	}
	return ohc
}

// SetUpdatedAt sets the "updated_at" field.
func (ohc *OrganizationHistoryCreate) SetUpdatedAt(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetUpdatedAt(t)
	return ohc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetUpdatedAt(*t)
	}
	return ohc
}

// SetHistoryTime sets the "history_time" field.
func (ohc *OrganizationHistoryCreate) SetHistoryTime(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetHistoryTime(t)
	return ohc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableHistoryTime(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetHistoryTime(*t)
	}
	return ohc
}

// SetOperation sets the "operation" field.
func (ohc *OrganizationHistoryCreate) SetOperation(et enthistory.OpType) *OrganizationHistoryCreate {
	ohc.mutation.SetOperation(et)
	return ohc
}

// SetRef sets the "ref" field.
func (ohc *OrganizationHistoryCreate) SetRef(u uuid.UUID) *OrganizationHistoryCreate {
	ohc.mutation.SetRef(u)
	return ohc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableRef(u *uuid.UUID) *OrganizationHistoryCreate {
	if u != nil {
		ohc.SetRef(*u)
	}
	return ohc
}

// SetUpdatedBy sets the "updated_by" field.
func (ohc *OrganizationHistoryCreate) SetUpdatedBy(u uuid.UUID) *OrganizationHistoryCreate {
	ohc.mutation.SetUpdatedBy(u)
	return ohc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableUpdatedBy(u *uuid.UUID) *OrganizationHistoryCreate {
	if u != nil {
		ohc.SetUpdatedBy(*u)
	}
	return ohc
}

// SetName sets the "name" field.
func (ohc *OrganizationHistoryCreate) SetName(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetName(s)
	return ohc
}

// SetInfo sets the "info" field.
func (ohc *OrganizationHistoryCreate) SetInfo(m map[string]interface{}) *OrganizationHistoryCreate {
	ohc.mutation.SetInfo(m)
	return ohc
}

// SetID sets the "id" field.
func (ohc *OrganizationHistoryCreate) SetID(i int) *OrganizationHistoryCreate {
	ohc.mutation.SetID(i)
	return ohc
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohc *OrganizationHistoryCreate) Mutation() *OrganizationHistoryMutation {
	return ohc.mutation
}

// Save creates the OrganizationHistory in the database.
func (ohc *OrganizationHistoryCreate) Save(ctx context.Context) (*OrganizationHistory, error) {
	ohc.defaults()
	return withHooks(ctx, ohc.sqlSave, ohc.mutation, ohc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ohc *OrganizationHistoryCreate) SaveX(ctx context.Context) *OrganizationHistory {
	v, err := ohc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohc *OrganizationHistoryCreate) Exec(ctx context.Context) error {
	_, err := ohc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohc *OrganizationHistoryCreate) ExecX(ctx context.Context) {
	if err := ohc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ohc *OrganizationHistoryCreate) defaults() {
	if _, ok := ohc.mutation.CreatedAt(); !ok {
		v := organizationhistory.DefaultCreatedAt()
		ohc.mutation.SetCreatedAt(v)
	}
	if _, ok := ohc.mutation.UpdatedAt(); !ok {
		v := organizationhistory.DefaultUpdatedAt()
		ohc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ohc.mutation.HistoryTime(); !ok {
		v := organizationhistory.DefaultHistoryTime()
		ohc.mutation.SetHistoryTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ohc *OrganizationHistoryCreate) check() error {
	if _, ok := ohc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrganizationHistory.created_at"`)}
	}
	if _, ok := ohc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrganizationHistory.updated_at"`)}
	}
	if _, ok := ohc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "OrganizationHistory.history_time"`)}
	}
	if _, ok := ohc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "OrganizationHistory.operation"`)}
	}
	if v, ok := ohc.mutation.Operation(); ok {
		if err := organizationhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "OrganizationHistory.operation": %w`, err)}
		}
	}
	if _, ok := ohc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrganizationHistory.name"`)}
	}
	return nil
}

func (ohc *OrganizationHistoryCreate) sqlSave(ctx context.Context) (*OrganizationHistory, error) {
	if err := ohc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ohc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ohc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ohc.mutation.id = &_node.ID
	ohc.mutation.done = true
	return _node, nil
}

func (ohc *OrganizationHistoryCreate) createSpec() (*OrganizationHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationHistory{config: ohc.config}
		_spec = sqlgraph.NewCreateSpec(organizationhistory.Table, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeInt))
	)
	if id, ok := ohc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ohc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ohc.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ohc.mutation.HistoryTime(); ok {
		_spec.SetField(organizationhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ohc.mutation.Operation(); ok {
		_spec.SetField(organizationhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ohc.mutation.Ref(); ok {
		_spec.SetField(organizationhistory.FieldRef, field.TypeUUID, value)
		_node.Ref = value
	}
	if value, ok := ohc.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = &value
	}
	if value, ok := ohc.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ohc.mutation.Info(); ok {
		_spec.SetField(organizationhistory.FieldInfo, field.TypeJSON, value)
		_node.Info = value
	}
	return _node, _spec
}

// OrganizationHistoryCreateBulk is the builder for creating many OrganizationHistory entities in bulk.
type OrganizationHistoryCreateBulk struct {
	config
	err      error
	builders []*OrganizationHistoryCreate
}

// Save creates the OrganizationHistory entities in the database.
func (ohcb *OrganizationHistoryCreateBulk) Save(ctx context.Context) ([]*OrganizationHistory, error) {
	if ohcb.err != nil {
		return nil, ohcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ohcb.builders))
	nodes := make([]*OrganizationHistory, len(ohcb.builders))
	mutators := make([]Mutator, len(ohcb.builders))
	for i := range ohcb.builders {
		func(i int, root context.Context) {
			builder := ohcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ohcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ohcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ohcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ohcb *OrganizationHistoryCreateBulk) SaveX(ctx context.Context) []*OrganizationHistory {
	v, err := ohcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohcb *OrganizationHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ohcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohcb *OrganizationHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ohcb.Exec(ctx); err != nil {
		panic(err)
	}
}
