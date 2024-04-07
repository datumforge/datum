// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, generated.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op generated.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op generated.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m generated.Mutation) error {
		return Denyf("generated/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The EmailVerificationTokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmailVerificationTokenQueryRuleFunc func(context.Context, *generated.EmailVerificationTokenQuery) error

// EvalQuery return f(ctx, q).
func (f EmailVerificationTokenQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.EmailVerificationTokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.EmailVerificationTokenQuery", q)
}

// The EmailVerificationTokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmailVerificationTokenMutationRuleFunc func(context.Context, *generated.EmailVerificationTokenMutation) error

// EvalMutation calls f(ctx, m).
func (f EmailVerificationTokenMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.EmailVerificationTokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.EmailVerificationTokenMutation", m)
}

// The EntitlementQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EntitlementQueryRuleFunc func(context.Context, *generated.EntitlementQuery) error

// EvalQuery return f(ctx, q).
func (f EntitlementQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.EntitlementQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.EntitlementQuery", q)
}

// The EntitlementMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EntitlementMutationRuleFunc func(context.Context, *generated.EntitlementMutation) error

// EvalMutation calls f(ctx, m).
func (f EntitlementMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.EntitlementMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.EntitlementMutation", m)
}

// The GroupQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupQueryRuleFunc func(context.Context, *generated.GroupQuery) error

// EvalQuery return f(ctx, q).
func (f GroupQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.GroupQuery", q)
}

// The GroupMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupMutationRuleFunc func(context.Context, *generated.GroupMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.GroupMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.GroupMutation", m)
}

// The GroupMembershipQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupMembershipQueryRuleFunc func(context.Context, *generated.GroupMembershipQuery) error

// EvalQuery return f(ctx, q).
func (f GroupMembershipQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupMembershipQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.GroupMembershipQuery", q)
}

// The GroupMembershipMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupMembershipMutationRuleFunc func(context.Context, *generated.GroupMembershipMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupMembershipMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.GroupMembershipMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.GroupMembershipMutation", m)
}

// The GroupSettingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupSettingQueryRuleFunc func(context.Context, *generated.GroupSettingQuery) error

// EvalQuery return f(ctx, q).
func (f GroupSettingQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupSettingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.GroupSettingQuery", q)
}

// The GroupSettingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupSettingMutationRuleFunc func(context.Context, *generated.GroupSettingMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupSettingMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.GroupSettingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.GroupSettingMutation", m)
}

// The IntegrationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type IntegrationQueryRuleFunc func(context.Context, *generated.IntegrationQuery) error

// EvalQuery return f(ctx, q).
func (f IntegrationQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.IntegrationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.IntegrationQuery", q)
}

// The IntegrationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type IntegrationMutationRuleFunc func(context.Context, *generated.IntegrationMutation) error

// EvalMutation calls f(ctx, m).
func (f IntegrationMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.IntegrationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.IntegrationMutation", m)
}

// The InviteQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type InviteQueryRuleFunc func(context.Context, *generated.InviteQuery) error

// EvalQuery return f(ctx, q).
func (f InviteQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.InviteQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.InviteQuery", q)
}

// The InviteMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type InviteMutationRuleFunc func(context.Context, *generated.InviteMutation) error

// EvalMutation calls f(ctx, m).
func (f InviteMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.InviteMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.InviteMutation", m)
}

// The OauthProviderQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OauthProviderQueryRuleFunc func(context.Context, *generated.OauthProviderQuery) error

// EvalQuery return f(ctx, q).
func (f OauthProviderQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OauthProviderQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OauthProviderQuery", q)
}

// The OauthProviderMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OauthProviderMutationRuleFunc func(context.Context, *generated.OauthProviderMutation) error

// EvalMutation calls f(ctx, m).
func (f OauthProviderMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OauthProviderMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OauthProviderMutation", m)
}

// The OhAuthTooTokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OhAuthTooTokenQueryRuleFunc func(context.Context, *generated.OhAuthTooTokenQuery) error

// EvalQuery return f(ctx, q).
func (f OhAuthTooTokenQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OhAuthTooTokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OhAuthTooTokenQuery", q)
}

