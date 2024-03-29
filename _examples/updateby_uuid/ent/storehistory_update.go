// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/predicate"
	"_examples/updateby_uuid/ent/storehistory"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StoreHistoryUpdate is the builder for updating StoreHistory entities.
type StoreHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *StoreHistoryMutation
}

// Where appends a list predicates to the StoreHistoryUpdate builder.
func (shu *StoreHistoryUpdate) Where(ps ...predicate.StoreHistory) *StoreHistoryUpdate {
	shu.mutation.Where(ps...)
	return shu
}

// SetUpdatedAt sets the "updated_at" field.
func (shu *StoreHistoryUpdate) SetUpdatedAt(t time.Time) *StoreHistoryUpdate {
	shu.mutation.SetUpdatedAt(t)
	return shu
}

// SetName sets the "name" field.
func (shu *StoreHistoryUpdate) SetName(s string) *StoreHistoryUpdate {
	shu.mutation.SetName(s)
	return shu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (shu *StoreHistoryUpdate) SetNillableName(s *string) *StoreHistoryUpdate {
	if s != nil {
		shu.SetName(*s)
	}
	return shu
}

// SetRegion sets the "region" field.
func (shu *StoreHistoryUpdate) SetRegion(s string) *StoreHistoryUpdate {
	shu.mutation.SetRegion(s)
	return shu
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (shu *StoreHistoryUpdate) SetNillableRegion(s *string) *StoreHistoryUpdate {
	if s != nil {
		shu.SetRegion(*s)
	}
	return shu
}

// SetOrganizationID sets the "organization_id" field.
func (shu *StoreHistoryUpdate) SetOrganizationID(u uuid.UUID) *StoreHistoryUpdate {
	shu.mutation.SetOrganizationID(u)
	return shu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (shu *StoreHistoryUpdate) SetNillableOrganizationID(u *uuid.UUID) *StoreHistoryUpdate {
	if u != nil {
		shu.SetOrganizationID(*u)
	}
	return shu
}

// Mutation returns the StoreHistoryMutation object of the builder.
func (shu *StoreHistoryUpdate) Mutation() *StoreHistoryMutation {
	return shu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (shu *StoreHistoryUpdate) Save(ctx context.Context) (int, error) {
	shu.defaults()
	return withHooks(ctx, shu.sqlSave, shu.mutation, shu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shu *StoreHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := shu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (shu *StoreHistoryUpdate) Exec(ctx context.Context) error {
	_, err := shu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shu *StoreHistoryUpdate) ExecX(ctx context.Context) {
	if err := shu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shu *StoreHistoryUpdate) defaults() {
	if _, ok := shu.mutation.UpdatedAt(); !ok {
		v := storehistory.UpdateDefaultUpdatedAt()
		shu.mutation.SetUpdatedAt(v)
	}
}

func (shu *StoreHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(storehistory.Table, storehistory.Columns, sqlgraph.NewFieldSpec(storehistory.FieldID, field.TypeInt))
	if ps := shu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := shu.mutation.UpdatedAt(); ok {
		_spec.SetField(storehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if shu.mutation.RefCleared() {
		_spec.ClearField(storehistory.FieldRef, field.TypeUUID)
	}
	if shu.mutation.UpdatedByCleared() {
		_spec.ClearField(storehistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := shu.mutation.Name(); ok {
		_spec.SetField(storehistory.FieldName, field.TypeString, value)
	}
	if value, ok := shu.mutation.Region(); ok {
		_spec.SetField(storehistory.FieldRegion, field.TypeString, value)
	}
	if value, ok := shu.mutation.OrganizationID(); ok {
		_spec.SetField(storehistory.FieldOrganizationID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, shu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	shu.mutation.done = true
	return n, nil
}

// StoreHistoryUpdateOne is the builder for updating a single StoreHistory entity.
type StoreHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StoreHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (shuo *StoreHistoryUpdateOne) SetUpdatedAt(t time.Time) *StoreHistoryUpdateOne {
	shuo.mutation.SetUpdatedAt(t)
	return shuo
}

// SetName sets the "name" field.
func (shuo *StoreHistoryUpdateOne) SetName(s string) *StoreHistoryUpdateOne {
	shuo.mutation.SetName(s)
	return shuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (shuo *StoreHistoryUpdateOne) SetNillableName(s *string) *StoreHistoryUpdateOne {
	if s != nil {
		shuo.SetName(*s)
	}
	return shuo
}

// SetRegion sets the "region" field.
func (shuo *StoreHistoryUpdateOne) SetRegion(s string) *StoreHistoryUpdateOne {
	shuo.mutation.SetRegion(s)
	return shuo
}

// SetNillableRegion sets the "region" field if the given value is not nil.
func (shuo *StoreHistoryUpdateOne) SetNillableRegion(s *string) *StoreHistoryUpdateOne {
	if s != nil {
		shuo.SetRegion(*s)
	}
	return shuo
}

// SetOrganizationID sets the "organization_id" field.
func (shuo *StoreHistoryUpdateOne) SetOrganizationID(u uuid.UUID) *StoreHistoryUpdateOne {
	shuo.mutation.SetOrganizationID(u)
	return shuo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (shuo *StoreHistoryUpdateOne) SetNillableOrganizationID(u *uuid.UUID) *StoreHistoryUpdateOne {
	if u != nil {
		shuo.SetOrganizationID(*u)
	}
	return shuo
}

// Mutation returns the StoreHistoryMutation object of the builder.
func (shuo *StoreHistoryUpdateOne) Mutation() *StoreHistoryMutation {
	return shuo.mutation
}

// Where appends a list predicates to the StoreHistoryUpdate builder.
func (shuo *StoreHistoryUpdateOne) Where(ps ...predicate.StoreHistory) *StoreHistoryUpdateOne {
	shuo.mutation.Where(ps...)
	return shuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (shuo *StoreHistoryUpdateOne) Select(field string, fields ...string) *StoreHistoryUpdateOne {
	shuo.fields = append([]string{field}, fields...)
	return shuo
}

// Save executes the query and returns the updated StoreHistory entity.
func (shuo *StoreHistoryUpdateOne) Save(ctx context.Context) (*StoreHistory, error) {
	shuo.defaults()
	return withHooks(ctx, shuo.sqlSave, shuo.mutation, shuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (shuo *StoreHistoryUpdateOne) SaveX(ctx context.Context) *StoreHistory {
	node, err := shuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (shuo *StoreHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := shuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (shuo *StoreHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := shuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (shuo *StoreHistoryUpdateOne) defaults() {
	if _, ok := shuo.mutation.UpdatedAt(); !ok {
		v := storehistory.UpdateDefaultUpdatedAt()
		shuo.mutation.SetUpdatedAt(v)
	}
}

func (shuo *StoreHistoryUpdateOne) sqlSave(ctx context.Context) (_node *StoreHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(storehistory.Table, storehistory.Columns, sqlgraph.NewFieldSpec(storehistory.FieldID, field.TypeInt))
	id, ok := shuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StoreHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := shuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, storehistory.FieldID)
		for _, f := range fields {
			if !storehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != storehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := shuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := shuo.mutation.UpdatedAt(); ok {
		_spec.SetField(storehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if shuo.mutation.RefCleared() {
		_spec.ClearField(storehistory.FieldRef, field.TypeUUID)
	}
	if shuo.mutation.UpdatedByCleared() {
		_spec.ClearField(storehistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := shuo.mutation.Name(); ok {
		_spec.SetField(storehistory.FieldName, field.TypeString, value)
	}
	if value, ok := shuo.mutation.Region(); ok {
		_spec.SetField(storehistory.FieldRegion, field.TypeString, value)
	}
	if value, ok := shuo.mutation.OrganizationID(); ok {
		_spec.SetField(storehistory.FieldOrganizationID, field.TypeUUID, value)
	}
	_node = &StoreHistory{config: shuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, shuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{storehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	shuo.mutation.done = true
	return _node, nil
}
