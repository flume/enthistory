// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/updateby_uuid/ent/organizationhistory"
	"_examples/updateby_uuid/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrganizationHistoryQuery is the builder for querying OrganizationHistory entities.
type OrganizationHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []organizationhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.OrganizationHistory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationHistoryQuery builder.
func (ohq *OrganizationHistoryQuery) Where(ps ...predicate.OrganizationHistory) *OrganizationHistoryQuery {
	ohq.predicates = append(ohq.predicates, ps...)
	return ohq
}

// Limit the number of records to be returned by this query.
func (ohq *OrganizationHistoryQuery) Limit(limit int) *OrganizationHistoryQuery {
	ohq.ctx.Limit = &limit
	return ohq
}

// Offset to start from.
func (ohq *OrganizationHistoryQuery) Offset(offset int) *OrganizationHistoryQuery {
	ohq.ctx.Offset = &offset
	return ohq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ohq *OrganizationHistoryQuery) Unique(unique bool) *OrganizationHistoryQuery {
	ohq.ctx.Unique = &unique
	return ohq
}

// Order specifies how the records should be ordered.
func (ohq *OrganizationHistoryQuery) Order(o ...organizationhistory.OrderOption) *OrganizationHistoryQuery {
	ohq.order = append(ohq.order, o...)
	return ohq
}

