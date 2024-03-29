// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/without_updatedby/ent/character"
	"_examples/without_updatedby/ent/residence"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ResidenceCreate is the builder for creating a Residence entity.
type ResidenceCreate struct {
	config
	mutation *ResidenceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (rc *ResidenceCreate) SetCreatedAt(t time.Time) *ResidenceCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *ResidenceCreate) SetNillableCreatedAt(t *time.Time) *ResidenceCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ResidenceCreate) SetUpdatedAt(t time.Time) *ResidenceCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *ResidenceCreate) SetNillableUpdatedAt(t *time.Time) *ResidenceCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetName sets the "name" field.
func (rc *ResidenceCreate) SetName(s string) *ResidenceCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetID sets the "id" field.
func (rc *ResidenceCreate) SetID(u uuid.UUID) *ResidenceCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ResidenceCreate) SetNillableID(u *uuid.UUID) *ResidenceCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// AddOccupantIDs adds the "occupants" edge to the Character entity by IDs.
func (rc *ResidenceCreate) AddOccupantIDs(ids ...int) *ResidenceCreate {
	rc.mutation.AddOccupantIDs(ids...)
	return rc
}

// AddOccupants adds the "occupants" edges to the Character entity.
func (rc *ResidenceCreate) AddOccupants(c ...*Character) *ResidenceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rc.AddOccupantIDs(ids...)
}

// Mutation returns the ResidenceMutation object of the builder.
func (rc *ResidenceCreate) Mutation() *ResidenceMutation {
	return rc.mutation
}

// Save creates the Residence in the database.
func (rc *ResidenceCreate) Save(ctx context.Context) (*Residence, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResidenceCreate) SaveX(ctx context.Context) *Residence {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResidenceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResidenceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResidenceCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := residence.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := residence.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := residence.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResidenceCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Residence.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Residence.updated_at"`)}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Residence.name"`)}
	}
	return nil
}

func (rc *ResidenceCreate) sqlSave(ctx context.Context) (*Residence, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ResidenceCreate) createSpec() (*Residence, *sqlgraph.CreateSpec) {
	var (
		_node = &Residence{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(residence.Table, sqlgraph.NewFieldSpec(residence.FieldID, field.TypeUUID))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(residence.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(residence.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(residence.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := rc.mutation.OccupantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   residence.OccupantsTable,
			Columns: []string{residence.OccupantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResidenceCreateBulk is the builder for creating many Residence entities in bulk.
type ResidenceCreateBulk struct {
	config
	err      error
	builders []*ResidenceCreate
}

// Save creates the Residence entities in the database.
func (rcb *ResidenceCreateBulk) Save(ctx context.Context) ([]*Residence, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Residence, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResidenceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResidenceCreateBulk) SaveX(ctx context.Context) []*Residence {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResidenceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResidenceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
