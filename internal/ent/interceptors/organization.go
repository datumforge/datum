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

// InterceptorOrganization is middleware to change the Organization query
func InterceptorOrganization() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.OrganizationFunc(func(ctx context.Context, q *generated.OrganizationQuery) (generated.Value, error) {
			// We only care these checks with authz is enabled, if this is empty skip interception checks
			if q.Authz.Ofga != nil {
				// run the query
				v, err := next.Query(ctx, q)
				if err != nil {
					return nil, err
				}

				return filterOrgsByAccess(ctx, q, v)
			}

			return next.Query(ctx, q)
		})
	})
}

// filterOrgsByAccess checks fga, using ListObjects, and ensure user has view access to an org before it is returned
// this checks both the org itself and any parent org in the request
func filterOrgsByAccess(ctx context.Context, q *generated.OrganizationQuery, v ent.Value) ([]*generated.Organization, error) {
	q.Logger.Debugw("intercepting list organization query")

	orgs, ok := v.([]*generated.Organization)
	if !ok {
		q.Logger.Errorw("unexpected type for organization query")

		return nil, ErrInternalServerError
	}

	// get userID for tuple checks
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		q.Logger.Errorw("unable to get user id from echo context")
		return nil, err
	}

	// See all orgs user has view access
	orgList, err := q.Authz.ListObjectsRequest(ctx, userID, "organization", fga.CanView)
	if err != nil {
		return nil, err
	}

	userOrgs := orgList.GetObjects()

	var accessibleOrgs []*generated.Organization

	for _, o := range orgs {
		entityType := strings.ToLower(o.Update().Mutation().Type())

		// check root org
		if !fga.ListContains(entityType, userOrgs, o.ID) {
			q.Logger.Infow("access denied to organization", "relation", fga.CanView, "organization_id", o.ID, "type", entityType)

			// go to next org, no need to check parent
			continue
		}

		// check parent org, if requested
		if o.ParentOrganizationID != "" {
			q.Logger.Debugw("checking parent organization access", "parent_organization_id", o.ParentOrganizationID)

			if !fga.ListContains(entityType, userOrgs, o.ParentOrganizationID) {
				q.Logger.Infow("access denied to parent organization", "relation", fga.CanView, "parent_organization_id", o.ParentOrganizationID)
			}
		}

		// add org to accessible orgs
		accessibleOrgs = append(accessibleOrgs, o)
	}

	// return updated orgs, if parent is denied, its removed from the result
	return accessibleOrgs, nil
}
