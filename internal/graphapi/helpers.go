package graphapi

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
)

// withTransactionalMutation automatically wrap the GraphQL mutations with a database transaction.
// This allows the ent.Client to commit at the end, or rollback the transaction in case of a GraphQL error.
func withTransactionalMutation(ctx context.Context) *ent.Client {
	return ent.FromContext(ctx)
}

// injectClient adds the db client to the context to be used with transactional mutations
func injectClient(client *ent.Client) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = ent.NewContext(ctx, client)
		return next(ctx)
	}
}

// setOrganizationInAuthContext sets the organization in the auth context based on the input if it is not already set
// in most cases this is a no-op because the organization id is set in the auth middleware
// only when multiple organizations are authorized (e.g. with a PAT) is this necessary
func setOrganizationInAuthContext(ctx context.Context, inputOrgID *string) error {
	orgID, err := auth.GetOrganizationIDFromContext(ctx)
	if err == nil && orgID != "" {
		return nil
	}

	if inputOrgID == nil {
		// this would happen on a PAT authenticated request because the org id is not set
		return fmt.Errorf("unable to determine organization id")
	}

	// ensure this org is authenticated
	orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
	if err != nil {
		return err
	}

	if !orgContains(orgIDs, *inputOrgID) {
		return fmt.Errorf("organization id %s not found in the authenticated organizations", orgID)
	}

	au, err := auth.GetAuthenticatedUserContext(ctx)
	if err != nil {
		return err
	}

	au.OrganizationID = *inputOrgID

	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return err
	}

	auth.SetAuthenticatedUserContext(ec, au)

	return nil
}

func orgContains(orgIDs []string, orgID string) bool {
	for _, id := range orgIDs {
		if id == orgID {
			return true
		}
	}

	return false
}
