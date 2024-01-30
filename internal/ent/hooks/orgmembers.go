package hooks

import (
	"context"

	"entgo.io/ent"

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

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
