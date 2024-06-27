package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
)

func HookFeature() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.FeatureFunc(func(ctx context.Context, mutation *generated.FeatureMutation) (generated.Value, error) {
			// set the display name if it is not set
			if mutation.Op() == ent.OpCreate {
				displayName, _ := mutation.DisplayName()
				if displayName == "" {
					name, ok := mutation.Name()
					if ok {
						mutation.SetDisplayName(name)
					}
				}
			}

			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			return retVal, err
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}
