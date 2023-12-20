package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookSession runs on refresh token creation and sets expires fields
func HookSession() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.SessionFunc(func(ctx context.Context, mutation *generated.SessionMutation) (generated.Value, error) {
			expires, _ := mutation.ExpiresAt()
			if expires.IsZero() {
				mutation.SetExpiresAt(time.Now().Add(time.Hour * 24 * 7)) // nolint: gomnd
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
