package interceptors

import (
	"context"

	"entgo.io/ent"
	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorGroup is middleware to change the Group query
func InterceptorGroup() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.GroupFunc(func(ctx context.Context, q *generated.GroupQuery) (generated.Value, error) {
			// run the query
			v, err := next.Query(ctx, q)
			if err != nil {
				return nil, err
			}

			return filterGroupsByAccess(ctx, q, v)
		})
	})
}

// filterGroupsByAccess checks fga, using ListObjects, and ensure user has view access to a group before it is returned
func filterGroupsByAccess(ctx context.Context, q *generated.GroupQuery, v ent.Value) ([]*generated.Group, error) {
	q.Logger.Debugw("intercepting list group query")

	// get userID for tuple checks
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		q.Logger.Errorw("unable to get user id from echo context")
		return nil, err
	}

	qc := ent.QueryFromContext(ctx)

	var groups []*generated.Group
	// check if query is for a single group on an exists query
	if qc.Op == "Exist" {
		groupIDs, ok := v.([]string)
		if !ok {
			q.Logger.Errorw("unexpected type for group exist query")

			return nil, ErrInternalServerError
		}

		for _, g := range groupIDs {
			groups = append(groups, &generated.Group{ID: g})
		}
	} else {
		var ok bool

		groups, ok = v.([]*generated.Group)
		if !ok {
			q.Logger.Errorw("unexpected type for group query")

			return nil, ErrInternalServerError
		}
	}

	// See all groups user has view access
	groupList, err := q.Authz.ListObjectsRequest(ctx, userID, "group", fgax.CanView)
	if err != nil {
		return nil, err
	}

	userGroups := groupList.GetObjects()

	var accessibleGroups []*generated.Group

	for _, g := range groups {
		entityType := "group"

		if !fgax.ListContains(entityType, userGroups, g.ID) {
			q.Logger.Infow("access denied to group", "relation", fgax.CanView, "group_id", g.ID, "type", entityType)

			continue
		}

		// add group to accessible group
		accessibleGroups = append(accessibleGroups, g)
	}

	// return updated groups
	return accessibleGroups, nil
}
