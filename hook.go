package enthistory

import (
	"context"
	"fmt"

	"entgo.io/ent"
)

type Mutation interface {
	Op() ent.Op
	CreateHistoryFromCreate(ctx context.Context) error
	CreateHistoryFromUpdate(ctx context.Context) error
	CreateHistoryFromDelete(ctx context.Context) error
}

type Mutator interface {
	Mutate(context.Context, Mutation) (ent.Value, error)
}

func On(hk ent.Hook, op ent.Op) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			hasOp := m.Op().Is(op)
			if hasOp {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

func HistoryTriggerInsert[T Mutation]() ent.Hook {
	return On(historyHookCreate[T](), ent.OpCreate)
}

func HistoryTriggerUpdate[T Mutation]() ent.Hook {
	return On(historyHookUpdate[T](), ent.OpUpdate|ent.OpUpdateOne)
}

func HistoryTriggerDelete[T Mutation]() ent.Hook {
	return On(historyHookDelete[T](), ent.OpDelete|ent.OpDeleteOne)
}

func getTypedMutation[T Mutation](m ent.Mutation) (T, error) {
	f, ok := any(m).(T)
	if !ok {
		return f, fmt.Errorf("expected appropriately typed mutation in schema hook, got: %+v", m)
	}
	return f, nil
}

func historyHookCreate[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			value, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			err = mutation.CreateHistoryFromCreate(ctx)
			if err != nil {
				return nil, err
			}

			return value, nil
		})
	}
}

func historyHookUpdate[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			err = mutation.CreateHistoryFromUpdate(ctx)
			if err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}

func historyHookDelete[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			err = mutation.CreateHistoryFromDelete(ctx)
			if err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}
