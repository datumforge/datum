package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
)

// HookCreateAPIToken runs on api token mutations and sets expires and owner id
func HookCreateAPIToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.PersonalAccessTokenFunc(func(ctx context.Context, mutation *generated.PersonalAccessTokenMutation) (generated.Value, error) {
			// default the expiration to 7 days
			expires, ok := mutation.ExpiresAt()
			if !ok || expires.IsZero() {
				mutation.SetExpiresAt(time.Now().Add(time.Hour * 24 * 7)) // nolint: gomnd
			}

			orgID, err := auth.GetOrganizationIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			// set organization on the token
			mutation.SetOwnerID(orgID)

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

// HookUpdateAPIToken runs on api token update and redacts the token
func HookUpdateAPIToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.PersonalAccessTokenFunc(func(ctx context.Context, mutation *generated.PersonalAccessTokenMutation) (generated.Value, error) {
			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			// redact the token
			pat, ok := retVal.(*generated.PersonalAccessToken)
			if !ok {
				return retVal, nil
			}

			pat.Token = redacted

			return pat, nil
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
