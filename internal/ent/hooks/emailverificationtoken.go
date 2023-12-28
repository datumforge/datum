package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookEmailVerificationToken runs on accesstoken mutations and sets expires
func HookEmailVerificationToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.EmailVerificationTokenFunc(func(ctx context.Context, mutation *generated.EmailVerificationTokenMutation) (generated.Value, error) {
			expires, _ := mutation.TTL()
			if expires.IsZero() {
				mutation.SetTTL(time.Now().Add(time.Hour * 24 * 7)) // nolint: gomnd
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
