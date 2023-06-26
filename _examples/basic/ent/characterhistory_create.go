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
	"github.com/flume/enthistory/_examples/basic/ent/characterhistory"
)

// CharacterHistoryCreate is the builder for creating a CharacterHistory entity.
type CharacterHistoryCreate struct {
	config
	mutation *CharacterHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (chc *CharacterHistoryCreate) SetHistoryTime(t time.Time) *CharacterHistoryCreate {
	chc.mutation.SetHistoryTime(t)
	return chc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (chc *CharacterHistoryCreate) SetNillableHistoryTime(t *time.Time) *CharacterHistoryCreate {
	if t != nil {
		chc.SetHistoryTime(*t)
	}
	return chc
}

// SetRef sets the "ref" field.
func (chc *CharacterHistoryCreate) SetRef(i int) *CharacterHistoryCreate {
	chc.mutation.SetRef(i)
	return chc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (chc *CharacterHistoryCreate) SetNillableRef(i *int) *CharacterHistoryCreate {
	if i != nil {
		chc.SetRef(*i)
	}
	return chc
}

// SetOperation sets the "operation" field.
func (chc *CharacterHistoryCreate) SetOperation(et enthistory.OpType) *CharacterHistoryCreate {
	chc.mutation.SetOperation(et)
	return chc
}

// SetUpdatedBy sets the "updated_by" field.
func (chc *CharacterHistoryCreate) SetUpdatedBy(i int) *CharacterHistoryCreate {
	chc.mutation.SetUpdatedBy(i)
	return chc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (chc *CharacterHistoryCreate) SetNillableUpdatedBy(i *int) *CharacterHistoryCreate {
	if i != nil {
		chc.SetUpdatedBy(*i)
	}
	return chc
}

// SetCreatedAt sets the "created_at" field.
func (chc *CharacterHistoryCreate) SetCreatedAt(t time.Time) *CharacterHistoryCreate {
	chc.mutation.SetCreatedAt(t)
	return chc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (chc *CharacterHistoryCreate) SetNillableCreatedAt(t *time.Time) *CharacterHistoryCreate {
	if t != nil {
		chc.SetCreatedAt(*t)
	}
	return chc
}

// SetUpdatedAt sets the "updated_at" field.
func (chc *CharacterHistoryCreate) SetUpdatedAt(t time.Time) *CharacterHistoryCreate {
	chc.mutation.SetUpdatedAt(t)
	return chc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (chc *CharacterHistoryCreate) SetNillableUpdatedAt(t *time.Time) *CharacterHistoryCreate {
	if t != nil {
		chc.SetUpdatedAt(*t)
	}
	return chc
}

// SetAge sets the "age" field.
func (chc *CharacterHistoryCreate) SetAge(i int) *CharacterHistoryCreate {
	chc.mutation.SetAge(i)
	return chc
}

// SetName sets the "name" field.
func (chc *CharacterHistoryCreate) SetName(s string) *CharacterHistoryCreate {
	chc.mutation.SetName(s)
	return chc
}

// Mutation returns the CharacterHistoryMutation object of the builder.
func (chc *CharacterHistoryCreate) Mutation() *CharacterHistoryMutation {
	return chc.mutation
}

// Save creates the CharacterHistory in the database.
func (chc *CharacterHistoryCreate) Save(ctx context.Context) (*CharacterHistory, error) {
	chc.defaults()
	return withHooks(ctx, chc.sqlSave, chc.mutation, chc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (chc *CharacterHistoryCreate) SaveX(ctx context.Context) *CharacterHistory {
	v, err := chc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chc *CharacterHistoryCreate) Exec(ctx context.Context) error {
	_, err := chc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chc *CharacterHistoryCreate) ExecX(ctx context.Context) {
	if err := chc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (chc *CharacterHistoryCreate) defaults() {
	if _, ok := chc.mutation.HistoryTime(); !ok {
		v := characterhistory.DefaultHistoryTime()
		chc.mutation.SetHistoryTime(v)
	}
	if _, ok := chc.mutation.CreatedAt(); !ok {
		v := characterhistory.DefaultCreatedAt()
		chc.mutation.SetCreatedAt(v)
	}
	if _, ok := chc.mutation.UpdatedAt(); !ok {
		v := characterhistory.DefaultUpdatedAt()
		chc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chc *CharacterHistoryCreate) check() error {
	if _, ok := chc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`ent: missing required field "CharacterHistory.history_time"`)}
	}
	if _, ok := chc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`ent: missing required field "CharacterHistory.operation"`)}
	}
	if v, ok := chc.mutation.Operation(); ok {
		if err := characterhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`ent: validator failed for field "CharacterHistory.operation": %w`, err)}
		}
	}
	if _, ok := chc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CharacterHistory.created_at"`)}
	}
	if _, ok := chc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CharacterHistory.updated_at"`)}
	}
	if _, ok := chc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "CharacterHistory.age"`)}
	}
	if v, ok := chc.mutation.Age(); ok {
		if err := characterhistory.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "CharacterHistory.age": %w`, err)}
		}
	}
	if _, ok := chc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CharacterHistory.name"`)}
	}
	return nil
}

func (chc *CharacterHistoryCreate) sqlSave(ctx context.Context) (*CharacterHistory, error) {
	if err := chc.check(); err != nil {
		return nil, err
	}
	_node, _spec := chc.createSpec()
	if err := sqlgraph.CreateNode(ctx, chc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	chc.mutation.id = &_node.ID
	chc.mutation.done = true
	return _node, nil
}

func (chc *CharacterHistoryCreate) createSpec() (*CharacterHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &CharacterHistory{config: chc.config}
		_spec = sqlgraph.NewCreateSpec(characterhistory.Table, sqlgraph.NewFieldSpec(characterhistory.FieldID, field.TypeInt))
	)
	if value, ok := chc.mutation.HistoryTime(); ok {
		_spec.SetField(characterhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := chc.mutation.Ref(); ok {
		_spec.SetField(characterhistory.FieldRef, field.TypeInt, value)
		_node.Ref = value
	}
	if value, ok := chc.mutation.Operation(); ok {
		_spec.SetField(characterhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := chc.mutation.UpdatedBy(); ok {
		_spec.SetField(characterhistory.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = &value
	}
	if value, ok := chc.mutation.CreatedAt(); ok {
		_spec.SetField(characterhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := chc.mutation.UpdatedAt(); ok {
		_spec.SetField(characterhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := chc.mutation.Age(); ok {
		_spec.SetField(characterhistory.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := chc.mutation.Name(); ok {
		_spec.SetField(characterhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// CharacterHistoryCreateBulk is the builder for creating many CharacterHistory entities in bulk.
type CharacterHistoryCreateBulk struct {
	config
	builders []*CharacterHistoryCreate
}

// Save creates the CharacterHistory entities in the database.
func (chcb *CharacterHistoryCreateBulk) Save(ctx context.Context) ([]*CharacterHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(chcb.builders))
	nodes := make([]*CharacterHistory, len(chcb.builders))
	mutators := make([]Mutator, len(chcb.builders))
	for i := range chcb.builders {
		func(i int, root context.Context) {
			builder := chcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, chcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, chcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, chcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (chcb *CharacterHistoryCreateBulk) SaveX(ctx context.Context) []*CharacterHistory {
	v, err := chcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chcb *CharacterHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := chcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcb *CharacterHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := chcb.Exec(ctx); err != nil {
		panic(err)
	}
}
