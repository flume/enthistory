// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/without_updatedby/ent/character"
	"_examples/without_updatedby/ent/predicate"
	"_examples/without_updatedby/ent/residence"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ResidenceUpdate is the builder for updating Residence entities.
type ResidenceUpdate struct {
	config
	hooks    []Hook
	mutation *ResidenceMutation
}

// Where appends a list predicates to the ResidenceUpdate builder.
func (ru *ResidenceUpdate) Where(ps ...predicate.Residence) *ResidenceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *ResidenceUpdate) SetUpdatedAt(t time.Time) *ResidenceUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetName sets the "name" field.
func (ru *ResidenceUpdate) SetName(s string) *ResidenceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *ResidenceUpdate) SetNillableName(s *string) *ResidenceUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// AddOccupantIDs adds the "occupants" edge to the Character entity by IDs.
func (ru *ResidenceUpdate) AddOccupantIDs(ids ...int) *ResidenceUpdate {
	ru.mutation.AddOccupantIDs(ids...)
	return ru
}

// AddOccupants adds the "occupants" edges to the Character entity.
func (ru *ResidenceUpdate) AddOccupants(c ...*Character) *ResidenceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.AddOccupantIDs(ids...)
}

// Mutation returns the ResidenceMutation object of the builder.
func (ru *ResidenceUpdate) Mutation() *ResidenceMutation {
	return ru.mutation
}

// ClearOccupants clears all "occupants" edges to the Character entity.
func (ru *ResidenceUpdate) ClearOccupants() *ResidenceUpdate {
	ru.mutation.ClearOccupants()
	return ru
}

// RemoveOccupantIDs removes the "occupants" edge to Character entities by IDs.
func (ru *ResidenceUpdate) RemoveOccupantIDs(ids ...int) *ResidenceUpdate {
	ru.mutation.RemoveOccupantIDs(ids...)
	return ru
}

// RemoveOccupants removes "occupants" edges to Character entities.
func (ru *ResidenceUpdate) RemoveOccupants(c ...*Character) *ResidenceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ru.RemoveOccupantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResidenceUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResidenceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResidenceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResidenceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ResidenceUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := residence.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

func (ru *ResidenceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(residence.Table, residence.Columns, sqlgraph.NewFieldSpec(residence.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(residence.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(residence.FieldName, field.TypeString, value)
	}
	if ru.mutation.OccupantsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedOccupantsIDs(); len(nodes) > 0 && !ru.mutation.OccupantsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.OccupantsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{residence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ResidenceUpdateOne is the builder for updating a single Residence entity.
type ResidenceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResidenceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *ResidenceUpdateOne) SetUpdatedAt(t time.Time) *ResidenceUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetName sets the "name" field.
func (ruo *ResidenceUpdateOne) SetName(s string) *ResidenceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *ResidenceUpdateOne) SetNillableName(s *string) *ResidenceUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// AddOccupantIDs adds the "occupants" edge to the Character entity by IDs.
func (ruo *ResidenceUpdateOne) AddOccupantIDs(ids ...int) *ResidenceUpdateOne {
	ruo.mutation.AddOccupantIDs(ids...)
	return ruo
}

// AddOccupants adds the "occupants" edges to the Character entity.
func (ruo *ResidenceUpdateOne) AddOccupants(c ...*Character) *ResidenceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.AddOccupantIDs(ids...)
}

// Mutation returns the ResidenceMutation object of the builder.
func (ruo *ResidenceUpdateOne) Mutation() *ResidenceMutation {
	return ruo.mutation
}

// ClearOccupants clears all "occupants" edges to the Character entity.
func (ruo *ResidenceUpdateOne) ClearOccupants() *ResidenceUpdateOne {
	ruo.mutation.ClearOccupants()
	return ruo
}

// RemoveOccupantIDs removes the "occupants" edge to Character entities by IDs.
func (ruo *ResidenceUpdateOne) RemoveOccupantIDs(ids ...int) *ResidenceUpdateOne {
	ruo.mutation.RemoveOccupantIDs(ids...)
	return ruo
}

// RemoveOccupants removes "occupants" edges to Character entities.
func (ruo *ResidenceUpdateOne) RemoveOccupants(c ...*Character) *ResidenceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ruo.RemoveOccupantIDs(ids...)
}

// Where appends a list predicates to the ResidenceUpdate builder.
func (ruo *ResidenceUpdateOne) Where(ps ...predicate.Residence) *ResidenceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResidenceUpdateOne) Select(field string, fields ...string) *ResidenceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Residence entity.
func (ruo *ResidenceUpdateOne) Save(ctx context.Context) (*Residence, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResidenceUpdateOne) SaveX(ctx context.Context) *Residence {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResidenceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResidenceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ResidenceUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := residence.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

func (ruo *ResidenceUpdateOne) sqlSave(ctx context.Context) (_node *Residence, err error) {
	_spec := sqlgraph.NewUpdateSpec(residence.Table, residence.Columns, sqlgraph.NewFieldSpec(residence.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Residence.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, residence.FieldID)
		for _, f := range fields {
			if !residence.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != residence.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(residence.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(residence.FieldName, field.TypeString, value)
	}
	if ruo.mutation.OccupantsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedOccupantsIDs(); len(nodes) > 0 && !ruo.mutation.OccupantsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.OccupantsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Residence{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{residence.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
