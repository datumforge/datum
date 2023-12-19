package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

// HookPersonalAccessToken runs on accesstoken mutations and sets expires
func HookPersonalAccessToken() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.PersonalAccessTokenFunc(func(ctx context.Context, m *generated.PersonalAccessTokenMutation) (ent.Value, error) {
			if m.Op().Is(ent.OpCreate) {
				m.SetExpiresAt(time.Now().Add(time.Hour * 24 * 7)) // nolint: gomnd
			}
			return next.Mutate(ctx, m)
		})
	}
}
