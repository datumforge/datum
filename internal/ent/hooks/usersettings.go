package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/pkg/rout"
)

// HookUserSetting runs on user settings mutations and validates input on update
func HookUserSetting() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserSettingFunc(func(ctx context.Context, mutation *generated.UserSettingMutation) (generated.Value, error) {
			org, ok := mutation.DefaultOrgID()
			if ok {
				// ensure user has access to the organization
				_, err := mutation.Client().OrgMembership.Query().Where(orgmembership.OrganizationID(org)).All(ctx)
				if err != nil {
					return nil, rout.InvalidField(rout.ErrOrganizationNotFound)
				}
			}

			// // delete tfa settings if tfa is disabled
			// tfaEnabled, ok := mutation.IsTfaEnabled()
			// if ok && !tfaEnabled {
			// 	userID, err := auth.GetUserIDFromContext(ctx)
			// 	if err != nil {
			// 		return nil, err
			// 	}

			// 	_, err = mutation.Client().TFASettings.Delete().Where(tfasettings.OwnerID(userID)).Exec(ctx)
			// 	if err != nil {
			// 		return nil, err
			// 	}
			// }

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