// First returns the first OrganizationHistory entity from the query.
// Returns a *NotFoundError when no OrganizationHistory was found.
func (ohq *OrganizationHistoryQuery) First(ctx context.Context) (*OrganizationHistory, error) {
	nodes, err := ohq.Limit(1).All(setContextOp(ctx, ohq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) FirstX(ctx context.Context) *OrganizationHistory {
	node, err := ohq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationHistory ID from the query.
// Returns a *NotFoundError when no OrganizationHistory ID was found.
func (ohq *OrganizationHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ohq.Limit(1).IDs(setContextOp(ctx, ohq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := ohq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationHistory entity is found.
// Returns a *NotFoundError when no OrganizationHistory entities are found.
func (ohq *OrganizationHistoryQuery) Only(ctx context.Context) (*OrganizationHistory, error) {
	nodes, err := ohq.Limit(2).All(setContextOp(ctx, ohq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationhistory.Label}
	default:
		return nil, &NotSingularError{organizationhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) OnlyX(ctx context.Context) *OrganizationHistory {
	node, err := ohq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationHistory ID in the query.
// Returns a *NotSingularError when more than one OrganizationHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (ohq *OrganizationHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ohq.Limit(2).IDs(setContextOp(ctx, ohq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationhistory.Label}
	default:
		err = &NotSingularError{organizationhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := ohq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationHistories.
func (ohq *OrganizationHistoryQuery) All(ctx context.Context) ([]*OrganizationHistory, error) {
	ctx = setContextOp(ctx, ohq.ctx, ent.OpQueryAll)
	if err := ohq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationHistory, *OrganizationHistoryQuery]()
	return withInterceptors[[]*OrganizationHistory](ctx, ohq, qr, ohq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) AllX(ctx context.Context) []*OrganizationHistory {
	nodes, err := ohq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationHistory IDs.
func (ohq *OrganizationHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ohq.ctx.Unique == nil && ohq.path != nil {
		ohq.Unique(true)
	}
	ctx = setContextOp(ctx, ohq.ctx, ent.OpQueryIDs)
	if err = ohq.Select(organizationhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := ohq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ohq *OrganizationHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ohq.ctx, ent.OpQueryCount)
	if err := ohq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ohq, querierCount[*OrganizationHistoryQuery](), ohq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) CountX(ctx context.Context) int {
	count, err := ohq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ohq *OrganizationHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ohq.ctx, ent.OpQueryExist)
	switch _, err := ohq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ohq *OrganizationHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := ohq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ohq *OrganizationHistoryQuery) Clone() *OrganizationHistoryQuery {
	if ohq == nil {
		return nil
	}
	return &OrganizationHistoryQuery{
		config:     ohq.config,
		ctx:        ohq.ctx.Clone(),
		order:      append([]organizationhistory.OrderOption{}, ohq.order...),
		inters:     append([]Interceptor{}, ohq.inters...),
		predicates: append([]predicate.OrganizationHistory{}, ohq.predicates...),
		// clone intermediate query.
		sql:  ohq.sql.Clone(),
		path: ohq.path,
	}
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
//	client.OrganizationHistory.Query().
//		GroupBy(organizationhistory.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ohq *OrganizationHistoryQuery) GroupBy(field string, fields ...string) *OrganizationHistoryGroupBy {
	ohq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationHistoryGroupBy{build: ohq}
	grbuild.flds = &ohq.ctx.Fields
	grbuild.label = organizationhistory.Label
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
//	client.OrganizationHistory.Query().
//		Select(organizationhistory.FieldCreatedAt).
//		Scan(ctx, &v)
func (ohq *OrganizationHistoryQuery) Select(fields ...string) *OrganizationHistorySelect {
	ohq.ctx.Fields = append(ohq.ctx.Fields, fields...)
	sbuild := &OrganizationHistorySelect{OrganizationHistoryQuery: ohq}
	sbuild.label = organizationhistory.Label
	sbuild.flds, sbuild.scan = &ohq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationHistorySelect configured with the given aggregations.
func (ohq *OrganizationHistoryQuery) Aggregate(fns ...AggregateFunc) *OrganizationHistorySelect {
	return ohq.Select().Aggregate(fns...)
}

func (ohq *OrganizationHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ohq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ohq); err != nil {
				return err
			}
		}
	}
	for _, f := range ohq.ctx.Fields {
		if !organizationhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ohq.path != nil {
		prev, err := ohq.path(ctx)
		if err != nil {
			return err
		}
		ohq.sql = prev
	}
	return nil
}

func (ohq *OrganizationHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationHistory, error) {
	var (
		nodes = []*OrganizationHistory{}
		_spec = ohq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationHistory{config: ohq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ohq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ohq *OrganizationHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ohq.querySpec()
	_spec.Node.Columns = ohq.ctx.Fields
	if len(ohq.ctx.Fields) > 0 {
		_spec.Unique = ohq.ctx.Unique != nil && *ohq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ohq.driver, _spec)
}

func (ohq *OrganizationHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationhistory.Table, organizationhistory.Columns, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeInt))
	_spec.From = ohq.sql
	if unique := ohq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ohq.path != nil {
		_spec.Unique = true
	}
	if fields := ohq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationhistory.FieldID)
		for i := range fields {
			if fields[i] != organizationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ohq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ohq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ohq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ohq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ohq *OrganizationHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ohq.driver.Dialect())
	t1 := builder.Table(organizationhistory.Table)
	columns := ohq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ohq.sql != nil {
		selector = ohq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ohq.ctx.Unique != nil && *ohq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ohq.predicates {
		p(selector)
	}
	for _, p := range ohq.order {
		p(selector)
	}
	if offset := ohq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ohq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrganizationHistoryGroupBy is the group-by builder for OrganizationHistory entities.
type OrganizationHistoryGroupBy struct {
	selector
	build *OrganizationHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ohgb *OrganizationHistoryGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationHistoryGroupBy {
	ohgb.fns = append(ohgb.fns, fns...)
	return ohgb
}

// Scan applies the selector query and scans the result into the given value.
func (ohgb *OrganizationHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ohgb.build.ctx, ent.OpQueryGroupBy)
	if err := ohgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationHistoryQuery, *OrganizationHistoryGroupBy](ctx, ohgb.build, ohgb, ohgb.build.inters, v)
}

func (ohgb *OrganizationHistoryGroupBy) sqlScan(ctx context.Context, root *OrganizationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ohgb.fns))
	for _, fn := range ohgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ohgb.flds)+len(ohgb.fns))
		for _, f := range *ohgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ohgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ohgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationHistorySelect is the builder for selecting fields of OrganizationHistory entities.
type OrganizationHistorySelect struct {
	*OrganizationHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ohs *OrganizationHistorySelect) Aggregate(fns ...AggregateFunc) *OrganizationHistorySelect {
	ohs.fns = append(ohs.fns, fns...)
	return ohs
}

// Scan applies the selector query and scans the result into the given value.
func (ohs *OrganizationHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ohs.ctx, ent.OpQuerySelect)
	if err := ohs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationHistoryQuery, *OrganizationHistorySelect](ctx, ohs.OrganizationHistoryQuery, ohs, ohs.inters, v)
}

func (ohs *OrganizationHistorySelect) sqlScan(ctx context.Context, root *OrganizationHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ohs.fns))
	for _, fn := range ohs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ohs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ohs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
