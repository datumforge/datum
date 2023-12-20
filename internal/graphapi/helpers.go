package graphapi

import (
	"context"

	"github.com/99designs/gqlgen/graphql"

	ent "github.com/datumforge/datum/internal/ent/generated"
)

// WithTransactionalMutation automatically wrap the GraphQL mutations with a database transaction.
// This allows the ent.Client to commit at the end, or rollback the transaction in case of a GraphQL error.
func WithTransactionalMutation(ctx context.Context) *ent.Client {
	return ent.FromContext(ctx)
}

func injectClient(client *ent.Client) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = ent.NewContext(ctx, client)
		return next(ctx)
	}
}
