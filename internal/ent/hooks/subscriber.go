package hooks

import (
	"context"

	"entgo.io/ent"

	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
)

// HookSubscriber runs on subscriber mutations and ensures uniqueness of email for root subscribers
func HookSubscriber() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.SubscriberFunc(func(ctx context.Context, mutation *generated.SubscriberMutation) (generated.Value, error) {
			// ensure uniqueness for null orgs, ent doesn't support IFNULL in unique constraints
			owner, ok := mutation.OwnerID()
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

			props := ph.NewProperties().
				Set("subscriber_email", email).
				Set("organization_name", owner)

			mutation.Analytics.Event("subscriber_created", props)

			// continue on
			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}
