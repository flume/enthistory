// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/predicate"
	"_examples/graphql/ent/testskiphistory"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TestSkipHistoryQuery is the builder for querying TestSkipHistory entities.
type TestSkipHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []testskiphistory.OrderOption
	inters     []Interceptor
	predicates []predicate.TestSkipHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*TestSkipHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestSkipHistoryQuery builder.
func (tshq *TestSkipHistoryQuery) Where(ps ...predicate.TestSkipHistory) *TestSkipHistoryQuery {
	tshq.predicates = append(tshq.predicates, ps...)
	return tshq
}

// Limit the number of records to be returned by this query.
func (tshq *TestSkipHistoryQuery) Limit(limit int) *TestSkipHistoryQuery {
	tshq.ctx.Limit = &limit
	return tshq
}

// Offset to start from.
func (tshq *TestSkipHistoryQuery) Offset(offset int) *TestSkipHistoryQuery {
	tshq.ctx.Offset = &offset
	return tshq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tshq *TestSkipHistoryQuery) Unique(unique bool) *TestSkipHistoryQuery {
	tshq.ctx.Unique = &unique
	return tshq
}

// Order specifies how the records should be ordered.
func (tshq *TestSkipHistoryQuery) Order(o ...testskiphistory.OrderOption) *TestSkipHistoryQuery {
	tshq.order = append(tshq.order, o...)
	return tshq
}

