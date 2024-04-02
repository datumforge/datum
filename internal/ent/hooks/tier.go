package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/tier"
)

// HookTier runs on tier mutations and ensures uniqueness of tier names for root tiers
func HookTier() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.TierFunc(func(ctx context.Context, mutation *generated.TierMutation) (generated.Value, error) {
			// ensure uniqueness for null orgs, ent doesn't support IFNULL in unique constraints
			_, ok := mutation.OwnerID()
			name, _ := mutation.Name()

			if !ok {
				exist, err := mutation.Client().Tier.Query().
					Where(
						tier.Name(name),
						tier.OwnerIDIsNil(),
						tier.DeletedAtIsNil(),
					).
					Exist(ctx)

				// if error, return
				if err != nil {
					return nil, err
				}

				// if the tier already is a tier to datum, say so
				if exist {
					return nil, ErrTierAlreadyExists
				}
			}

			// continue on
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
