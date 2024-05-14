package interceptors

import (
	"context"

	"entgo.io/ent"

	"github.com/99designs/gqlgen/graphql"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/pkg/auth"
)

func InterceptorUser() ent.Interceptor {
	return intercept.TraverseUser(func(ctx context.Context, q *generated.UserQuery) error {
		// bypass filter if the request is allowed, this happens when a user is
		// being created, via invite or other method by another authenticated user
		// or in tests
		if _, allow := privacy.DecisionFromContext(ctx); allow {
			return nil
		}

		// allow users to be created without filtering
		rootFieldCtx := graphql.GetRootFieldContext(ctx)
		if rootFieldCtx != nil && rootFieldCtx.Object == "createUser" {
			return nil
		}

		switch filterType(ctx) {
		// if we are looking at a user in the context of an organization or group
		// filter for just those users
		case "org":
			orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
			if err == nil {
				q.Where(user.HasOrgMembershipsWith(
					orgmembership.HasOrganizationWith(
						organization.IDIn(orgIDs...),
					),
				))

				return nil
			}
		case "user":
			// if we are looking at self
			userID, err := auth.GetUserIDFromContext(ctx)
			if err == nil {
				q.Where(user.ID(userID))

				return nil
			}
		default:
			return nil
		}

		return nil
	})
}

// filterType returns the type of filter to apply to the query
func filterType(ctx context.Context) string {
	rootFieldCtx := graphql.GetRootFieldContext(ctx)

	if rootFieldCtx != nil {
		if rootFieldCtx.Object == "updateGroup" || rootFieldCtx.Object == "createGroup" {
			return "org"
		}
	}

	qCtx := ent.QueryFromContext(ctx)

	switch qCtx.Type {
	case "OrgMembership", "GroupMembership", "Group":
		return "org"
	case "Organization":
		return "" // no filter because this is filtered at the org level
	default:
		return "user"
	}
}
