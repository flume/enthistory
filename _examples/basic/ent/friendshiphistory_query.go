// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/flume/enthistory/_examples/basic/ent/friendshiphistory"
	"github.com/flume/enthistory/_examples/basic/ent/predicate"
)

// FriendshipHistoryQuery is the builder for querying FriendshipHistory entities.
type FriendshipHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []friendshiphistory.OrderOption
	inters     []Interceptor
	predicates []predicate.FriendshipHistory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FriendshipHistoryQuery builder.
func (fhq *FriendshipHistoryQuery) Where(ps ...predicate.FriendshipHistory) *FriendshipHistoryQuery {
	fhq.predicates = append(fhq.predicates, ps...)
	return fhq
}

// Limit the number of records to be returned by this query.
func (fhq *FriendshipHistoryQuery) Limit(limit int) *FriendshipHistoryQuery {
	fhq.ctx.Limit = &limit
	return fhq
}

// Offset to start from.
func (fhq *FriendshipHistoryQuery) Offset(offset int) *FriendshipHistoryQuery {
	fhq.ctx.Offset = &offset
	return fhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fhq *FriendshipHistoryQuery) Unique(unique bool) *FriendshipHistoryQuery {
	fhq.ctx.Unique = &unique
	return fhq
}

// Order specifies how the records should be ordered.
func (fhq *FriendshipHistoryQuery) Order(o ...friendshiphistory.OrderOption) *FriendshipHistoryQuery {
	fhq.order = append(fhq.order, o...)
	return fhq
}

