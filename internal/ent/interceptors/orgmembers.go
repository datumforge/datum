package interceptors

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorOrgMembers is middleware to change the Org Members query
func InterceptorOrgMembers() ent.Interceptor {
	return intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
		// Organization list queries should not be filtered by organization id
		ctxQuery := ent.QueryFromContext(ctx)
		if ctxQuery.Type == "Organization" {
			return nil
		}

		orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
		if err != nil {
			return err
		}

		// sets the organization id on the query for the current organization
		q.WhereP(orgmembership.OrganizationIDIn(orgIDs...))

		return nil
	})
}
