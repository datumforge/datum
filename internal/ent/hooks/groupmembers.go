package hooks

import (
	"context"

	"entgo.io/ent"

	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
)

func HookGroupMembers() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.GroupMembershipFunc(func(ctx context.Context, mutation *generated.GroupMembershipMutation) (generated.Value, error) {
			// check role, if its not set the default is member
			userID, ok := mutation.UserID()
			if !ok {
				return next.Mutate(ctx, mutation)
			}

			groupID, ok := mutation.GroupID()
			if !ok {
				return next.Mutate(ctx, mutation)
			}

			group, err := mutation.Client().Group.Get(ctx, groupID)
			if err != nil {
				// group not found, let the default validation handle it
				return next.Mutate(ctx, mutation)
			}

			// ensure user is a member of the organization
			exists, err := mutation.Client().OrgMembership.Query().
				Where(orgmembership.UserID(userID)).
				Where(orgmembership.OrganizationID(group.OwnerID)).
				Exist(ctx)
			if err != nil {
				return nil, err
			}

			if !exists {
				return nil, ErrUserNotInOrg
			}

			role, _ := mutation.Role()

			props := ph.NewProperties().
				Set("group_id", groupID).
				Set("user_id", userID).
				Set("role", role)

			mutation.Analytics.Event("group_membership", props)

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
