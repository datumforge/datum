package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/utils/gravatar"
)

// HookUser runs on user mutations validate and hash the password and set default values that are not provided
func HookUser() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, mutation *generated.UserMutation) (generated.Value, error) {
			if password, ok := mutation.Password(); ok {
				// validate password before its encrypted
				if passwd.Strength(password) < passwd.Moderate {
					return nil, auth.ErrPasswordTooWeak
				}

				hash, err := passwd.CreateDerivedKey(password)
				if err != nil {
					return nil, err
				}

				mutation.SetPassword(hash)
			}

			if email, ok := mutation.Email(); ok {
				url := gravatar.New(email, nil)
				mutation.SetAvatarRemoteURL(url)

				// use the email as the display name, if not provided
				if displayName, ok := mutation.DisplayName(); ok {
					if displayName == "" {
						mutation.SetDisplayName(email)
					}
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}
