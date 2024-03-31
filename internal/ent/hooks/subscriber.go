package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
)

// HookSubscriber runs on subscriber mutations and ensures uniqueness of email for root subscribers
func HookSubscriber() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.SubscriberFunc(func(ctx context.Context, mutation *generated.SubscriberMutation) (generated.Value, error) {
			// ensure uniqueness for null orgs, ent doesn't support IFNULL in unique constraints
			_, ok := mutation.OwnerID()
			email, _ := mutation.Email()

			if !ok {
				exist, err := mutation.Client().Subscriber.Query().
					Where(
						subscriber.EmailEQ(email),
						subscriber.OwnerIDIsNil(),
						subscriber.DeletedAtIsNil(),
					).
					Exist(ctx)

				// if error, return
				if err != nil {
					return nil, err
				}

				// if the user already is a subscriber to datum, say so
				if exist {
					return nil, ErrUserAlreadySubscriber
				}
			}

			// continue on
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
