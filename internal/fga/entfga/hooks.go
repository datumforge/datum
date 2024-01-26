package entfga

import (
	"context"
	"fmt"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/entx"
)

type Mutation interface {
	Op() ent.Op
	CreateTuplesFromCreate(ctx context.Context) error
	CreateTuplesFromUpdate(ctx context.Context) error
	CreateTuplesFromDelete(ctx context.Context) error
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

func AuthzHooks[T Mutation]() []ent.Hook {
	return []ent.Hook{
		On(authzHookCreate[T](), ent.OpCreate),
		On(authzHookUpdate[T](), ent.OpUpdate|ent.OpUpdateOne),
		On(authzHookDelete[T](), ent.OpDelete|ent.OpDeleteOne),
	}
}

func getTypedMutation[T Mutation](m ent.Mutation) (T, error) {
	f, ok := any(m).(T)
	if !ok {
		return f, fmt.Errorf("expected appropriately typed mutation in schema hook, got: %+v", m) //nolint:goerr113
	}

	return f, nil
}

func authzHookCreate[T Mutation]() ent.Hook {
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

			err = mutation.CreateTuplesFromCreate(ctx)
			if err != nil {
				return nil, err
			}

			return value, nil
		})
	}
}

func authzHookUpdate[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			if entx.CheckIsSoftDelete(ctx) {
				return next.Mutate(ctx, m)
			}

			if err = mutation.CreateTuplesFromUpdate(ctx); err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}

func authzHookDelete[T Mutation]() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mutation, err := getTypedMutation[T](m)
			if err != nil {
				return nil, err
			}

			if err = mutation.CreateTuplesFromDelete(ctx); err != nil {
				return nil, err
			}

			return next.Mutate(ctx, m)
		})
	}
}
