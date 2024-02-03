// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/organization"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/predicate"
	"github.com/flume/enthistory/_examples/updateby_uuid/ent/store"
	"github.com/google/uuid"
)

// OrganizationUpdate is the builder for updating Organization entities.
type OrganizationUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationMutation
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ou *OrganizationUpdate) Where(ps ...predicate.Organization) *OrganizationUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrganizationUpdate) SetUpdatedAt(t time.Time) *OrganizationUpdate {
	ou.mutation.SetUpdatedAt(t)
	return ou
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableUpdatedAt(t *time.Time) *OrganizationUpdate {
	if t != nil {
		ou.SetUpdatedAt(*t)
	}
	return ou
}

// SetName sets the "name" field.
func (ou *OrganizationUpdate) SetName(s string) *OrganizationUpdate {
	ou.mutation.SetName(s)
	return ou
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ou *OrganizationUpdate) SetNillableName(s *string) *OrganizationUpdate {
	if s != nil {
		ou.SetName(*s)
	}
	return ou
}

// SetInfo sets the "info" field.
func (ou *OrganizationUpdate) SetInfo(m map[string]interface{}) *OrganizationUpdate {
	ou.mutation.SetInfo(m)
	return ou
}

// ClearInfo clears the value of the "info" field.
func (ou *OrganizationUpdate) ClearInfo() *OrganizationUpdate {
	ou.mutation.ClearInfo()
	return ou
}

// AddOrganizationStoreIDs adds the "organization_stores" edge to the Store entity by IDs.
func (ou *OrganizationUpdate) AddOrganizationStoreIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.AddOrganizationStoreIDs(ids...)
	return ou
}

// AddOrganizationStores adds the "organization_stores" edges to the Store entity.
func (ou *OrganizationUpdate) AddOrganizationStores(s ...*Store) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ou.AddOrganizationStoreIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ou *OrganizationUpdate) Mutation() *OrganizationMutation {
	return ou.mutation
}

// ClearOrganizationStores clears all "organization_stores" edges to the Store entity.
func (ou *OrganizationUpdate) ClearOrganizationStores() *OrganizationUpdate {
	ou.mutation.ClearOrganizationStores()
	return ou
}

// RemoveOrganizationStoreIDs removes the "organization_stores" edge to Store entities by IDs.
func (ou *OrganizationUpdate) RemoveOrganizationStoreIDs(ids ...uuid.UUID) *OrganizationUpdate {
	ou.mutation.RemoveOrganizationStoreIDs(ids...)
	return ou
}

// RemoveOrganizationStores removes "organization_stores" edges to Store entities.
func (ou *OrganizationUpdate) RemoveOrganizationStores(s ...*Store) *OrganizationUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ou.RemoveOrganizationStoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrganizationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrganizationUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrganizationUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrganizationUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrganizationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ou.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ou.mutation.Info(); ok {
		_spec.SetField(organization.FieldInfo, field.TypeJSON, value)
	}
	if ou.mutation.InfoCleared() {
		_spec.ClearField(organization.FieldInfo, field.TypeJSON)
	}
	if ou.mutation.OrganizationStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.RemovedOrganizationStoresIDs(); len(nodes) > 0 && !ou.mutation.OrganizationStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ou.mutation.OrganizationStoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrganizationUpdateOne is the builder for updating a single Organization entity.
type OrganizationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrganizationUpdateOne) SetUpdatedAt(t time.Time) *OrganizationUpdateOne {
	ouo.mutation.SetUpdatedAt(t)
	return ouo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrganizationUpdateOne {
	if t != nil {
		ouo.SetUpdatedAt(*t)
	}
	return ouo
}

// SetName sets the "name" field.
func (ouo *OrganizationUpdateOne) SetName(s string) *OrganizationUpdateOne {
	ouo.mutation.SetName(s)
	return ouo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ouo *OrganizationUpdateOne) SetNillableName(s *string) *OrganizationUpdateOne {
	if s != nil {
		ouo.SetName(*s)
	}
	return ouo
}

// SetInfo sets the "info" field.
func (ouo *OrganizationUpdateOne) SetInfo(m map[string]interface{}) *OrganizationUpdateOne {
	ouo.mutation.SetInfo(m)
	return ouo
}

// ClearInfo clears the value of the "info" field.
func (ouo *OrganizationUpdateOne) ClearInfo() *OrganizationUpdateOne {
	ouo.mutation.ClearInfo()
	return ouo
}

// AddOrganizationStoreIDs adds the "organization_stores" edge to the Store entity by IDs.
func (ouo *OrganizationUpdateOne) AddOrganizationStoreIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.AddOrganizationStoreIDs(ids...)
	return ouo
}

// AddOrganizationStores adds the "organization_stores" edges to the Store entity.
func (ouo *OrganizationUpdateOne) AddOrganizationStores(s ...*Store) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ouo.AddOrganizationStoreIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (ouo *OrganizationUpdateOne) Mutation() *OrganizationMutation {
	return ouo.mutation
}

// ClearOrganizationStores clears all "organization_stores" edges to the Store entity.
func (ouo *OrganizationUpdateOne) ClearOrganizationStores() *OrganizationUpdateOne {
	ouo.mutation.ClearOrganizationStores()
	return ouo
}

// RemoveOrganizationStoreIDs removes the "organization_stores" edge to Store entities by IDs.
func (ouo *OrganizationUpdateOne) RemoveOrganizationStoreIDs(ids ...uuid.UUID) *OrganizationUpdateOne {
	ouo.mutation.RemoveOrganizationStoreIDs(ids...)
	return ouo
}

// RemoveOrganizationStores removes "organization_stores" edges to Store entities.
func (ouo *OrganizationUpdateOne) RemoveOrganizationStores(s ...*Store) *OrganizationUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ouo.RemoveOrganizationStoreIDs(ids...)
}

// Where appends a list predicates to the OrganizationUpdate builder.
func (ouo *OrganizationUpdateOne) Where(ps ...predicate.Organization) *OrganizationUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrganizationUpdateOne) Select(field string, fields ...string) *OrganizationUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Organization entity.
func (ouo *OrganizationUpdateOne) Save(ctx context.Context) (*Organization, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) SaveX(ctx context.Context) *Organization {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrganizationUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrganizationUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrganizationUpdateOne) sqlSave(ctx context.Context) (_node *Organization, err error) {
	_spec := sqlgraph.NewUpdateSpec(organization.Table, organization.Columns, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Organization.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organization.FieldID)
		for _, f := range fields {
			if !organization.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organization.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ouo.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Info(); ok {
		_spec.SetField(organization.FieldInfo, field.TypeJSON, value)
	}
	if ouo.mutation.InfoCleared() {
		_spec.ClearField(organization.FieldInfo, field.TypeJSON)
	}
	if ouo.mutation.OrganizationStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.RemovedOrganizationStoresIDs(); len(nodes) > 0 && !ouo.mutation.OrganizationStoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ouo.mutation.OrganizationStoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OrganizationStoresTable,
			Columns: []string{organization.OrganizationStoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Organization{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organization.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
