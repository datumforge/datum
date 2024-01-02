// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next generated.Querier) generated.Querier {
	return generated.QuerierFunc(func(ctx context.Context, q generated.Query) (generated.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q generated.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The EmailVerificationTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type EmailVerificationTokenFunc func(context.Context, *generated.EmailVerificationTokenQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f EmailVerificationTokenFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.EmailVerificationTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.EmailVerificationTokenQuery", q)
}

// The TraverseEmailVerificationToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEmailVerificationToken func(context.Context, *generated.EmailVerificationTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEmailVerificationToken) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEmailVerificationToken) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.EmailVerificationTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.EmailVerificationTokenQuery", q)
}

// The EntitlementFunc type is an adapter to allow the use of ordinary function as a Querier.
type EntitlementFunc func(context.Context, *generated.EntitlementQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f EntitlementFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.EntitlementQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.EntitlementQuery", q)
}

// The TraverseEntitlement type is an adapter to allow the use of ordinary function as Traverser.
type TraverseEntitlement func(context.Context, *generated.EntitlementQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseEntitlement) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseEntitlement) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.EntitlementQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.EntitlementQuery", q)
}

// The GroupFunc type is an adapter to allow the use of ordinary function as a Querier.
type GroupFunc func(context.Context, *generated.GroupQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f GroupFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.GroupQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.GroupQuery", q)
}

// The TraverseGroup type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGroup func(context.Context, *generated.GroupQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGroup) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGroup) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.GroupQuery", q)
}

// The GroupSettingFunc type is an adapter to allow the use of ordinary function as a Querier.
type GroupSettingFunc func(context.Context, *generated.GroupSettingQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f GroupSettingFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.GroupSettingQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.GroupSettingQuery", q)
}

// The TraverseGroupSetting type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGroupSetting func(context.Context, *generated.GroupSettingQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGroupSetting) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGroupSetting) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupSettingQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.GroupSettingQuery", q)
}

// The IntegrationFunc type is an adapter to allow the use of ordinary function as a Querier.
type IntegrationFunc func(context.Context, *generated.IntegrationQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f IntegrationFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.IntegrationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.IntegrationQuery", q)
}

// The TraverseIntegration type is an adapter to allow the use of ordinary function as Traverser.
type TraverseIntegration func(context.Context, *generated.IntegrationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseIntegration) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseIntegration) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.IntegrationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.IntegrationQuery", q)
}

// The OauthProviderFunc type is an adapter to allow the use of ordinary function as a Querier.
type OauthProviderFunc func(context.Context, *generated.OauthProviderQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OauthProviderFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OauthProviderQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OauthProviderQuery", q)
}

// The TraverseOauthProvider type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOauthProvider func(context.Context, *generated.OauthProviderQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOauthProvider) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOauthProvider) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OauthProviderQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OauthProviderQuery", q)
}

// The OhAuthTooTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type OhAuthTooTokenFunc func(context.Context, *generated.OhAuthTooTokenQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OhAuthTooTokenFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OhAuthTooTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OhAuthTooTokenQuery", q)
}

// The TraverseOhAuthTooToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOhAuthTooToken func(context.Context, *generated.OhAuthTooTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOhAuthTooToken) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOhAuthTooToken) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OhAuthTooTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OhAuthTooTokenQuery", q)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrganizationFunc func(context.Context, *generated.OrganizationQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OrganizationFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OrganizationQuery", q)
}

// The TraverseOrganization type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrganization func(context.Context, *generated.OrganizationQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrganization) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrganization) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OrganizationQuery", q)
}

// The OrganizationSettingFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrganizationSettingFunc func(context.Context, *generated.OrganizationSettingQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OrganizationSettingFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OrganizationSettingQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OrganizationSettingQuery", q)
}

// The TraverseOrganizationSetting type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrganizationSetting func(context.Context, *generated.OrganizationSettingQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrganizationSetting) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrganizationSetting) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationSettingQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OrganizationSettingQuery", q)
}

// The PasswordResetTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type PasswordResetTokenFunc func(context.Context, *generated.PasswordResetTokenQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f PasswordResetTokenFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.PasswordResetTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.PasswordResetTokenQuery", q)
}

// The TraversePasswordResetToken type is an adapter to allow the use of ordinary function as Traverser.
type TraversePasswordResetToken func(context.Context, *generated.PasswordResetTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePasswordResetToken) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePasswordResetToken) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PasswordResetTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.PasswordResetTokenQuery", q)
}

