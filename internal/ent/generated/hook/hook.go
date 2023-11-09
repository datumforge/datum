// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// The EntitlementFunc type is an adapter to allow the use of ordinary
// function as Entitlement mutator.
type EntitlementFunc func(context.Context, *generated.EntitlementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *generated.GroupMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupMutation", m)
}

// The GroupSettingsFunc type is an adapter to allow the use of ordinary
// function as GroupSettings mutator.
type GroupSettingsFunc func(context.Context, *generated.GroupSettingsMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupSettingsFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupSettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupSettingsMutation", m)
}

// The IntegrationFunc type is an adapter to allow the use of ordinary
// function as Integration mutator.
type IntegrationFunc func(context.Context, *generated.IntegrationMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f IntegrationFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.IntegrationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.IntegrationMutation", m)
}

// The OauthProviderFunc type is an adapter to allow the use of ordinary
// function as OauthProvider mutator.
type OauthProviderFunc func(context.Context, *generated.OauthProviderMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OauthProviderFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OauthProviderMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OauthProviderMutation", m)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary
// function as Organization mutator.
type OrganizationFunc func(context.Context, *generated.OrganizationMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationMutation", m)
}

// The OrganizationSettingsFunc type is an adapter to allow the use of ordinary
// function as OrganizationSettings mutator.
type OrganizationSettingsFunc func(context.Context, *generated.OrganizationSettingsMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationSettingsFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationSettingsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationSettingsMutation", m)
}

// The PersonalAccessTokenFunc type is an adapter to allow the use of ordinary
// function as PersonalAccessToken mutator.
type PersonalAccessTokenFunc func(context.Context, *generated.PersonalAccessTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f PersonalAccessTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.PersonalAccessTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.PersonalAccessTokenMutation", m)
}

// The RefreshTokenFunc type is an adapter to allow the use of ordinary
// function as RefreshToken mutator.
type RefreshTokenFunc func(context.Context, *generated.RefreshTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f RefreshTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.RefreshTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.RefreshTokenMutation", m)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *generated.SessionMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.SessionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.SessionMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *generated.UserMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, generated.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generated.Op) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk generated.Hook, cond Condition) generated.Hook {
	return func(next generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(ctx context.Context, m generated.Mutation) (generated.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generated.Delete|generated.Create)
func On(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generated.Update|generated.UpdateOne)
func Unless(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generated.Hook {
	return func(generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(context.Context, generated.Mutation) (generated.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generated.Hook {
//		return []generated.Hook{
//			Reject(generated.Delete|generated.Update),
//		}
//	}
func Reject(op generated.Op) generated.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generated.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generated.Hook) Chain {
	return Chain{append([]generated.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generated.Hook {
	return func(mutator generated.Mutator) generated.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generated.Hook) Chain {
	newHooks := make([]generated.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
