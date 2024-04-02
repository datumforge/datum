package interceptors

import (
	"context"

	"entgo.io/ent"
	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorTier is middleware to change the Tier query
func InterceptorTier() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.TierFunc(func(ctx context.Context, q *generated.TierQuery) (generated.Value, error) {
			// run the query
			v, err := next.Query(ctx, q)
			if err != nil {
				return nil, err
			}

			return filterTierByAccess(ctx, q, v)
		})
	})
}

// filterTierByAccess checks fga, using ListObjects, and ensure user has view access to a owning org before it is returned
func filterTierByAccess(ctx context.Context, q *generated.TierQuery, v ent.Value) ([]*generated.Tier, error) {
	q.Logger.Debugw("intercepting list tier query")

	tier, ok := v.([]*generated.Tier)
	if !ok {
		q.Logger.Infow("unexpected type for tier query, will continue without filtering")

		return nil, nil
	}

	// get userID for tuple checks
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		q.Logger.Errorw("unable to get user id from echo context")
		return nil, err
	}

	// See all orgs user has view access
	orgList, err := q.Authz.ListObjectsRequest(ctx, userID, "organization", fgax.CanView)
	if err != nil {
		return nil, err
	}

	// Check for system admin
	isAdmin, err := q.Authz.CheckSystemAdminRole(ctx, userID)
	if err != nil {
		return nil, err
	}

	userOrgs := orgList.GetObjects()

	var accessibleTier []*generated.Tier

	for _, s := range tier {
		if s.OwnerID != "" {
			entityType := "organization"

			if !fgax.ListContains(entityType, userOrgs, s.OwnerID) {
				q.Logger.Debugw("access denied to organization", "relation", fgax.CanView, "organization", s.OwnerID, "type", entityType)

				continue
			}
			// add tier to accessible tier
			accessibleTier = append(accessibleTier, s)
		} else if isAdmin {
			// add tier to accessible tier
			accessibleTier = append(accessibleTier, s)
		}
	}

	// return updated tier
	return accessibleTier, nil
}
