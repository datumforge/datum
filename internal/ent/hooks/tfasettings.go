package hooks

import (
	"context"
	"database/sql"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/utils/totp"
)

func HookEnableTFA() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.TFASettingsFunc(func(ctx context.Context, mutation *generated.TFASettingsMutation) (generated.Value, error) {
			u, err := constructTOTPUser(ctx, mutation)
			if err != nil {
				return nil, err
			}

			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			// update user settings
			_, err = mutation.Client().UserSetting.Update().
				Where(usersetting.UserID(u.ID)).
				SetIsTfaEnabled(true). // set tfa enabled to true
				Save(ctx)

			return retVal, err
		})
	}, ent.OpCreate)
}

func HookTFAUpdate() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.TFASettingsFunc(func(ctx context.Context, mutation *generated.TFASettingsMutation) (generated.Value, error) {
			// once verified, create recovery codes
			verified, ok := mutation.Verified()

			// if recovery codes are cleared, generate new ones
			cleared := mutation.RecoveryCodesCleared()

			if (ok && verified) || cleared {
				u, err := constructTOTPUser(ctx, mutation)
				if err != nil {
					return nil, err
				}

				u.TFASecret, err = mutation.TOTP.TOTPManager.TOTPSecret(u)
				if err != nil {
					return nil, err
				}

				codes := mutation.TOTP.TOTPManager.GenerateRecoveryCodes()
				mutation.SetRecoveryCodes(codes)
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}

func constructTOTPUser(ctx context.Context, mutation *generated.TFASettingsMutation) (*totp.User, error) {
	userID, ok := mutation.OwnerID()
	if !ok {
		var err error

		userID, err = auth.GetUserIDFromContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	u := &totp.User{
		ID: userID,
	}

	// get the user object
	user, err := mutation.Client().User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	// get the full setting object
	setting, err := mutation.Client().UserSetting.Query().Where(usersetting.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// set the TFA settings
	u.IsEmailOTPAllowed, _ = mutation.EmailOtpAllowed()
	u.IsPhoneOTPAllowed, _ = mutation.PhoneOtpAllowed()
	u.IsTOTPAllowed, _ = mutation.TotpAllowed()

	// setup account name fields
	u.Email = sql.NullString{
		String: user.Email,
	}

	phoneNumber := setting.PhoneNumber
	if phoneNumber != nil {
		u.Phone = sql.NullString{
			String: *setting.PhoneNumber,
		}
	}

	return u, nil
}
