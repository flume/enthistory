// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/testexclude"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestExcludeCreate is the builder for creating a TestExclude entity.
type TestExcludeCreate struct {
	config
	mutation *TestExcludeMutation
	hooks    []Hook
}

// SetOtherID sets the "other_id" field.
func (tec *TestExcludeCreate) SetOtherID(u uuid.UUID) *TestExcludeCreate {
	tec.mutation.SetOtherID(u)
	return tec
}

// SetNillableOtherID sets the "other_id" field if the given value is not nil.
func (tec *TestExcludeCreate) SetNillableOtherID(u *uuid.UUID) *TestExcludeCreate {
	if u != nil {
		tec.SetOtherID(*u)
	}
	return tec
}

// SetName sets the "name" field.
func (tec *TestExcludeCreate) SetName(s string) *TestExcludeCreate {
	tec.mutation.SetName(s)
	return tec
}

// SetID sets the "id" field.
func (tec *TestExcludeCreate) SetID(u uuid.UUID) *TestExcludeCreate {
	tec.mutation.SetID(u)
	return tec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tec *TestExcludeCreate) SetNillableID(u *uuid.UUID) *TestExcludeCreate {
	if u != nil {
		tec.SetID(*u)
	}
	return tec
}

// Mutation returns the TestExcludeMutation object of the builder.
func (tec *TestExcludeCreate) Mutation() *TestExcludeMutation {
	return tec.mutation
}

// Save creates the TestExclude in the database.
func (tec *TestExcludeCreate) Save(ctx context.Context) (*TestExclude, error) {
	tec.defaults()
	return withHooks(ctx, tec.sqlSave, tec.mutation, tec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tec *TestExcludeCreate) SaveX(ctx context.Context) *TestExclude {
	v, err := tec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tec *TestExcludeCreate) Exec(ctx context.Context) error {
	_, err := tec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tec *TestExcludeCreate) ExecX(ctx context.Context) {
	if err := tec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tec *TestExcludeCreate) defaults() {
	if _, ok := tec.mutation.ID(); !ok {
		v := testexclude.DefaultID()
		tec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tec *TestExcludeCreate) check() error {
	if _, ok := tec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "TestExclude.name"`)}
	}
	if v, ok := tec.mutation.Name(); ok {
		if err := testexclude.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "TestExclude.name": %w`, err)}
		}
	}
	return nil
}

func (tec *TestExcludeCreate) sqlSave(ctx context.Context) (*TestExclude, error) {
	if err := tec.check(); err != nil {
		return nil, err
	}
	_node, _spec := tec.createSpec()
	if err := sqlgraph.CreateNode(ctx, tec.driver, _spec); err != nil {
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
	tec.mutation.id = &_node.ID
	tec.mutation.done = true
	return _node, nil
}

func (tec *TestExcludeCreate) createSpec() (*TestExclude, *sqlgraph.CreateSpec) {
	var (
		_node = &TestExclude{config: tec.config}
		_spec = sqlgraph.NewCreateSpec(testexclude.Table, sqlgraph.NewFieldSpec(testexclude.FieldID, field.TypeUUID))
	)
	if id, ok := tec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tec.mutation.OtherID(); ok {
		_spec.SetField(testexclude.FieldOtherID, field.TypeUUID, value)
		_node.OtherID = value
	}
	if value, ok := tec.mutation.Name(); ok {
		_spec.SetField(testexclude.FieldName, field.TypeString, value)
		_node.Name = value
	}
	return _node, _spec
}

// TestExcludeCreateBulk is the builder for creating many TestExclude entities in bulk.
type TestExcludeCreateBulk struct {
	config
	err      error
	builders []*TestExcludeCreate
}

// Save creates the TestExclude entities in the database.
func (tecb *TestExcludeCreateBulk) Save(ctx context.Context) ([]*TestExclude, error) {
	if tecb.err != nil {
		return nil, tecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tecb.builders))
	nodes := make([]*TestExclude, len(tecb.builders))
	mutators := make([]Mutator, len(tecb.builders))
	for i := range tecb.builders {
		func(i int, root context.Context) {
			builder := tecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestExcludeMutation)
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
					_, err = mutators[i+1].Mutate(root, tecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tecb *TestExcludeCreateBulk) SaveX(ctx context.Context) []*TestExclude {
	v, err := tecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tecb *TestExcludeCreateBulk) Exec(ctx context.Context) error {
	_, err := tecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tecb *TestExcludeCreateBulk) ExecX(ctx context.Context) {
	if err := tecb.Exec(ctx); err != nil {
		panic(err)
	}
}