// The OhAuthTooTokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OhAuthTooTokenMutationRuleFunc func(context.Context, *generated.OhAuthTooTokenMutation) error

// EvalMutation calls f(ctx, m).
func (f OhAuthTooTokenMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OhAuthTooTokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OhAuthTooTokenMutation", m)
}

// The OrgMembershipQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrgMembershipQueryRuleFunc func(context.Context, *generated.OrgMembershipQuery) error

// EvalQuery return f(ctx, q).
func (f OrgMembershipQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrgMembershipQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrgMembershipQuery", q)
}

// The OrgMembershipMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrgMembershipMutationRuleFunc func(context.Context, *generated.OrgMembershipMutation) error

// EvalMutation calls f(ctx, m).
func (f OrgMembershipMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrgMembershipMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrgMembershipMutation", m)
}

// The OrganizationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationQueryRuleFunc func(context.Context, *generated.OrganizationQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrganizationQuery", q)
}

// The OrganizationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationMutationRuleFunc func(context.Context, *generated.OrganizationMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrganizationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrganizationMutation", m)
}

// The OrganizationHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationHistoryQueryRuleFunc func(context.Context, *generated.OrganizationHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrganizationHistoryQuery", q)
}

// The OrganizationHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationHistoryMutationRuleFunc func(context.Context, *generated.OrganizationHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrganizationHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrganizationHistoryMutation", m)
}

// The OrganizationSettingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationSettingQueryRuleFunc func(context.Context, *generated.OrganizationSettingQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationSettingQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationSettingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrganizationSettingQuery", q)
}

// The OrganizationSettingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationSettingMutationRuleFunc func(context.Context, *generated.OrganizationSettingMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationSettingMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrganizationSettingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrganizationSettingMutation", m)
}

// The OrganizationSettingHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationSettingHistoryQueryRuleFunc func(context.Context, *generated.OrganizationSettingHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationSettingHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationSettingHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrganizationSettingHistoryQuery", q)
}

// The OrganizationSettingHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationSettingHistoryMutationRuleFunc func(context.Context, *generated.OrganizationSettingHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationSettingHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrganizationSettingHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrganizationSettingHistoryMutation", m)
}

// The PasswordResetTokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PasswordResetTokenQueryRuleFunc func(context.Context, *generated.PasswordResetTokenQuery) error

// EvalQuery return f(ctx, q).
func (f PasswordResetTokenQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PasswordResetTokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.PasswordResetTokenQuery", q)
}

// The PasswordResetTokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PasswordResetTokenMutationRuleFunc func(context.Context, *generated.PasswordResetTokenMutation) error

// EvalMutation calls f(ctx, m).
func (f PasswordResetTokenMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.PasswordResetTokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.PasswordResetTokenMutation", m)
}

// The PersonalAccessTokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PersonalAccessTokenQueryRuleFunc func(context.Context, *generated.PersonalAccessTokenQuery) error

// EvalQuery return f(ctx, q).
func (f PersonalAccessTokenQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PersonalAccessTokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.PersonalAccessTokenQuery", q)
}

// The PersonalAccessTokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PersonalAccessTokenMutationRuleFunc func(context.Context, *generated.PersonalAccessTokenMutation) error

// EvalMutation calls f(ctx, m).
func (f PersonalAccessTokenMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.PersonalAccessTokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.PersonalAccessTokenMutation", m)
}

// The SubscriberQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SubscriberQueryRuleFunc func(context.Context, *generated.SubscriberQuery) error

// EvalQuery return f(ctx, q).
func (f SubscriberQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.SubscriberQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.SubscriberQuery", q)
}

// The SubscriberMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SubscriberMutationRuleFunc func(context.Context, *generated.SubscriberMutation) error

// EvalMutation calls f(ctx, m).
func (f SubscriberMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.SubscriberMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.SubscriberMutation", m)
}

// The TFASettingsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TFASettingsQueryRuleFunc func(context.Context, *generated.TFASettingsQuery) error

// EvalQuery return f(ctx, q).
func (f TFASettingsQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.TFASettingsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.TFASettingsQuery", q)
}

// The TFASettingsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TFASettingsMutationRuleFunc func(context.Context, *generated.TFASettingsMutation) error

