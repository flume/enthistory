// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/basic/ent/character"
	"_examples/basic/ent/friendship"
	"_examples/basic/ent/predicate"
	"_examples/basic/ent/residence"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CharacterQuery is the builder for querying Character entities.
type CharacterQuery struct {
	config
	ctx             *QueryContext
	order           []character.OrderOption
	inters          []Interceptor
	predicates      []predicate.Character
	withFriends     *CharacterQuery
	withResidence   *ResidenceQuery
	withFriendships *FriendshipQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CharacterQuery builder.
func (cq *CharacterQuery) Where(ps ...predicate.Character) *CharacterQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CharacterQuery) Limit(limit int) *CharacterQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CharacterQuery) Offset(offset int) *CharacterQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CharacterQuery) Unique(unique bool) *CharacterQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CharacterQuery) Order(o ...character.OrderOption) *CharacterQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryFriends chains the current query on the "friends" edge.
func (cq *CharacterQuery) QueryFriends() *CharacterQuery {
	query := (&CharacterClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, character.FriendsTable, character.FriendsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryResidence chains the current query on the "residence" edge.
func (cq *CharacterQuery) QueryResidence() *ResidenceQuery {
	query := (&ResidenceClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(residence.Table, residence.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, character.ResidenceTable, character.ResidenceColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriendships chains the current query on the "friendships" edge.
func (cq *CharacterQuery) QueryFriendships() *FriendshipQuery {
	query := (&FriendshipClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(friendship.Table, friendship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, character.FriendshipsTable, character.FriendshipsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Character entity from the query.
// Returns a *NotFoundError when no Character was found.
func (cq *CharacterQuery) First(ctx context.Context) (*Character, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{character.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CharacterQuery) FirstX(ctx context.Context) *Character {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Character ID from the query.
// Returns a *NotFoundError when no Character ID was found.
func (cq *CharacterQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{character.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CharacterQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Character entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Character entity is found.
// Returns a *NotFoundError when no Character entities are found.
func (cq *CharacterQuery) Only(ctx context.Context) (*Character, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{character.Label}
	default:
		return nil, &NotSingularError{character.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CharacterQuery) OnlyX(ctx context.Context) *Character {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Character ID in the query.
// Returns a *NotSingularError when more than one Character ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CharacterQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{character.Label}
	default:
		err = &NotSingularError{character.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CharacterQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Characters.
func (cq *CharacterQuery) All(ctx context.Context) ([]*Character, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Character, *CharacterQuery]()
	return withInterceptors[[]*Character](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CharacterQuery) AllX(ctx context.Context) []*Character {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Character IDs.
func (cq *CharacterQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(character.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CharacterQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CharacterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CharacterQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CharacterQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CharacterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CharacterQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CharacterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CharacterQuery) Clone() *CharacterQuery {
	if cq == nil {
		return nil
	}
	return &CharacterQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]character.OrderOption{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Character{}, cq.predicates...),
		withFriends:     cq.withFriends.Clone(),
		withResidence:   cq.withResidence.Clone(),
		withFriendships: cq.withFriendships.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithFriends tells the query-builder to eager-load the nodes that are connected to
// the "friends" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithFriends(opts ...func(*CharacterQuery)) *CharacterQuery {
	query := (&CharacterClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withFriends = query
	return cq
}

// WithResidence tells the query-builder to eager-load the nodes that are connected to
// the "residence" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithResidence(opts ...func(*ResidenceQuery)) *CharacterQuery {
	query := (&ResidenceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withResidence = query
	return cq
}

// WithFriendships tells the query-builder to eager-load the nodes that are connected to
// the "friendships" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithFriendships(opts ...func(*FriendshipQuery)) *CharacterQuery {
	query := (&FriendshipClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withFriendships = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Character.Query().
//		GroupBy(character.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CharacterQuery) GroupBy(field string, fields ...string) *CharacterGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CharacterGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = character.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Character.Query().
//		Select(character.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CharacterQuery) Select(fields ...string) *CharacterSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CharacterSelect{CharacterQuery: cq}
	sbuild.label = character.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CharacterSelect configured with the given aggregations.
func (cq *CharacterQuery) Aggregate(fns ...AggregateFunc) *CharacterSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CharacterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !character.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CharacterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Character, error) {
	var (
		nodes       = []*Character{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [3]bool{
			cq.withFriends != nil,
			cq.withResidence != nil,
			cq.withFriendships != nil,
		}
	)
	if cq.withResidence != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, character.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Character).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Character{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withFriends; query != nil {
		if err := cq.loadFriends(ctx, query, nodes,
			func(n *Character) { n.Edges.Friends = []*Character{} },
			func(n *Character, e *Character) { n.Edges.Friends = append(n.Edges.Friends, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withResidence; query != nil {
		if err := cq.loadResidence(ctx, query, nodes, nil,
			func(n *Character, e *Residence) { n.Edges.Residence = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withFriendships; query != nil {
		if err := cq.loadFriendships(ctx, query, nodes,
			func(n *Character) { n.Edges.Friendships = []*Friendship{} },
			func(n *Character, e *Friendship) { n.Edges.Friendships = append(n.Edges.Friendships, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CharacterQuery) loadFriends(ctx context.Context, query *CharacterQuery, nodes []*Character, init func(*Character), assign func(*Character, *Character)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Character)
	nids := make(map[int]map[*Character]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(character.FriendsTable)
		s.Join(joinT).On(s.C(character.FieldID), joinT.C(character.FriendsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(character.FriendsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(character.FriendsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Character]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Character](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "friends" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (cq *CharacterQuery) loadResidence(ctx context.Context, query *ResidenceQuery, nodes []*Character, init func(*Character), assign func(*Character, *Residence)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Character)
	for i := range nodes {
		if nodes[i].residence_occupants == nil {
			continue
		}
		fk := *nodes[i].residence_occupants
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(residence.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "residence_occupants" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CharacterQuery) loadFriendships(ctx context.Context, query *FriendshipQuery, nodes []*Character, init func(*Character), assign func(*Character, *Friendship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Character)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(friendship.FieldCharacterID)
	}
	query.Where(predicate.Friendship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(character.FriendshipsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CharacterID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "character_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CharacterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CharacterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for i := range fields {
			if fields[i] != character.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CharacterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(character.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = character.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CharacterGroupBy is the group-by builder for Character entities.
type CharacterGroupBy struct {
	selector
	build *CharacterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CharacterGroupBy) Aggregate(fns ...AggregateFunc) *CharacterGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CharacterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterQuery, *CharacterGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CharacterGroupBy) sqlScan(ctx context.Context, root *CharacterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CharacterSelect is the builder for selecting fields of Character entities.
type CharacterSelect struct {
	*CharacterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CharacterSelect) Aggregate(fns ...AggregateFunc) *CharacterSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CharacterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterQuery, *CharacterSelect](ctx, cs.CharacterQuery, cs, cs.inters, v)
}

func (cs *CharacterSelect) sqlScan(ctx context.Context, root *CharacterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
