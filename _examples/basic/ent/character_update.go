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

	"github.com/flume/enthistory/_examples/basic/ent/character"
	"github.com/flume/enthistory/_examples/basic/ent/friendship"
	"github.com/flume/enthistory/_examples/basic/ent/predicate"
)

// CharacterUpdate is the builder for updating Character entities.
type CharacterUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterMutation
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cu *CharacterUpdate) Where(ps ...predicate.Character) *CharacterUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CharacterUpdate) SetUpdatedAt(t time.Time) *CharacterUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cu *CharacterUpdate) SetNillableUpdatedAt(t *time.Time) *CharacterUpdate {
	if t != nil {
		cu.SetUpdatedAt(*t)
	}
	return cu
}

// SetAge sets the "age" field.
func (cu *CharacterUpdate) SetAge(i int) *CharacterUpdate {
	cu.mutation.ResetAge()
	cu.mutation.SetAge(i)
	return cu
}

// AddAge adds i to the "age" field.
func (cu *CharacterUpdate) AddAge(i int) *CharacterUpdate {
	cu.mutation.AddAge(i)
	return cu
}

// SetName sets the "name" field.
func (cu *CharacterUpdate) SetName(s string) *CharacterUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddFriendIDs adds the "friends" edge to the Character entity by IDs.
func (cu *CharacterUpdate) AddFriendIDs(ids ...int) *CharacterUpdate {
	cu.mutation.AddFriendIDs(ids...)
	return cu
}

// AddFriends adds the "friends" edges to the Character entity.
func (cu *CharacterUpdate) AddFriends(c ...*Character) *CharacterUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddFriendIDs(ids...)
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (cu *CharacterUpdate) AddFriendshipIDs(ids ...int) *CharacterUpdate {
	cu.mutation.AddFriendshipIDs(ids...)
	return cu
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (cu *CharacterUpdate) AddFriendships(f ...*Friendship) *CharacterUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.AddFriendshipIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cu *CharacterUpdate) Mutation() *CharacterMutation {
	return cu.mutation
}

// ClearFriends clears all "friends" edges to the Character entity.
func (cu *CharacterUpdate) ClearFriends() *CharacterUpdate {
	cu.mutation.ClearFriends()
	return cu
}

// RemoveFriendIDs removes the "friends" edge to Character entities by IDs.
func (cu *CharacterUpdate) RemoveFriendIDs(ids ...int) *CharacterUpdate {
	cu.mutation.RemoveFriendIDs(ids...)
	return cu
}

// RemoveFriends removes "friends" edges to Character entities.
func (cu *CharacterUpdate) RemoveFriends(c ...*Character) *CharacterUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveFriendIDs(ids...)
}

// ClearFriendships clears all "friendships" edges to the Friendship entity.
func (cu *CharacterUpdate) ClearFriendships() *CharacterUpdate {
	cu.mutation.ClearFriendships()
	return cu
}

// RemoveFriendshipIDs removes the "friendships" edge to Friendship entities by IDs.
func (cu *CharacterUpdate) RemoveFriendshipIDs(ids ...int) *CharacterUpdate {
	cu.mutation.RemoveFriendshipIDs(ids...)
	return cu
}

// RemoveFriendships removes "friendships" edges to Friendship entities.
func (cu *CharacterUpdate) RemoveFriendships(f ...*Friendship) *CharacterUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cu.RemoveFriendshipIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CharacterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CharacterUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CharacterUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CharacterUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CharacterUpdate) check() error {
	if v, ok := cu.mutation.Age(); ok {
		if err := character.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Character.age": %w`, err)}
		}
	}
	return nil
}

