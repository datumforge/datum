// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// The EmailVerificationTokenFunc type is an adapter to allow the use of ordinary
// function as EmailVerificationToken mutator.
type EmailVerificationTokenFunc func(context.Context, *generated.EmailVerificationTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EmailVerificationTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EmailVerificationTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EmailVerificationTokenMutation", m)
}

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

// The GroupMembershipFunc type is an adapter to allow the use of ordinary
// function as GroupMembership mutator.
type GroupMembershipFunc func(context.Context, *generated.GroupMembershipMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupMembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupMembershipMutation", m)
}

// The GroupSettingFunc type is an adapter to allow the use of ordinary
// function as GroupSetting mutator.
type GroupSettingFunc func(context.Context, *generated.GroupSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupSettingMutation", m)
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

// The InviteFunc type is an adapter to allow the use of ordinary
// function as Invite mutator.
type InviteFunc func(context.Context, *generated.InviteMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f InviteFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.InviteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.InviteMutation", m)
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

// The OhAuthTooTokenFunc type is an adapter to allow the use of ordinary
// function as OhAuthTooToken mutator.
type OhAuthTooTokenFunc func(context.Context, *generated.OhAuthTooTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OhAuthTooTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OhAuthTooTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OhAuthTooTokenMutation", m)
}

// The OrgMembershipFunc type is an adapter to allow the use of ordinary
// function as OrgMembership mutator.
type OrgMembershipFunc func(context.Context, *generated.OrgMembershipMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrgMembershipFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrgMembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrgMembershipMutation", m)
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

// The OrganizationSettingFunc type is an adapter to allow the use of ordinary
// function as OrganizationSetting mutator.
type OrganizationSettingFunc func(context.Context, *generated.OrganizationSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationSettingMutation", m)
}

// The PasswordResetTokenFunc type is an adapter to allow the use of ordinary
// function as PasswordResetToken mutator.
type PasswordResetTokenFunc func(context.Context, *generated.PasswordResetTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f PasswordResetTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.PasswordResetTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.PasswordResetTokenMutation", m)
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

// The UserSettingFunc type is an adapter to allow the use of ordinary
// function as UserSetting mutator.
type UserSettingFunc func(context.Context, *generated.UserSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserSettingMutation", m)
}

// The WebauthnFunc type is an adapter to allow the use of ordinary
// function as Webauthn mutator.
type WebauthnFunc func(context.Context, *generated.WebauthnMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f WebauthnFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.WebauthnMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.WebauthnMutation", m)
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
