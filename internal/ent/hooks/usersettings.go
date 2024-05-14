package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/tfasetting"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
)

// HookUserSetting runs on user settings mutations and validates input on update
func HookUserSetting() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserSettingFunc(func(ctx context.Context, mutation *generated.UserSettingMutation) (generated.Value, error) {
			org, ok := mutation.DefaultOrgID()
			if ok {
				// ensure user has access to the organization
				// the ID is always set on update
				userSettingID, _ := mutation.ID()

				owner, err := mutation.Client().User.Query().Where(user.HasSettingWith(usersetting.ID(userSettingID))).Only(ctx)
				if err != nil {
					return nil, err
				}

				allow, err := mutation.Authz.CheckOrgAccess(ctx, owner.ID, auth.UserSubjectType, org, fgax.CanView)
				if err != nil || !allow {
					return nil, rout.InvalidField(rout.ErrOrganizationNotFound)
				}
			}

			// delete tfa setting if tfa is disabled
			tfaEnabled, ok := mutation.IsTfaEnabled()
			if ok && !tfaEnabled {
				userID, err := auth.GetUserIDFromContext(ctx)
				if err != nil {
					return nil, err
				}

				_, err = mutation.Client().TFASetting.Delete().Where(tfasetting.OwnerID(userID)).Exec(ctx)
				if err != nil {
					return nil, err
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
