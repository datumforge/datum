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

// The GroupSettingsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GroupSettingsQueryRuleFunc func(context.Context, *generated.GroupSettingsQuery) error

// EvalQuery return f(ctx, q).
func (f GroupSettingsQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupSettingsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.GroupSettingsQuery", q)
}

// The GroupSettingsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GroupSettingsMutationRuleFunc func(context.Context, *generated.GroupSettingsMutation) error

// EvalMutation calls f(ctx, m).
func (f GroupSettingsMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.GroupSettingsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.GroupSettingsMutation", m)
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

// The OrganizationSettingsQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type OrganizationSettingsQueryRuleFunc func(context.Context, *generated.OrganizationSettingsQuery) error

// EvalQuery return f(ctx, q).
func (f OrganizationSettingsQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationSettingsQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.OrganizationSettingsQuery", q)
}

// The OrganizationSettingsMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type OrganizationSettingsMutationRuleFunc func(context.Context, *generated.OrganizationSettingsMutation) error

// EvalMutation calls f(ctx, m).
func (f OrganizationSettingsMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.OrganizationSettingsMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.OrganizationSettingsMutation", m)
}

// The RefreshTokenQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RefreshTokenQueryRuleFunc func(context.Context, *generated.RefreshTokenQuery) error

// EvalQuery return f(ctx, q).
func (f RefreshTokenQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.RefreshTokenQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.RefreshTokenQuery", q)
}

// The RefreshTokenMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RefreshTokenMutationRuleFunc func(context.Context, *generated.RefreshTokenMutation) error

// EvalMutation calls f(ctx, m).
func (f RefreshTokenMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.RefreshTokenMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.RefreshTokenMutation", m)
}

// The SessionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SessionQueryRuleFunc func(context.Context, *generated.SessionQuery) error

// EvalQuery return f(ctx, q).
func (f SessionQueryRuleFunc) EvalQuery(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.SessionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("generated/privacy: unexpected query type %T, expect *generated.SessionQuery", q)
}

// The SessionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SessionMutationRuleFunc func(context.Context, *generated.SessionMutation) error

// EvalMutation calls f(ctx, m).
func (f SessionMutationRuleFunc) EvalMutation(ctx context.Context, m generated.Mutation) error {
	if m, ok := m.(*generated.SessionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("generated/privacy: unexpected mutation type %T, expect *generated.SessionMutation", m)
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
	case *generated.GroupQuery:
		return q.Filter(), nil
	case *generated.GroupSettingsQuery:
		return q.Filter(), nil
	case *generated.IntegrationQuery:
		return q.Filter(), nil
	case *generated.OrganizationQuery:
		return q.Filter(), nil
	case *generated.OrganizationSettingsQuery:
		return q.Filter(), nil
	case *generated.RefreshTokenQuery:
		return q.Filter(), nil
	case *generated.SessionQuery:
		return q.Filter(), nil
	case *generated.UserQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m generated.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *generated.GroupMutation:
		return m.Filter(), nil
	case *generated.GroupSettingsMutation:
		return m.Filter(), nil
	case *generated.IntegrationMutation:
		return m.Filter(), nil
	case *generated.OrganizationMutation:
		return m.Filter(), nil
	case *generated.OrganizationSettingsMutation:
		return m.Filter(), nil
	case *generated.RefreshTokenMutation:
		return m.Filter(), nil
	case *generated.SessionMutation:
		return m.Filter(), nil
	case *generated.UserMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("generated/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
