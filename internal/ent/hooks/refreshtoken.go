package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookRefreshToken runs on refresh token creation and sets expires fields
func HookRefreshToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.RefreshTokenFunc(func(ctx context.Context, mutation *generated.RefreshTokenMutation) (generated.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				expires, _ := mutation.ExpiresAt()
				if expires.IsZero() {
					mutation.SetExpiresAt(time.Now().Add(time.Hour * 24 * 7)) // nolint: gomnd
				}
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate|ent.OpUpdateOne)
}