// First returns the first TestSkipHistory entity from the query.
// Returns a *NotFoundError when no TestSkipHistory was found.
func (tshq *TestSkipHistoryQuery) First(ctx context.Context) (*TestSkipHistory, error) {
	nodes, err := tshq.Limit(1).All(setContextOp(ctx, tshq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testskiphistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) FirstX(ctx context.Context) *TestSkipHistory {
	node, err := tshq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestSkipHistory ID from the query.
// Returns a *NotFoundError when no TestSkipHistory ID was found.
func (tshq *TestSkipHistoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tshq.Limit(1).IDs(setContextOp(ctx, tshq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testskiphistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := tshq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestSkipHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestSkipHistory entity is found.
// Returns a *NotFoundError when no TestSkipHistory entities are found.
func (tshq *TestSkipHistoryQuery) Only(ctx context.Context) (*TestSkipHistory, error) {
	nodes, err := tshq.Limit(2).All(setContextOp(ctx, tshq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testskiphistory.Label}
	default:
		return nil, &NotSingularError{testskiphistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) OnlyX(ctx context.Context) *TestSkipHistory {
	node, err := tshq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestSkipHistory ID in the query.
// Returns a *NotSingularError when more than one TestSkipHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (tshq *TestSkipHistoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = tshq.Limit(2).IDs(setContextOp(ctx, tshq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testskiphistory.Label}
	default:
		err = &NotSingularError{testskiphistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := tshq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestSkipHistories.
func (tshq *TestSkipHistoryQuery) All(ctx context.Context) ([]*TestSkipHistory, error) {
	ctx = setContextOp(ctx, tshq.ctx, "All")
	if err := tshq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestSkipHistory, *TestSkipHistoryQuery]()
	return withInterceptors[[]*TestSkipHistory](ctx, tshq, qr, tshq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) AllX(ctx context.Context) []*TestSkipHistory {
	nodes, err := tshq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestSkipHistory IDs.
func (tshq *TestSkipHistoryQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if tshq.ctx.Unique == nil && tshq.path != nil {
		tshq.Unique(true)
	}
	ctx = setContextOp(ctx, tshq.ctx, "IDs")
	if err = tshq.Select(testskiphistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := tshq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tshq *TestSkipHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tshq.ctx, "Count")
	if err := tshq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tshq, querierCount[*TestSkipHistoryQuery](), tshq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) CountX(ctx context.Context) int {
	count, err := tshq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tshq *TestSkipHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tshq.ctx, "Exist")
	switch _, err := tshq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tshq *TestSkipHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := tshq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestSkipHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tshq *TestSkipHistoryQuery) Clone() *TestSkipHistoryQuery {
	if tshq == nil {
		return nil
	}
	return &TestSkipHistoryQuery{
		config:     tshq.config,
		ctx:        tshq.ctx.Clone(),
		order:      append([]testskiphistory.OrderOption{}, tshq.order...),
		inters:     append([]Interceptor{}, tshq.inters...),
		predicates: append([]predicate.TestSkipHistory{}, tshq.predicates...),
		// clone intermediate query.
		sql:  tshq.sql.Clone(),
		path: tshq.path,
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
//	client.TestSkipHistory.Query().
//		GroupBy(testskiphistory.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tshq *TestSkipHistoryQuery) GroupBy(field string, fields ...string) *TestSkipHistoryGroupBy {
	tshq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestSkipHistoryGroupBy{build: tshq}
	grbuild.flds = &tshq.ctx.Fields
	grbuild.label = testskiphistory.Label
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
//	client.TestSkipHistory.Query().
//		Select(testskiphistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (tshq *TestSkipHistoryQuery) Select(fields ...string) *TestSkipHistorySelect {
	tshq.ctx.Fields = append(tshq.ctx.Fields, fields...)
	sbuild := &TestSkipHistorySelect{TestSkipHistoryQuery: tshq}
	sbuild.label = testskiphistory.Label
	sbuild.flds, sbuild.scan = &tshq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestSkipHistorySelect configured with the given aggregations.
func (tshq *TestSkipHistoryQuery) Aggregate(fns ...AggregateFunc) *TestSkipHistorySelect {
	return tshq.Select().Aggregate(fns...)
}

func (tshq *TestSkipHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tshq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tshq); err != nil {
				return err
			}
		}
	}
	for _, f := range tshq.ctx.Fields {
		if !testskiphistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tshq.path != nil {
		prev, err := tshq.path(ctx)
		if err != nil {
			return err
		}
		tshq.sql = prev
	}
	return nil
}

func (tshq *TestSkipHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestSkipHistory, error) {
	var (
		nodes = []*TestSkipHistory{}
		_spec = tshq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestSkipHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestSkipHistory{config: tshq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tshq.modifiers) > 0 {
		_spec.Modifiers = tshq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tshq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range tshq.loadTotal {
		if err := tshq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tshq *TestSkipHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tshq.querySpec()
	if len(tshq.modifiers) > 0 {
		_spec.Modifiers = tshq.modifiers
	}
	_spec.Node.Columns = tshq.ctx.Fields
	if len(tshq.ctx.Fields) > 0 {
		_spec.Unique = tshq.ctx.Unique != nil && *tshq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tshq.driver, _spec)
}

func (tshq *TestSkipHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testskiphistory.Table, testskiphistory.Columns, sqlgraph.NewFieldSpec(testskiphistory.FieldID, field.TypeUUID))
	_spec.From = tshq.sql
	if unique := tshq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tshq.path != nil {
		_spec.Unique = true
	}
	if fields := tshq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testskiphistory.FieldID)
		for i := range fields {
			if fields[i] != testskiphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tshq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tshq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tshq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tshq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tshq *TestSkipHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tshq.driver.Dialect())
	t1 := builder.Table(testskiphistory.Table)
	columns := tshq.ctx.Fields
	if len(columns) == 0 {
		columns = testskiphistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tshq.sql != nil {
		selector = tshq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tshq.ctx.Unique != nil && *tshq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tshq.predicates {
		p(selector)
	}
	for _, p := range tshq.order {
		p(selector)
	}
	if offset := tshq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tshq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestSkipHistoryGroupBy is the group-by builder for TestSkipHistory entities.
type TestSkipHistoryGroupBy struct {
	selector
	build *TestSkipHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tshgb *TestSkipHistoryGroupBy) Aggregate(fns ...AggregateFunc) *TestSkipHistoryGroupBy {
	tshgb.fns = append(tshgb.fns, fns...)
	return tshgb
}

// Scan applies the selector query and scans the result into the given value.
func (tshgb *TestSkipHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tshgb.build.ctx, "GroupBy")
	if err := tshgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestSkipHistoryQuery, *TestSkipHistoryGroupBy](ctx, tshgb.build, tshgb, tshgb.build.inters, v)
}

func (tshgb *TestSkipHistoryGroupBy) sqlScan(ctx context.Context, root *TestSkipHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tshgb.fns))
	for _, fn := range tshgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tshgb.flds)+len(tshgb.fns))
		for _, f := range *tshgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tshgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tshgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestSkipHistorySelect is the builder for selecting fields of TestSkipHistory entities.
type TestSkipHistorySelect struct {
	*TestSkipHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tshs *TestSkipHistorySelect) Aggregate(fns ...AggregateFunc) *TestSkipHistorySelect {
	tshs.fns = append(tshs.fns, fns...)
	return tshs
}

// Scan applies the selector query and scans the result into the given value.
func (tshs *TestSkipHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tshs.ctx, "Select")
	if err := tshs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestSkipHistoryQuery, *TestSkipHistorySelect](ctx, tshs.TestSkipHistoryQuery, tshs, tshs.inters, v)
}

func (tshs *TestSkipHistorySelect) sqlScan(ctx context.Context, root *TestSkipHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tshs.fns))
	for _, fn := range tshs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tshs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tshs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
