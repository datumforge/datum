package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookPasswordResetToken runs on reset token mutations and sets expires
func HookPasswordResetToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.PasswordResetTokenFunc(func(ctx context.Context, mutation *generated.PasswordResetTokenMutation) (generated.Value, error) {
			expires, _ := mutation.TTL()
			if expires.IsZero() {
				mutation.SetTTL(time.Now().Add(time.Minute * 15)) // nolint: gomnd
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
