// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/todo"
	"_examples/graphql/ent/todohistory"
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[uuid.UUID]
	PageInfo       = entgql.PageInfo[uuid.UUID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// TodoEdge is the edge representation of Todo.
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TodoConnection is the connection containing edges to Todo.
type TodoConnection struct {
	Edges      []*TodoEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *TodoConnection) build(nodes []*Todo, pager *todoPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Todo
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Todo {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Todo {
			return nodes[i]
		}
	}
	c.Edges = make([]*TodoEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &TodoEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// TodoPaginateOption enables pagination customization.
type TodoPaginateOption func(*todoPager) error

// WithTodoOrder configures pagination ordering.
func WithTodoOrder(order *TodoOrder) TodoPaginateOption {
	if order == nil {
		order = DefaultTodoOrder
	}
	o := *order
	return func(pager *todoPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTodoOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTodoFilter configures pagination filter.
func WithTodoFilter(filter func(*TodoQuery) (*TodoQuery, error)) TodoPaginateOption {
	return func(pager *todoPager) error {
		if filter == nil {
			return errors.New("TodoQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type todoPager struct {
	reverse bool
	order   *TodoOrder
	filter  func(*TodoQuery) (*TodoQuery, error)
}

func newTodoPager(opts []TodoPaginateOption, reverse bool) (*todoPager, error) {
	pager := &todoPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTodoOrder
	}
	return pager, nil
}

func (p *todoPager) applyFilter(query *TodoQuery) (*TodoQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *todoPager) toCursor(t *Todo) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *todoPager) applyCursors(query *TodoQuery, after, before *Cursor) (*TodoQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultTodoOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *todoPager) applyOrder(query *TodoQuery) *TodoQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultTodoOrder.Field {
		query = query.Order(DefaultTodoOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *todoPager) orderExpr(query *TodoQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultTodoOrder.Field {
			b.Comma().Ident(DefaultTodoOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Todo.
func (t *TodoQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoPaginateOption,
) (*TodoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}
	conn := &TodoConnection{Edges: []*TodoEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = t.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if t, err = pager.applyCursors(t, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		t.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := t.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	t = pager.applyOrder(t)
	nodes, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// TodoOrderFieldName orders Todo by name.
	TodoOrderFieldName = &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.Name, nil
		},
		column: todo.FieldName,
		toTerm: todo.ByName,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f TodoOrderField) String() string {
	var str string
	switch f.column {
	case TodoOrderFieldName.column:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f TodoOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *TodoOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TodoOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *TodoOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid TodoOrderField", str)
	}
	return nil
}

// TodoOrderField defines the ordering field of Todo.
type TodoOrderField struct {
	// Value extracts the ordering value from the given Todo.
	Value    func(*Todo) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) todo.OrderOption
	toCursor func(*Todo) Cursor
}

// TodoOrder defines the ordering of Todo.
type TodoOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *TodoOrderField `json:"field"`
}

// DefaultTodoOrder is the default ordering of Todo.
var DefaultTodoOrder = &TodoOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.ID, nil
		},
		column: todo.FieldID,
		toTerm: todo.ByID,
		toCursor: func(t *Todo) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Todo into TodoEdge.
func (t *Todo) ToEdge(order *TodoOrder) *TodoEdge {
	if order == nil {
		order = DefaultTodoOrder
	}
	return &TodoEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// TodoHistoryEdge is the edge representation of TodoHistory.
type TodoHistoryEdge struct {
	Node   *TodoHistory `json:"node"`
	Cursor Cursor       `json:"cursor"`
}

// TodoHistoryConnection is the connection containing edges to TodoHistory.
type TodoHistoryConnection struct {
	Edges      []*TodoHistoryEdge `json:"edges"`
	PageInfo   PageInfo           `json:"pageInfo"`
	TotalCount int                `json:"totalCount"`
}

func (c *TodoHistoryConnection) build(nodes []*TodoHistory, pager *todohistoryPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *TodoHistory
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *TodoHistory {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *TodoHistory {
			return nodes[i]
		}
	}
	c.Edges = make([]*TodoHistoryEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &TodoHistoryEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// TodoHistoryPaginateOption enables pagination customization.
type TodoHistoryPaginateOption func(*todohistoryPager) error

// WithTodoHistoryOrder configures pagination ordering.
func WithTodoHistoryOrder(order *TodoHistoryOrder) TodoHistoryPaginateOption {
	if order == nil {
		order = DefaultTodoHistoryOrder
	}
	o := *order
	return func(pager *todohistoryPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTodoHistoryOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTodoHistoryFilter configures pagination filter.
func WithTodoHistoryFilter(filter func(*TodoHistoryQuery) (*TodoHistoryQuery, error)) TodoHistoryPaginateOption {
	return func(pager *todohistoryPager) error {
		if filter == nil {
			return errors.New("TodoHistoryQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type todohistoryPager struct {
	reverse bool
	order   *TodoHistoryOrder
	filter  func(*TodoHistoryQuery) (*TodoHistoryQuery, error)
}

func newTodoHistoryPager(opts []TodoHistoryPaginateOption, reverse bool) (*todohistoryPager, error) {
	pager := &todohistoryPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTodoHistoryOrder
	}
	return pager, nil
}

func (p *todohistoryPager) applyFilter(query *TodoHistoryQuery) (*TodoHistoryQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *todohistoryPager) toCursor(th *TodoHistory) Cursor {
	return p.order.Field.toCursor(th)
}

func (p *todohistoryPager) applyCursors(query *TodoHistoryQuery, after, before *Cursor) (*TodoHistoryQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultTodoHistoryOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *todohistoryPager) applyOrder(query *TodoHistoryQuery) *TodoHistoryQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultTodoHistoryOrder.Field {
		query = query.Order(DefaultTodoHistoryOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *todohistoryPager) orderExpr(query *TodoHistoryQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultTodoHistoryOrder.Field {
			b.Comma().Ident(DefaultTodoHistoryOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to TodoHistory.
func (th *TodoHistoryQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoHistoryPaginateOption,
) (*TodoHistoryConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoHistoryPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if th, err = pager.applyFilter(th); err != nil {
		return nil, err
	}
	conn := &TodoHistoryConnection{Edges: []*TodoHistoryEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = th.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if th, err = pager.applyCursors(th, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		th.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := th.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	th = pager.applyOrder(th)
	nodes, err := th.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// TodoHistoryOrderFieldName orders TodoHistory by name.
	TodoHistoryOrderFieldName = &TodoHistoryOrderField{
		Value: func(th *TodoHistory) (ent.Value, error) {
			return th.Name, nil
		},
		column: todohistory.FieldName,
		toTerm: todohistory.ByName,
		toCursor: func(th *TodoHistory) Cursor {
			return Cursor{
				ID:    th.ID,
				Value: th.Name,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f TodoHistoryOrderField) String() string {
	var str string
	switch f.column {
	case TodoHistoryOrderFieldName.column:
		str = "NAME"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f TodoHistoryOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *TodoHistoryOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TodoHistoryOrderField %T must be a string", v)
	}
	switch str {
	case "NAME":
		*f = *TodoHistoryOrderFieldName
	default:
		return fmt.Errorf("%s is not a valid TodoHistoryOrderField", str)
	}
	return nil
}

// TodoHistoryOrderField defines the ordering field of TodoHistory.
type TodoHistoryOrderField struct {
	// Value extracts the ordering value from the given TodoHistory.
	Value    func(*TodoHistory) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) todohistory.OrderOption
	toCursor func(*TodoHistory) Cursor
}

// TodoHistoryOrder defines the ordering of TodoHistory.
type TodoHistoryOrder struct {
	Direction OrderDirection         `json:"direction"`
	Field     *TodoHistoryOrderField `json:"field"`
}

// DefaultTodoHistoryOrder is the default ordering of TodoHistory.
var DefaultTodoHistoryOrder = &TodoHistoryOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &TodoHistoryOrderField{
		Value: func(th *TodoHistory) (ent.Value, error) {
			return th.ID, nil
		},
		column: todohistory.FieldID,
		toTerm: todohistory.ByID,
		toCursor: func(th *TodoHistory) Cursor {
			return Cursor{ID: th.ID}
		},
	},
}

// ToEdge converts TodoHistory into TodoHistoryEdge.
func (th *TodoHistory) ToEdge(order *TodoHistoryOrder) *TodoHistoryEdge {
	if order == nil {
		order = DefaultTodoHistoryOrder
	}
	return &TodoHistoryEdge{
		Node:   th,
		Cursor: order.Field.toCursor(th),
	}
}
