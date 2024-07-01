package hooks

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"github.com/99designs/gqlgen/graphql"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/enums"
)

func HookOrgMembers() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrgMembershipFunc(func(ctx context.Context, mutation *generated.OrgMembershipMutation) (generated.Value, error) {
			// check role, if its not set the default is member
			role, _ := mutation.Role()
			if role == enums.RoleOwner {
				return next.Mutate(ctx, mutation)
			}

			orgID, exists := mutation.OrganizationID()
			if !exists || orgID == "" {
				var err error
				// get the organization based on authorized context if its not set
				orgID, err = auth.GetOrganizationIDFromContext(ctx)
				if err != nil {
					return nil, fmt.Errorf("failed to get organization id from context: %w", err)
				}

				// set organization id in mutation
				mutation.SetOrganizationID(orgID)
			}

			// get the organization
			org, err := mutation.Client().Organization.Get(ctx, orgID)
			if err != nil {
				mutation.Logger.Errorw("error getting organization", "error", err)

				return nil, err
			}

			// do not allow members to be added to personal orgs
			if org.PersonalOrg {
				return nil, ErrPersonalOrgsNoMembers
			}

			retValue, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			// check to see if the default org needs to be updated for the user
			if err := updateOrgMemberDefaultOrgOnCreate(ctx, mutation, orgID); err != nil {
				return retValue, err
			}

			if userID, ok := mutation.UserID(); ok {
				role, _ := mutation.Role()

				// allow the user to be pulled directly with a GET User, which is not allowed by default
				// the traverser will not allow this, so we need to create a new context
				allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

				user, err := mutation.Client().User.Get(allowCtx, userID)
				if err != nil {
					mutation.Logger.Errorw("error getting user", "error", err)

					return nil, err
				}

				orgName, err := auth.GetOrganizationNameFromContext(ctx)
				if err != nil {
					mutation.Logger.Errorw("error getting org name from context", "error", err)

					return nil, err
				}

				props := ph.NewProperties().
					Set("organization_name", orgName).
					Set("user_name", user.FirstName+user.LastName).
					Set("join_role", role.String())

				mutation.Analytics.Event("org_membership", props)
			}

			return retValue, err
		})
	}, ent.OpCreate)
}

// HookOrgMembersDelete is a hook that runs during the delete operation of an org membership
func HookOrgMembersDelete() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrgMembershipFunc(func(ctx context.Context, mutation *generated.OrgMembershipMutation) (generated.Value, error) {
			// we only want to do this on direct deleteOrgMembership operations
			// deleteOrganization will be handled by the organization hook
			rootFieldCtx := graphql.GetRootFieldContext(ctx)
			if rootFieldCtx == nil || rootFieldCtx.Object != "deleteOrgMembership" {
				return next.Mutate(ctx, mutation)
			}

			// get the existing org membership
			id, ok := mutation.ID()
			if !ok {
				return nil, fmt.Errorf("%w: %s", ErrInvalidInput, "id is required")
			}

			// get the org membership
			orgMembership, err := mutation.Client().OrgMembership.Get(ctx, id)
			if err != nil {
				return nil, err
			}

			retValue, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			allowCtx := privacy.DecisionContext(ctx, privacy.Allow)
			if _, err = checkAndUpdateDefaultOrg(allowCtx, orgMembership.UserID, orgMembership.OrganizationID); err != nil {
				return nil, err
			}

			return retValue, err
		})
	}, ent.OpDeleteOne|ent.OpDelete|ent.OpUpdate|ent.OpUpdateOne)
}

// updateOrgMemberDefaultOrgOnCreate updates the user's default org if the user has no default org or
// the default org is their personal org
func updateOrgMemberDefaultOrgOnCreate(ctx context.Context, mutation *generated.OrgMembershipMutation, orgID string) error {
	// allow access to the user settings; this is required to update the default org
	allowCtx := privacy.DecisionContext(ctx, privacy.Allow)

	// get the user id from the mutation, this is a required field
	userID, ok := mutation.UserID()
	if !ok {
		// this should never happen because the mutation should have already failed
		return fmt.Errorf("%w: %s", ErrInvalidInput, "user id is required")
	}

	return updateDefaultOrgIfPersonal(allowCtx, userID, orgID, mutation.Client())
}
