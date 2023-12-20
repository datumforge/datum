package graphapi

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	ent "github.com/datumforge/datum/internal/ent/generated"
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
