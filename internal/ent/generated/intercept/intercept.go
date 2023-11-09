// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettings"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/refreshtoken"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
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

// The GroupSettingsFunc type is an adapter to allow the use of ordinary function as a Querier.
type GroupSettingsFunc func(context.Context, *generated.GroupSettingsQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f GroupSettingsFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.GroupSettingsQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.GroupSettingsQuery", q)
}

// The TraverseGroupSettings type is an adapter to allow the use of ordinary function as Traverser.
type TraverseGroupSettings func(context.Context, *generated.GroupSettingsQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseGroupSettings) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseGroupSettings) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.GroupSettingsQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.GroupSettingsQuery", q)
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

// The OrganizationSettingsFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrganizationSettingsFunc func(context.Context, *generated.OrganizationSettingsQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f OrganizationSettingsFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.OrganizationSettingsQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.OrganizationSettingsQuery", q)
}

// The TraverseOrganizationSettings type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrganizationSettings func(context.Context, *generated.OrganizationSettingsQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrganizationSettings) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrganizationSettings) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.OrganizationSettingsQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.OrganizationSettingsQuery", q)
}

// The RefreshTokenFunc type is an adapter to allow the use of ordinary function as a Querier.
type RefreshTokenFunc func(context.Context, *generated.RefreshTokenQuery) (generated.Value, error)

// Query calls f(ctx, q).
func (f RefreshTokenFunc) Query(ctx context.Context, q generated.Query) (generated.Value, error) {
	if q, ok := q.(*generated.RefreshTokenQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *generated.RefreshTokenQuery", q)
}

// The TraverseRefreshToken type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRefreshToken func(context.Context, *generated.RefreshTokenQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRefreshToken) Intercept(next generated.Querier) generated.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRefreshToken) Traverse(ctx context.Context, q generated.Query) error {
	if q, ok := q.(*generated.RefreshTokenQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *generated.RefreshTokenQuery", q)
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

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q generated.Query) (Query, error) {
	switch q := q.(type) {
	case *generated.EntitlementQuery:
		return &query[*generated.EntitlementQuery, predicate.Entitlement, entitlement.OrderOption]{typ: generated.TypeEntitlement, tq: q}, nil
	case *generated.GroupQuery:
		return &query[*generated.GroupQuery, predicate.Group, group.OrderOption]{typ: generated.TypeGroup, tq: q}, nil
	case *generated.GroupSettingsQuery:
		return &query[*generated.GroupSettingsQuery, predicate.GroupSettings, groupsettings.OrderOption]{typ: generated.TypeGroupSettings, tq: q}, nil
	case *generated.IntegrationQuery:
		return &query[*generated.IntegrationQuery, predicate.Integration, integration.OrderOption]{typ: generated.TypeIntegration, tq: q}, nil
	case *generated.OrganizationQuery:
		return &query[*generated.OrganizationQuery, predicate.Organization, organization.OrderOption]{typ: generated.TypeOrganization, tq: q}, nil
	case *generated.OrganizationSettingsQuery:
		return &query[*generated.OrganizationSettingsQuery, predicate.OrganizationSettings, organizationsettings.OrderOption]{typ: generated.TypeOrganizationSettings, tq: q}, nil
	case *generated.RefreshTokenQuery:
		return &query[*generated.RefreshTokenQuery, predicate.RefreshToken, refreshtoken.OrderOption]{typ: generated.TypeRefreshToken, tq: q}, nil
	case *generated.SessionQuery:
		return &query[*generated.SessionQuery, predicate.Session, session.OrderOption]{typ: generated.TypeSession, tq: q}, nil
	case *generated.UserQuery:
		return &query[*generated.UserQuery, predicate.User, user.OrderOption]{typ: generated.TypeUser, tq: q}, nil
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
