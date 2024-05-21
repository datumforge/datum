package interceptors

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
)

// InterceptorOrganization is middleware to change the Organization query
func InterceptorOrganization() ent.Interceptor {
	return intercept.TraverseOrganization(func(ctx context.Context, q *generated.OrganizationQuery) error {
		// by pass checks on invite or pre-allowed request
		if _, allow := privacy.DecisionFromContext(ctx); allow {
			return nil
		}

		if rule.ContextHasPrivacyTokenOfType(ctx, &token.OrgInviteToken{}) ||
			rule.ContextHasPrivacyTokenOfType(ctx, &token.SignUpToken{}) {
			return nil
		}

		// if this is an API token, only allow the query if it is for the organization
		if auth.IsAPITokenAuthentication(ctx) {
			orgID, err := auth.GetOrganizationIDFromContext(ctx)
			if err != nil {
				return err
			}

			q.Where(organization.IDEQ(orgID))

			return nil
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		// add filter to query
		q.Where(organization.HasMembersWith(orgmembership.HasUserWith(user.ID(userID))))

		return nil
	})
}
