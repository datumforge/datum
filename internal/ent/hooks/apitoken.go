package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/pkg/auth"
	sliceutil "github.com/datumforge/datum/pkg/utils/slice"
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
			tuples, err := createScopeTuples(token.Scopes, orgID, token.ID)
			if err != nil {
				return retVal, err
			}

			// create the relationship tuples if we have any
			if len(tuples) > 0 {
				if _, err := mutation.Authz.WriteTupleKeys(ctx, tuples, nil); err != nil {
					mutation.Logger.Errorw("failed to create relationship tuple", "error", err)

					return nil, err
				}
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

			// create the relationship tuples in fga for the token
			newScopes, err := getNewScopes(ctx, mutation)
			if err != nil {
				return at, err
			}

			tuples, err := createScopeTuples(newScopes, at.OwnerID, at.ID)
			if err != nil {
				return retVal, err
			}

			// create the relationship tuples if we have any
			if len(tuples) > 0 {
				if _, err := mutation.Authz.WriteTupleKeys(ctx, tuples, nil); err != nil {
					mutation.Logger.Errorw("failed to create relationship tuple", "error", err)

					return nil, err
				}
			}

			return at, nil
		})
	}, ent.OpUpdate|ent.OpUpdateOne)
}

// createScopeTuples creates the relationship tuples for the token
func createScopeTuples(scopes []string, orgID, tokenID string) (tuples []fgax.TupleKey, err error) {
	// create the relationship tuples in fga for the token
	// TODO (sfunk): this shouldn't be a static list
	for _, scope := range scopes {
		var relation string

		switch scope {
		case "read":
			relation = "can_view"
		case "write":
			relation = "can_edit"
		case "delete":
			relation = "can_delete"
		}

		var apiKeyTuple fgax.TupleKey

		apiKeyTuple, err = getTupleKey(tokenID, "service", orgID, "organization", relation)
		if err != nil {
			return
		}

		tuples = append(tuples, apiKeyTuple)
	}

	return
}

// getNewScopes returns the new scopes that were added to the token during an update
// NOTE: there is an AppendedScopes on the mutation, but this is not populated
// so calculating the new scopes for now
func getNewScopes(ctx context.Context, mutation *generated.APITokenMutation) ([]string, error) {
	scopes, ok := mutation.Scopes()
	if !ok {
		return nil, nil
	}

	oldScopes, err := mutation.OldScopes(ctx)
	if err != nil {
		return nil, err
	}

	var newScopes []string

	for _, scope := range scopes {
		if !sliceutil.Contains(oldScopes, scope) {
			newScopes = append(newScopes, scope)
		}
	}

	return newScopes, nil
}
