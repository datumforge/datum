package interceptors

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorOrganizationSetting is middleware to change the org setting query
func InterceptorOrganizationSetting() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return intercept.OrganizationSettingFunc(func(ctx context.Context, q *generated.OrganizationSettingQuery) (generated.Value, error) {
			// Organization list queries should not be filtered by organization id
			// Same with OrganizationSetting queries with the Only operation
			ctxQuery := ent.QueryFromContext(ctx)
			if ctxQuery.Type == "Organization" || ctxQuery.Op == "Only" {
				return next.Query(ctx, q)
			}

			orgID, err := auth.GetOrganizationIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			// sets the organization id on the query for the current organization
			return next.Query(ctx, q.Where(organizationsetting.OrganizationID(orgID)))
		})
	})
}
