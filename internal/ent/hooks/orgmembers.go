package hooks

import (
	"context"

	"entgo.io/ent"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/analytics"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

func HookOrgMembers() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrgMembershipFunc(func(ctx context.Context, mutation *generated.OrgMembershipMutation) (generated.Value, error) {
			// check role, if its not set the default is member
			role, _ := mutation.Role()
			if role == enums.RoleOwner {
				return next.Mutate(ctx, mutation)
			}

			// get the organization based on input
			orgID, exists := mutation.OrganizationID()
			if exists {
				org, err := mutation.Client().Organization.Get(ctx, orgID)
				if err != nil {
					mutation.Logger.Errorw("error getting organization", "error", err)

					return nil, err
				}

				// do not allow members to be added to personal orgs
				if org.PersonalOrg {
					return nil, ErrPersonalOrgsNoMembers
				}
			}

			retValue, err := next.Mutate(ctx, mutation)

			if userID, ok := mutation.UserID(); ok {
				role, _ := mutation.Role()
				org, err := mutation.Client().Organization.Get(ctx, orgID)
				if err != nil {
					mutation.Logger.Errorw("error getting organization", "error", err)

					return nil, err
				}

				user, err := mutation.Client().User.Get(ctx, userID)
				if err != nil {
					mutation.Logger.Errorw("error getting user", "error", err)

					return nil, err
				}

				props := ph.NewProperties().
					Set("organization_name", org.Name).
					Set("user_name", user.FirstName+user.LastName).
					Set("join_role", role.String())

				analytics.OrganizationEvent(orgID, userID, "organization_membership", props)
				analytics.UserEvent(userID, "organization_membership", props)
			}

			return retValue, err
		})
	}, ent.OpCreate)
}
