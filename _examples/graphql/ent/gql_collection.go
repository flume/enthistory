// Code generated by ent, DO NOT EDIT.

package ent

import (
	"_examples/graphql/ent/testexclude"
	"_examples/graphql/ent/todo"
	"_examples/graphql/ent/todohistory"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (te *TestExcludeQuery) CollectFields(ctx context.Context, satisfies ...string) (*TestExcludeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return te, nil
	}
	if err := te.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return te, nil
}

func (te *TestExcludeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(testexclude.Columns))
		selectedFields = []string{testexclude.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "otherID":
			if _, ok := fieldSeen[testexclude.FieldOtherID]; !ok {
				selectedFields = append(selectedFields, testexclude.FieldOtherID)
				fieldSeen[testexclude.FieldOtherID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[testexclude.FieldName]; !ok {
				selectedFields = append(selectedFields, testexclude.FieldName)
				fieldSeen[testexclude.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		te.Select(selectedFields...)
	}
	return nil
}

type testexcludePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TestExcludePaginateOption
}

func newTestExcludePaginateArgs(rv map[string]any) *testexcludePaginateArgs {
	args := &testexcludePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TestExcludeOrder{Field: &TestExcludeOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTestExcludeOrder(order))
			}
		case *TestExcludeOrder:
			if v != nil {
				args.opts = append(args.opts, WithTestExcludeOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TestExcludeWhereInput); ok {
		args.opts = append(args.opts, WithTestExcludeFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TodoQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TodoQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(todo.Columns))
		selectedFields = []string{todo.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "otherID":
			if _, ok := fieldSeen[todo.FieldOtherID]; !ok {
				selectedFields = append(selectedFields, todo.FieldOtherID)
				fieldSeen[todo.FieldOtherID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[todo.FieldName]; !ok {
				selectedFields = append(selectedFields, todo.FieldName)
				fieldSeen[todo.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type todoPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoPaginateOption
}

func newTodoPaginateArgs(rv map[string]any) *todoPaginateArgs {
	args := &todoPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TodoOrder{Field: &TodoOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTodoOrder(order))
			}
		case *TodoOrder:
			if v != nil {
				args.opts = append(args.opts, WithTodoOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TodoWhereInput); ok {
		args.opts = append(args.opts, WithTodoFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (th *TodoHistoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoHistoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return th, nil
	}
	if err := th.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return th, nil
}

func (th *TodoHistoryQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(todohistory.Columns))
		selectedFields = []string{todohistory.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "historyTime":
			if _, ok := fieldSeen[todohistory.FieldHistoryTime]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldHistoryTime)
				fieldSeen[todohistory.FieldHistoryTime] = struct{}{}
			}
		case "operation":
			if _, ok := fieldSeen[todohistory.FieldOperation]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldOperation)
				fieldSeen[todohistory.FieldOperation] = struct{}{}
			}
		case "ref":
			if _, ok := fieldSeen[todohistory.FieldRef]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldRef)
				fieldSeen[todohistory.FieldRef] = struct{}{}
			}
		case "updatedBy":
			if _, ok := fieldSeen[todohistory.FieldUpdatedBy]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldUpdatedBy)
				fieldSeen[todohistory.FieldUpdatedBy] = struct{}{}
			}
		case "otherID":
			if _, ok := fieldSeen[todohistory.FieldOtherID]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldOtherID)
				fieldSeen[todohistory.FieldOtherID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[todohistory.FieldName]; !ok {
				selectedFields = append(selectedFields, todohistory.FieldName)
				fieldSeen[todohistory.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		th.Select(selectedFields...)
	}
	return nil
}

type todohistoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoHistoryPaginateOption
}

func newTodoHistoryPaginateArgs(rv map[string]any) *todohistoryPaginateArgs {
	args := &todohistoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TodoHistoryOrder{Field: &TodoHistoryOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTodoHistoryOrder(order))
			}
		case *TodoHistoryOrder:
			if v != nil {
				args.opts = append(args.opts, WithTodoHistoryOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TodoHistoryWhereInput); ok {
		args.opts = append(args.opts, WithTodoHistoryFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
