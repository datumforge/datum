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

// InterceptorUser returns an ent interceptor for user that filters users based on the context of the query
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
			// if we want to get all users, don't apply any filters
			return nil
		}

		return nil
	})
}

// filterType returns the type of filter to apply to the query
func filterType(ctx context.Context) string {
	rootFieldCtx := graphql.GetRootFieldContext(ctx)

	// the extended resolvers allow members to be adding on creation or update of a group
	// so we need to filter for the org
	if rootFieldCtx != nil {
		if rootFieldCtx.Object == "updateGroup" || rootFieldCtx.Object == "createGroup" {
			return "org"
		}
	}

	qCtx := ent.QueryFromContext(ctx)
	if qCtx == nil {
		return ""
	}

	switch qCtx.Type {
	case "OrgMembership", "GroupMembership", "Group":
		return "org"
	case "Organization":
		return "" // no filter because this is filtered at the org level
	default:
		return "user"
	}
}
