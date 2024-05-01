package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
)

// HookCreateAPIToken runs on api token mutations and sets expires and owner id
func HookCreateAPIToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.APITokenFunc(func(ctx context.Context, mutation *generated.APITokenMutation) (generated.Value, error) {
			orgID, err := auth.GetOrganizationIDFromContext(ctx)
			if err != nil {
				return nil, err
			}

			// set organization on the token
			mutation.SetOwnerID(orgID)

			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			token, ok := retVal.(*generated.APIToken)
			if !ok {
				return retVal, err
			}

			// create the relationship tuples in fga for the token
			tuples := []fgax.TupleKey{}

			// TODO (sfunk): this shouldn't be a static list
			for _, scope := range token.Scopes {
				var relation string

				switch scope {
				case "read":
					relation = "can_view"
				case "write":
					relation = "can_edit"
				case "delete":
					relation = "can_delete"
				}

				apiKeyTuple, err := getTupleKey(token.ID, "service", orgID, "organization", relation)
				if err != nil {
					return nil, err
				}

				tuples = append(tuples, apiKeyTuple)
			}

			if _, err := mutation.Authz.WriteTupleKeys(ctx, tuples, nil); err != nil {
				mutation.Logger.Errorw("failed to create relationship tuple", "error", err)

				return nil, err
			}

			return retVal, err
		})
	}, ent.OpCreate)
}

// HookUpdateAPIToken runs on api token update and redacts the token
func HookUpdateAPIToken() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.APITokenFunc(func(ctx context.Context, mutation *generated.APITokenMutation) (generated.Value, error) {
			retVal, err := next.Mutate(ctx, mutation)
			if err != nil {
				return nil, err
			}

			// redact the token
			at, ok := retVal.(*generated.APIToken)
			if !ok {
				return retVal, nil
			}

			at.Token = redacted

			return at, nil
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}
