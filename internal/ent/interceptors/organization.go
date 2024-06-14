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

		orgIDs, err := getAllActorOrgIDs(ctx)
		if err != nil {
			return err
		}

		q.Where(organization.IDIn(orgIDs...))

		return nil
	})
}

// getAllActorOrgIDs returns all the organization IDs that the user is a member of, either directly or indirectly via a parent organization
func getAllActorOrgIDs(ctx context.Context) ([]string, error) {
	// allow the request, otherwise we would be in an infinite loop, as this function is called by the interceptor
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	directOrgIDs, err := generated.FromContext(allowCtx).Organization.
		Query().
		Where(organization.HasMembersWith(orgmembership.HasUserWith(user.ID(userID)))).
		Select(organization.FieldID).
		Strings(allowCtx)
	if err != nil {
		return nil, err
	}

	return getAllRelatedOrgs(allowCtx, directOrgIDs)
}

// getAllRelatedOrgs returns all the combined directly related orgs in addition to any child organizations of the parent organizations
func getAllRelatedOrgs(ctx context.Context, directOrgIDs []string) ([]string, error) {
	allOrgsIDs := directOrgIDs

	for _, id := range directOrgIDs {
		co, err := getChildOrgIDs(ctx, id)
		if err != nil {
			return nil, err
		}

		allOrgsIDs = append(allOrgsIDs, co...)
	}

	return allOrgsIDs, nil
}

// getChildOrgIDs returns all the child organizations of the parent organization
func getChildOrgIDs(ctx context.Context, parentOrgID string) ([]string, error) {
	// given an organization id, get all the children of that organization
	childOrgs, err := generated.FromContext(ctx).Organization.
		Query().
		Where(
			organization.HasParentWith(organization.ID(parentOrgID)),
		).
		Select(organization.FieldID).
		Strings(ctx)
	if err != nil {
		return nil, err
	}

	allOrgsIDs := childOrgs

	for _, orgID := range childOrgs {
		// recursively get all the children of the children
		coIDs, err := getChildOrgIDs(ctx, orgID)
		if err != nil {
			return nil, err
		}

		allOrgsIDs = append(allOrgsIDs, coIDs...)
	}

	return allOrgsIDs, nil
}