// The PersonalAccessTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type PersonalAccessTokenFunc func(context.Context, *generated.PersonalAccessTokenQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f PersonalAccessTokenFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.PersonalAccessTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.PersonalAccessTokenQuery", q)
}

// The TraversePersonalAccessToken type is an adapter to allow the use of ordinary function as Traverser.
type TraversePersonalAccessToken func(context.Context, *generated.PersonalAccessTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraversePersonalAccessToken) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraversePersonalAccessToken) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.PersonalAccessTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.PersonalAccessTokenQuery", q)
}

// The SessionFunc type is an adapter to allow the use of ordinary function as a Querier.
type SessionFunc func(context.Context, *generated.SessionQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f SessionFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.SessionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.SessionQuery", q)
}

// The TraverseSession type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSession func(context.Context, *generated.SessionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSession) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSession) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.SessionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.SessionQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *generated.UserQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *generated.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.UserQuery", q)
}

// The UserSettingFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserSettingFunc func(context.Context, *generated.UserSettingQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f UserSettingFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.UserSettingQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.UserSettingQuery", q)
}

// The TraverseUserSetting type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUserSetting func(context.Context, *generated.UserSettingQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUserSetting) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUserSetting) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.UserSettingQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.UserSettingQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q generated.Query) (Query, error) {
	switch q := q.(type) {
	case *generated.EmailVerificationTokenQuery:
		return &query[*generated.EmailVerificationTokenQuery, predicate.EmailVerificationToken, emailverificationtoken.OrderOption]{typ: generated.TypeEmailVerificationToken, tq: q}, nil
	case *generated.EntitlementQuery:
		return &query[*generated.EntitlementQuery, predicate.Entitlement, entitlement.OrderOption]{typ: generated.TypeEntitlement, tq: q}, nil
	case *generated.GroupQuery:
		return &query[*generated.GroupQuery, predicate.Group, group.OrderOption]{typ: generated.TypeGroup, tq: q}, nil
	case *generated.GroupSettingQuery:
		return &query[*generated.GroupSettingQuery, predicate.GroupSetting, groupsetting.OrderOption]{typ: generated.TypeGroupSetting, tq: q}, nil
	case *generated.IntegrationQuery:
		return &query[*generated.IntegrationQuery, predicate.Integration, integration.OrderOption]{typ: generated.TypeIntegration, tq: q}, nil
	case *generated.OauthProviderQuery:
		return &query[*generated.OauthProviderQuery, predicate.OauthProvider, oauthprovider.OrderOption]{typ: generated.TypeOauthProvider, tq: q}, nil
	case *generated.OhAuthTooTokenQuery:
		return &query[*generated.OhAuthTooTokenQuery, predicate.OhAuthTooToken, ohauthtootoken.OrderOption]{typ: generated.TypeOhAuthTooToken, tq: q}, nil
	case *generated.OrganizationQuery:
		return &query[*generated.OrganizationQuery, predicate.Organization, organization.OrderOption]{typ: generated.TypeOrganization, tq: q}, nil
	case *generated.OrganizationSettingQuery:
		return &query[*generated.OrganizationSettingQuery, predicate.OrganizationSetting, organizationsetting.OrderOption]{typ: generated.TypeOrganizationSetting, tq: q}, nil
	case *generated.PasswordResetTokenQuery:
		return &query[*generated.PasswordResetTokenQuery, predicate.PasswordResetToken, passwordresettoken.OrderOption]{typ: generated.TypePasswordResetToken, tq: q}, nil
	case *generated.PersonalAccessTokenQuery:
		return &query[*generated.PersonalAccessTokenQuery, predicate.PersonalAccessToken, personalaccesstoken.OrderOption]{typ: generated.TypePersonalAccessToken, tq: q}, nil
	case *generated.SessionQuery:
		return &query[*generated.SessionQuery, predicate.Session, session.OrderOption]{typ: generated.TypeSession, tq: q}, nil
	case *generated.UserQuery:
		return &query[*generated.UserQuery, predicate.User, user.OrderOption]{typ: generated.TypeUser, tq: q}, nil
	case *generated.UserSettingQuery:
		return &query[*generated.UserSettingQuery, predicate.UserSetting, usersetting.OrderOption]{typ: generated.TypeUserSetting, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