// First returns the first FriendshipHistory entity from the query.
// Returns a *NotFoundError when no FriendshipHistory was found.
func (fhq *FriendshipHistoryQuery) First(ctx context.Context) (*FriendshipHistory, error) {
	nodes, err := fhq.Limit(1).All(setContextOp(ctx, fhq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{friendshiphistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) FirstX(ctx context.Context) *FriendshipHistory {
	node, err := fhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FriendshipHistory ID from the query.
// Returns a *NotFoundError when no FriendshipHistory ID was found.
func (fhq *FriendshipHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fhq.Limit(1).IDs(setContextOp(ctx, fhq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{friendshiphistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := fhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FriendshipHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FriendshipHistory entity is found.
// Returns a *NotFoundError when no FriendshipHistory entities are found.
func (fhq *FriendshipHistoryQuery) Only(ctx context.Context) (*FriendshipHistory, error) {
	nodes, err := fhq.Limit(2).All(setContextOp(ctx, fhq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{friendshiphistory.Label}
	default:
		return nil, &NotSingularError{friendshiphistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) OnlyX(ctx context.Context) *FriendshipHistory {
	node, err := fhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FriendshipHistory ID in the query.
// Returns a *NotSingularError when more than one FriendshipHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (fhq *FriendshipHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fhq.Limit(2).IDs(setContextOp(ctx, fhq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{friendshiphistory.Label}
	default:
		err = &NotSingularError{friendshiphistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := fhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FriendshipHistories.
func (fhq *FriendshipHistoryQuery) All(ctx context.Context) ([]*FriendshipHistory, error) {
	ctx = setContextOp(ctx, fhq.ctx, "All")
	if err := fhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FriendshipHistory, *FriendshipHistoryQuery]()
	return withInterceptors[[]*FriendshipHistory](ctx, fhq, qr, fhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) AllX(ctx context.Context) []*FriendshipHistory {
	nodes, err := fhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FriendshipHistory IDs.
func (fhq *FriendshipHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fhq.ctx.Unique == nil && fhq.path != nil {
		fhq.Unique(true)
	}
	ctx = setContextOp(ctx, fhq.ctx, "IDs")
	if err = fhq.Select(friendshiphistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := fhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fhq *FriendshipHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fhq.ctx, "Count")
	if err := fhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fhq, querierCount[*FriendshipHistoryQuery](), fhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) CountX(ctx context.Context) int {
	count, err := fhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fhq *FriendshipHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fhq.ctx, "Exist")
	switch _, err := fhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fhq *FriendshipHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := fhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FriendshipHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fhq *FriendshipHistoryQuery) Clone() *FriendshipHistoryQuery {
	if fhq == nil {
		return nil
	}
	return &FriendshipHistoryQuery{
		config:     fhq.config,
		ctx:        fhq.ctx.Clone(),
		order:      append([]friendshiphistory.OrderOption{}, fhq.order...),
		inters:     append([]Interceptor{}, fhq.inters...),
		predicates: append([]predicate.FriendshipHistory{}, fhq.predicates...),
		// clone intermediate query.
		sql:  fhq.sql.Clone(),
		path: fhq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FriendshipHistory.Query().
//		GroupBy(friendshiphistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fhq *FriendshipHistoryQuery) GroupBy(field string, fields ...string) *FriendshipHistoryGroupBy {
	fhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FriendshipHistoryGroupBy{build: fhq}
	grbuild.flds = &fhq.ctx.Fields
	grbuild.label = friendshiphistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.FriendshipHistory.Query().
//		Select(friendshiphistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (fhq *FriendshipHistoryQuery) Select(fields ...string) *FriendshipHistorySelect {
	fhq.ctx.Fields = append(fhq.ctx.Fields, fields...)
	sbuild := &FriendshipHistorySelect{FriendshipHistoryQuery: fhq}
	sbuild.label = friendshiphistory.Label
	sbuild.flds, sbuild.scan = &fhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FriendshipHistorySelect configured with the given aggregations.
func (fhq *FriendshipHistoryQuery) Aggregate(fns ...AggregateFunc) *FriendshipHistorySelect {
	return fhq.Select().Aggregate(fns...)
}

func (fhq *FriendshipHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fhq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fhq); err != nil {
				return err
			}
		}
	}
	for _, f := range fhq.ctx.Fields {
		if !friendshiphistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fhq.path != nil {
		prev, err := fhq.path(ctx)
		if err != nil {
			return err
		}
		fhq.sql = prev
	}
	return nil
}

func (fhq *FriendshipHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FriendshipHistory, error) {
	var (
		nodes = []*FriendshipHistory{}
		_spec = fhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FriendshipHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FriendshipHistory{config: fhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fhq *FriendshipHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fhq.querySpec()
	_spec.Node.Columns = fhq.ctx.Fields
	if len(fhq.ctx.Fields) > 0 {
		_spec.Unique = fhq.ctx.Unique != nil && *fhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fhq.driver, _spec)
}

func (fhq *FriendshipHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(friendshiphistory.Table, friendshiphistory.Columns, sqlgraph.NewFieldSpec(friendshiphistory.FieldID, field.TypeInt))
	_spec.From = fhq.sql
	if unique := fhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fhq.path != nil {
		_spec.Unique = true
	}
	if fields := fhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendshiphistory.FieldID)
		for i := range fields {
			if fields[i] != friendshiphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fhq *FriendshipHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fhq.driver.Dialect())
	t1 := builder.Table(friendshiphistory.Table)
	columns := fhq.ctx.Fields
	if len(columns) == 0 {
		columns = friendshiphistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fhq.sql != nil {
		selector = fhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fhq.ctx.Unique != nil && *fhq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fhq.predicates {
		p(selector)
	}
	for _, p := range fhq.order {
		p(selector)
	}
	if offset := fhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FriendshipHistoryGroupBy is the group-by builder for FriendshipHistory entities.
type FriendshipHistoryGroupBy struct {
	selector
	build *FriendshipHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fhgb *FriendshipHistoryGroupBy) Aggregate(fns ...AggregateFunc) *FriendshipHistoryGroupBy {
	fhgb.fns = append(fhgb.fns, fns...)
	return fhgb
}

// Scan applies the selector query and scans the result into the given value.
func (fhgb *FriendshipHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fhgb.build.ctx, "GroupBy")
	if err := fhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipHistoryQuery, *FriendshipHistoryGroupBy](ctx, fhgb.build, fhgb, fhgb.build.inters, v)
}

func (fhgb *FriendshipHistoryGroupBy) sqlScan(ctx context.Context, root *FriendshipHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fhgb.fns))
	for _, fn := range fhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fhgb.flds)+len(fhgb.fns))
		for _, f := range *fhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FriendshipHistorySelect is the builder for selecting fields of FriendshipHistory entities.
type FriendshipHistorySelect struct {
	*FriendshipHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fhs *FriendshipHistorySelect) Aggregate(fns ...AggregateFunc) *FriendshipHistorySelect {
	fhs.fns = append(fhs.fns, fns...)
	return fhs
}

// Scan applies the selector query and scans the result into the given value.
func (fhs *FriendshipHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fhs.ctx, "Select")
	if err := fhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipHistoryQuery, *FriendshipHistorySelect](ctx, fhs.FriendshipHistoryQuery, fhs, fhs.inters, v)
}

func (fhs *FriendshipHistorySelect) sqlScan(ctx context.Context, root *FriendshipHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fhs.fns))
	for _, fn := range fhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