// EvalMutation calls f(ctx, m).
func (f TFASettingsMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.TFASettingsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.TFASettingsMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *generated.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *generated.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.UserMutation", m)
}

// The UserSettingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserSettingQueryRuleFunc func(context.Context, *generated.UserSettingQuery) error

// EvalQuery return f(ctx, q).
func (f UserSettingQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.UserSettingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.UserSettingQuery", q)
}

// The UserSettingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserSettingMutationRuleFunc func(context.Context, *generated.UserSettingMutation) error

// EvalMutation calls f(ctx, m).
func (f UserSettingMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.UserSettingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.UserSettingMutation", m)
}

// The WebauthnQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type WebauthnQueryRuleFunc func(context.Context, *generated.WebauthnQuery) error

// EvalQuery return f(ctx, q).
func (f WebauthnQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.WebauthnQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.WebauthnQuery", q)
}

// The WebauthnMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type WebauthnMutationRuleFunc func(context.Context, *generated.WebauthnMutation) error

// EvalMutation calls f(ctx, m).
func (f WebauthnMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.WebauthnMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.WebauthnMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q generated.Query) (Filter, error) {
	switch q := q.(type) {
	case *generated.EmailVerificationTokenQuery:
		return q.Filter(), nil
	case *generated.EntitlementQuery:
		return q.Filter(), nil
	case *generated.GroupQuery:
		return q.Filter(), nil
	case *generated.GroupMembershipQuery:
		return q.Filter(), nil
	case *generated.GroupSettingQuery:
		return q.Filter(), nil
	case *generated.IntegrationQuery:
		return q.Filter(), nil
	case *generated.InviteQuery:
		return q.Filter(), nil
	case *generated.OauthProviderQuery:
		return q.Filter(), nil
	case *generated.OhAuthTooTokenQuery:
		return q.Filter(), nil
	case *generated.OrgMembershipQuery:
		return q.Filter(), nil
	case *generated.OrganizationQuery:
		return q.Filter(), nil
	case *generated.OrganizationHistoryQuery:
		return q.Filter(), nil
	case *generated.OrganizationSettingQuery:
		return q.Filter(), nil
	case *generated.OrganizationSettingHistoryQuery:
		return q.Filter(), nil
	case *generated.PasswordResetTokenQuery:
		return q.Filter(), nil
	case *generated.PersonalAccessTokenQuery:
		return q.Filter(), nil
	case *generated.SubscriberQuery:
		return q.Filter(), nil
	case *generated.TFASettingsQuery:
		return q.Filter(), nil
	case *generated.UserQuery:
		return q.Filter(), nil
	case *generated.UserSettingQuery:
		return q.Filter(), nil
	case *generated.WebauthnQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m generated.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *generated.EmailVerificationTokenMutation:
		return m.Filter(), nil
	case *generated.EntitlementMutation:
		return m.Filter(), nil
	case *generated.GroupMutation:
		return m.Filter(), nil
	case *generated.GroupMembershipMutation:
		return m.Filter(), nil
	case *generated.GroupSettingMutation:
		return m.Filter(), nil
	case *generated.IntegrationMutation:
		return m.Filter(), nil
	case *generated.InviteMutation:
		return m.Filter(), nil
	case *generated.OauthProviderMutation:
		return m.Filter(), nil
	case *generated.OhAuthTooTokenMutation:
		return m.Filter(), nil
	case *generated.OrgMembershipMutation:
		return m.Filter(), nil
	case *generated.OrganizationMutation:
		return m.Filter(), nil
	case *generated.OrganizationHistoryMutation:
		return m.Filter(), nil
	case *generated.OrganizationSettingMutation:
		return m.Filter(), nil
	case *generated.OrganizationSettingHistoryMutation:
		return m.Filter(), nil
	case *generated.PasswordResetTokenMutation:
		return m.Filter(), nil
	case *generated.PersonalAccessTokenMutation:
		return m.Filter(), nil
	case *generated.SubscriberMutation:
		return m.Filter(), nil
	case *generated.TFASettingsMutation:
		return m.Filter(), nil
	case *generated.UserMutation:
		return m.Filter(), nil
	case *generated.UserSettingMutation:
		return m.Filter(), nil
	case *generated.WebauthnMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
