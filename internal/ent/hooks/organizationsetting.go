package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
)

// HookOrganizationSetting runs on organization settings mutations and validates input on update
func HookOrganizationSetting() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.OrganizationSettingFunc(func(ctx context.Context, mutation *generated.OrganizationSettingMutation) (generated.Value, error) {

			//			if email, ok := mutation.BillingEmail(); ok {
			//				url := gravatar.New(email, nil)
			//				mutation.SetAvatarRemoteURL(url)
			//
			//				// use the email as the display name, if not provided on creation
			//				if mutation.Op().Is(ent.OpCreate) {
			//					displayName, _ := mutation.DisplayName()
			//					if displayName == "" {
			//						mutation.SetDisplayName(email)
			//					}
			//				}
			//			}

			userID, err := auth.GetUserIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			orgID, ok := mutation.OrganizationID()
			if !ok {
				return nil, rout.InvalidField(rout.ErrOrganizationNotFound)
			}

			exists, err := mutation.Client().OrgMembership.Query().
				Where(orgmembership.UserID(userID)).
				Where(orgmembership.OrganizationID(orgID)).
				Exist(ctx)
			if err != nil {
				return nil, err
			}

			if !exists {
				return nil, ErrUserNotInOrg
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
