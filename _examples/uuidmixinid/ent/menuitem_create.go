// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/uuidmixinid/ent/menuitem"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MenuItemCreate is the builder for creating a MenuItem entity.
type MenuItemCreate struct {
	config
	mutation *MenuItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mic *MenuItemCreate) SetCreatedAt(t time.Time) *MenuItemCreate {
	mic.mutation.SetCreatedAt(t)
	return mic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mic *MenuItemCreate) SetNillableCreatedAt(t *time.Time) *MenuItemCreate {
	if t != nil {
		mic.SetCreatedAt(*t)
	}
	return mic
}

// SetUpdatedAt sets the "updated_at" field.
func (mic *MenuItemCreate) SetUpdatedAt(t time.Time) *MenuItemCreate {
	mic.mutation.SetUpdatedAt(t)
	return mic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mic *MenuItemCreate) SetNillableUpdatedAt(t *time.Time) *MenuItemCreate {
	if t != nil {
		mic.SetUpdatedAt(*t)
	}
	return mic
}

// SetName sets the "name" field.
func (mic *MenuItemCreate) SetName(s string) *MenuItemCreate {
	mic.mutation.SetName(s)
	return mic
}

// SetPrice sets the "price" field.
func (mic *MenuItemCreate) SetPrice(f float64) *MenuItemCreate {
	mic.mutation.SetPrice(f)
	return mic
}

// SetDescription sets the "description" field.
func (mic *MenuItemCreate) SetDescription(s string) *MenuItemCreate {
	mic.mutation.SetDescription(s)
	return mic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mic *MenuItemCreate) SetNillableDescription(s *string) *MenuItemCreate {
	if s != nil {
		mic.SetDescription(*s)
	}
	return mic
}

// SetID sets the "id" field.
func (mic *MenuItemCreate) SetID(u uuid.UUID) *MenuItemCreate {
	mic.mutation.SetID(u)
	return mic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mic *MenuItemCreate) SetNillableID(u *uuid.UUID) *MenuItemCreate {
	if u != nil {
		mic.SetID(*u)
	}
	return mic
}

// Mutation returns the MenuItemMutation object of the builder.
func (mic *MenuItemCreate) Mutation() *MenuItemMutation {
	return mic.mutation
}

// Save creates the MenuItem in the database.
func (mic *MenuItemCreate) Save(ctx context.Context) (*MenuItem, error) {
	mic.defaults()
	return withHooks(ctx, mic.sqlSave, mic.mutation, mic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MenuItemCreate) SaveX(ctx context.Context) *MenuItem {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MenuItemCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MenuItemCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mic *MenuItemCreate) defaults() {
	if _, ok := mic.mutation.CreatedAt(); !ok {
		v := menuitem.DefaultCreatedAt()
		mic.mutation.SetCreatedAt(v)
	}
	if _, ok := mic.mutation.UpdatedAt(); !ok {
		v := menuitem.DefaultUpdatedAt()
		mic.mutation.SetUpdatedAt(v)
	}
	if _, ok := mic.mutation.ID(); !ok {
		v := menuitem.DefaultID()
		mic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MenuItemCreate) check() error {
	if _, ok := mic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MenuItem.created_at"`)}
	}
	if _, ok := mic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MenuItem.updated_at"`)}
	}
	if _, ok := mic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MenuItem.name"`)}
	}
	if v, ok := mic.mutation.Name(); ok {
		if err := menuitem.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MenuItem.name": %w`, err)}
		}
	}
	if _, ok := mic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "MenuItem.price"`)}
	}
	if v, ok := mic.mutation.Price(); ok {
		if err := menuitem.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "MenuItem.price": %w`, err)}
		}
	}
	return nil
}

func (mic *MenuItemCreate) sqlSave(ctx context.Context) (*MenuItem, error) {
	if err := mic.check(); err != nil {
		return nil, err
	}
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mic.mutation.id = &_node.ID
	mic.mutation.done = true
	return _node, nil
}

func (mic *MenuItemCreate) createSpec() (*MenuItem, *sqlgraph.CreateSpec) {
	var (
		_node = &MenuItem{config: mic.config}
		_spec = sqlgraph.NewCreateSpec(menuitem.Table, sqlgraph.NewFieldSpec(menuitem.FieldID, field.TypeUUID))
	)
	if id, ok := mic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mic.mutation.CreatedAt(); ok {
		_spec.SetField(menuitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mic.mutation.UpdatedAt(); ok {
		_spec.SetField(menuitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mic.mutation.Name(); ok {
		_spec.SetField(menuitem.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mic.mutation.Price(); ok {
		_spec.SetField(menuitem.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := mic.mutation.Description(); ok {
		_spec.SetField(menuitem.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// MenuItemCreateBulk is the builder for creating many MenuItem entities in bulk.
type MenuItemCreateBulk struct {
	config
	err      error
	builders []*MenuItemCreate
}

// Save creates the MenuItem entities in the database.
func (micb *MenuItemCreateBulk) Save(ctx context.Context) ([]*MenuItem, error) {
	if micb.err != nil {
		return nil, micb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MenuItem, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuItemMutation)
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
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MenuItemCreateBulk) SaveX(ctx context.Context) []*MenuItem {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MenuItemCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MenuItemCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}
