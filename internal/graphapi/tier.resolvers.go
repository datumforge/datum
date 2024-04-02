package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateTier is the resolver for the createTier field.
func (r *mutationResolver) CreateTier(ctx context.Context, input generated.CreateTierInput) (*TierCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateTier - createTier"))
}

// UpdateTier is the resolver for the updateTier field.
func (r *mutationResolver) UpdateTier(ctx context.Context, id string, input generated.UpdateTierInput) (*TierUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateTier - updateTier"))
}

// DeleteTier is the resolver for the deleteTier field.
func (r *mutationResolver) DeleteTier(ctx context.Context, id string) (*TierDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteTier - deleteTier"))
}

// Tier is the resolver for the tier field.
func (r *queryResolver) Tier(ctx context.Context, id string) (*generated.Tier, error) {
	panic(fmt.Errorf("not implemented: Tier - tier"))
}
