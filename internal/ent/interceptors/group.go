package interceptors

import (
	"context"
	"strings"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// InterceptorGroup is middleware to change the Group query
func InterceptorGroup() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.GroupFunc(func(ctx context.Context, q *generated.GroupQuery) (generated.Value, error) {
			// We only care these checks with authz is enabled, if this is empty skip interception checks
			if q.Authz.Ofga != nil {
				// run the query
				v, err := next.Query(ctx, q)
				if err != nil {
					return nil, err
				}

				return filterGroupsByAccess(ctx, q, v)
			}

			return next.Query(ctx, q)
		})
	})
}

// filterGroupsByAccess checks fga, using ListObjects, and ensure user has view access to a group before it is returned
func filterGroupsByAccess(ctx context.Context, q *generated.GroupQuery, v ent.Value) ([]*generated.Group, error) {
	q.Logger.Debugw("intercepting list group query")

	groups, ok := v.([]*generated.Group)
	if !ok {
		q.Logger.Errorw("unexpected type for group query")

		return nil, ErrInternalServerError
	}

	// get userID for tuple checks
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		q.Logger.Errorw("unable to get user id from echo context")
		return nil, err
	}

	// See all groups user has view access
	groupList, err := q.Authz.ListObjectsRequest(ctx, userID, "group", fga.CanView)
	if err != nil {
		return nil, err
	}

	userGroups := groupList.GetObjects()

	var accessibleGroups []*generated.Group

	for _, g := range groups {
		entityType := strings.ToLower(g.Update().Mutation().Type())

		if !fga.ListContains(entityType, userGroups, g.ID) {
			q.Logger.Infow("access denied to group", "relation", fga.CanView, "group_id", g.ID, "type", entityType)

			continue
		}

		// add group to accessible group
		accessibleGroups = append(accessibleGroups, g)
	}

	// return updated groups
	return accessibleGroups, nil
}