func (cu *CharacterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(character.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Age(); ok {
		_spec.SetField(character.FieldAge, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedAge(); ok {
		_spec.AddField(character.FieldAge, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if cu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		createE := &FriendshipCreate{config: cu.config, mutation: newFriendshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !cu.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: cu.config, mutation: newFriendshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: cu.config, mutation: newFriendshipMutation(cu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedFriendshipsIDs(); len(nodes) > 0 && !cu.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CharacterUpdateOne is the builder for updating a single Character entity.
type CharacterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CharacterUpdateOne) SetUpdatedAt(t time.Time) *CharacterUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cuo *CharacterUpdateOne) SetNillableUpdatedAt(t *time.Time) *CharacterUpdateOne {
	if t != nil {
		cuo.SetUpdatedAt(*t)
	}
	return cuo
}

// SetAge sets the "age" field.
func (cuo *CharacterUpdateOne) SetAge(i int) *CharacterUpdateOne {
	cuo.mutation.ResetAge()
	cuo.mutation.SetAge(i)
	return cuo
}

// AddAge adds i to the "age" field.
func (cuo *CharacterUpdateOne) AddAge(i int) *CharacterUpdateOne {
	cuo.mutation.AddAge(i)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CharacterUpdateOne) SetName(s string) *CharacterUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddFriendIDs adds the "friends" edge to the Character entity by IDs.
func (cuo *CharacterUpdateOne) AddFriendIDs(ids ...int) *CharacterUpdateOne {
	cuo.mutation.AddFriendIDs(ids...)
	return cuo
}

// AddFriends adds the "friends" edges to the Character entity.
func (cuo *CharacterUpdateOne) AddFriends(c ...*Character) *CharacterUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddFriendIDs(ids...)
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (cuo *CharacterUpdateOne) AddFriendshipIDs(ids ...int) *CharacterUpdateOne {
	cuo.mutation.AddFriendshipIDs(ids...)
	return cuo
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (cuo *CharacterUpdateOne) AddFriendships(f ...*Friendship) *CharacterUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.AddFriendshipIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cuo *CharacterUpdateOne) Mutation() *CharacterMutation {
	return cuo.mutation
}

// ClearFriends clears all "friends" edges to the Character entity.
func (cuo *CharacterUpdateOne) ClearFriends() *CharacterUpdateOne {
	cuo.mutation.ClearFriends()
	return cuo
}

// RemoveFriendIDs removes the "friends" edge to Character entities by IDs.
func (cuo *CharacterUpdateOne) RemoveFriendIDs(ids ...int) *CharacterUpdateOne {
	cuo.mutation.RemoveFriendIDs(ids...)
	return cuo
}

// RemoveFriends removes "friends" edges to Character entities.
func (cuo *CharacterUpdateOne) RemoveFriends(c ...*Character) *CharacterUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveFriendIDs(ids...)
}

// ClearFriendships clears all "friendships" edges to the Friendship entity.
func (cuo *CharacterUpdateOne) ClearFriendships() *CharacterUpdateOne {
	cuo.mutation.ClearFriendships()
	return cuo
}

// RemoveFriendshipIDs removes the "friendships" edge to Friendship entities by IDs.
func (cuo *CharacterUpdateOne) RemoveFriendshipIDs(ids ...int) *CharacterUpdateOne {
	cuo.mutation.RemoveFriendshipIDs(ids...)
	return cuo
}

// RemoveFriendships removes "friendships" edges to Friendship entities.
func (cuo *CharacterUpdateOne) RemoveFriendships(f ...*Friendship) *CharacterUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cuo.RemoveFriendshipIDs(ids...)
}

// Where appends a list predicates to the CharacterUpdate builder.
func (cuo *CharacterUpdateOne) Where(ps ...predicate.Character) *CharacterUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CharacterUpdateOne) Select(field string, fields ...string) *CharacterUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Character entity.
func (cuo *CharacterUpdateOne) Save(ctx context.Context) (*Character, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CharacterUpdateOne) SaveX(ctx context.Context) *Character {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CharacterUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CharacterUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CharacterUpdateOne) check() error {
	if v, ok := cuo.mutation.Age(); ok {
		if err := character.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Character.age": %w`, err)}
		}
	}
	return nil
}

func (cuo *CharacterUpdateOne) sqlSave(ctx context.Context) (_node *Character, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Character.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for _, f := range fields {
			if !character.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != character.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(character.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Age(); ok {
		_spec.SetField(character.FieldAge, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedAge(); ok {
		_spec.AddField(character.FieldAge, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
	}
	if cuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		createE := &FriendshipCreate{config: cuo.config, mutation: newFriendshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedFriendsIDs(); len(nodes) > 0 && !cuo.mutation.FriendsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: cuo.config, mutation: newFriendshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.FriendsTable,
			Columns: character.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: cuo.config, mutation: newFriendshipMutation(cuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedFriendshipsIDs(); len(nodes) > 0 && !cuo.mutation.FriendshipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.FriendshipsTable,
			Columns: []string{character.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Character{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{character.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
