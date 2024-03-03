// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/organizationhistory"
	"_examples/updateby_uuid/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationHistoryUpdate is the builder for updating OrganizationHistory entities.
type OrganizationHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationHistoryMutation
}

// Where appends a list predicates to the OrganizationHistoryUpdate builder.
func (ohu *OrganizationHistoryUpdate) Where(ps ...predicate.OrganizationHistory) *OrganizationHistoryUpdate {
	ohu.mutation.Where(ps...)
	return ohu
}

// SetName sets the "name" field.
func (ohu *OrganizationHistoryUpdate) SetName(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetName(s)
	return ohu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableName(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetName(*s)
	}
	return ohu
}

// SetInfo sets the "info" field.
func (ohu *OrganizationHistoryUpdate) SetInfo(m map[string]interface{}) *OrganizationHistoryUpdate {
	ohu.mutation.SetInfo(m)
	return ohu
}

// ClearInfo clears the value of the "info" field.
func (ohu *OrganizationHistoryUpdate) ClearInfo() *OrganizationHistoryUpdate {
	ohu.mutation.ClearInfo()
	return ohu
}

// SetUpdatedAt sets the "updated_at" field.
func (ohu *OrganizationHistoryUpdate) SetUpdatedAt(t time.Time) *OrganizationHistoryUpdate {
	ohu.mutation.SetUpdatedAt(t)
	return ohu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryUpdate {
	if t != nil {
		ohu.SetUpdatedAt(*t)
	}
	return ohu
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohu *OrganizationHistoryUpdate) Mutation() *OrganizationHistoryMutation {
	return ohu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ohu *OrganizationHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ohu.sqlSave, ohu.mutation, ohu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohu *OrganizationHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ohu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ohu *OrganizationHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ohu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohu *OrganizationHistoryUpdate) ExecX(ctx context.Context) {
	if err := ohu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ohu *OrganizationHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationhistory.Table, organizationhistory.Columns, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeInt))
	if ps := ohu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohu.mutation.RefCleared() {
		_spec.ClearField(organizationhistory.FieldRef, field.TypeUUID)
	}
	if ohu.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := ohu.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ohu.mutation.Info(); ok {
		_spec.SetField(organizationhistory.FieldInfo, field.TypeJSON, value)
	}
	if ohu.mutation.InfoCleared() {
		_spec.ClearField(organizationhistory.FieldInfo, field.TypeJSON)
	}
	if value, ok := ohu.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ohu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ohu.mutation.done = true
	return n, nil
}

// OrganizationHistoryUpdateOne is the builder for updating a single OrganizationHistory entity.
type OrganizationHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationHistoryMutation
}

// SetName sets the "name" field.
func (ohuo *OrganizationHistoryUpdateOne) SetName(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetName(s)
	return ohuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableName(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetName(*s)
	}
	return ohuo
}

// SetInfo sets the "info" field.
func (ohuo *OrganizationHistoryUpdateOne) SetInfo(m map[string]interface{}) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetInfo(m)
	return ohuo
}

// ClearInfo clears the value of the "info" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearInfo() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearInfo()
	return ohuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ohuo *OrganizationHistoryUpdateOne) SetUpdatedAt(t time.Time) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetUpdatedAt(t)
	return ohuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryUpdateOne {
	if t != nil {
		ohuo.SetUpdatedAt(*t)
	}
	return ohuo
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohuo *OrganizationHistoryUpdateOne) Mutation() *OrganizationHistoryMutation {
	return ohuo.mutation
}

// Where appends a list predicates to the OrganizationHistoryUpdate builder.
func (ohuo *OrganizationHistoryUpdateOne) Where(ps ...predicate.OrganizationHistory) *OrganizationHistoryUpdateOne {
	ohuo.mutation.Where(ps...)
	return ohuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ohuo *OrganizationHistoryUpdateOne) Select(field string, fields ...string) *OrganizationHistoryUpdateOne {
	ohuo.fields = append([]string{field}, fields...)
	return ohuo
}

// Save executes the query and returns the updated OrganizationHistory entity.
func (ohuo *OrganizationHistoryUpdateOne) Save(ctx context.Context) (*OrganizationHistory, error) {
	return withHooks(ctx, ohuo.sqlSave, ohuo.mutation, ohuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohuo *OrganizationHistoryUpdateOne) SaveX(ctx context.Context) *OrganizationHistory {
	node, err := ohuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ohuo *OrganizationHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ohuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohuo *OrganizationHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ohuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ohuo *OrganizationHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationhistory.Table, organizationhistory.Columns, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeInt))
	id, ok := ohuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ohuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationhistory.FieldID)
		for _, f := range fields {
			if !organizationhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ohuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohuo.mutation.RefCleared() {
		_spec.ClearField(organizationhistory.FieldRef, field.TypeUUID)
	}
	if ohuo.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := ohuo.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ohuo.mutation.Info(); ok {
		_spec.SetField(organizationhistory.FieldInfo, field.TypeJSON, value)
	}
	if ohuo.mutation.InfoCleared() {
		_spec.ClearField(organizationhistory.FieldInfo, field.TypeJSON)
	}
	if value, ok := ohuo.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &OrganizationHistory{config: ohuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ohuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ohuo.mutation.done = true
	return _node, nil
}
