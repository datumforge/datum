package interceptors

import (
	"context"

	"entgo.io/ent"
	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorSubscriber is middleware to change the Subscriber query
func InterceptorSubscriber() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.SubscriberFunc(func(ctx context.Context, q *generated.SubscriberQuery) (generated.Value, error) {
			// run the query
			v, err := next.Query(ctx, q)
			if err != nil {
				return nil, err
			}

			// bypass checks on subscriber creation
			if rule.ContextHasPrivacyTokenOfType(ctx, &token.VerifyToken{}) || rule.ContextHasPrivacyTokenOfType(ctx, &token.SignUpToken{}) {
				return v, nil
			}

			return filterSubscribersByAccess(ctx, q, v)
		})
	})
}

// filterSubscribersByAccess checks fga, using ListObjects, and ensure user has view access to a owning org before it is returned
func filterSubscribersByAccess(ctx context.Context, q *generated.SubscriberQuery, v ent.Value) ([]*generated.Subscriber, error) {
	q.Logger.Debugw("intercepting list subscriber query")

	// return early if no subscribers
	if v == nil {
		return nil, nil
	}

	qc := ent.QueryFromContext(ctx)

	var (
		subscribers []*generated.Subscriber
		err         error
	)

	// check if query is for an exists query, which returns a slice of subscriber ids
	// instead of the subscriber objects
	if qc.Op == ExistOperation {
		subs, ok := v.([]string)
		if !ok {
			q.Logger.Errorw("unexpected type for subscriber exist query")

			return nil, ErrInternalServerError
		}

		// return early if no subscribers
		if len(subs) == 0 {
			return nil, nil
		}

		// get the full subscriber objects to get the OwnerID
		subscribers, err = q.Where(subscriber.IDIn(subs...)).All(ctx)
		if err != nil {
			q.Logger.Errorw("unable to get subscribers", "error", err)

			return nil, err
		}
	} else {
		var ok bool

		subscribers, ok = v.([]*generated.Subscriber)
		if !ok {
			q.Logger.Infow("unexpected type for subscriber query, will continue without filtering")

			return nil, nil
		}
	}

	// get userID for tuple checks
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		q.Logger.Errorw("unable to get user id from echo context")
		return nil, err
	}

	// See all orgs user has view access
	orgList, err := q.Authz.ListObjectsRequest(ctx, userID, auth.GetAuthzSubjectType(ctx), "organization", fgax.CanView)
	if err != nil {
		return nil, err
	}

	// Check for system admin
	isAdmin, err := q.Authz.CheckSystemAdminRole(ctx, userID)
	if err != nil {
		return nil, err
	}

	userOrgs := orgList.GetObjects()

	var accessibleSubscribers []*generated.Subscriber

	for _, s := range subscribers {
		if s.OwnerID != "" {
			entityType := "organization"

			if !fgax.ListContains(entityType, userOrgs, s.OwnerID) {
				q.Logger.Debugw("access denied to organization", "relation", fgax.CanView, "organization", s.OwnerID, "type", entityType)

				continue
			}
			// add subscriber to accessible subscribers
			accessibleSubscribers = append(accessibleSubscribers, s)
		} else if isAdmin {
			// add subscriber to accessible subscribers
			accessibleSubscribers = append(accessibleSubscribers, s)
		}
	}

	// return updated subscribers
	return accessibleSubscribers, nil
}
