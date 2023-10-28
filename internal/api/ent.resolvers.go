package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, where *generated.GroupWhereInput) (*generated.GroupConnection, error) {
	panic(fmt.Errorf("not implemented: Groups - groups"))
}

// GroupSettingsSlice is the resolver for the groupSettingsSlice field.
func (r *queryResolver) GroupSettingsSlice(ctx context.Context) ([]*generated.GroupSettings, error) {
	panic(fmt.Errorf("not implemented: GroupSettingsSlice - groupSettingsSlice"))
}

// Integrations is the resolver for the integrations field.
func (r *queryResolver) Integrations(ctx context.Context) ([]*generated.Integration, error) {
	return r.client.Integration.Query().AllX(ctx), nil
}

// Memberships is the resolver for the memberships field.
func (r *queryResolver) Memberships(ctx context.Context) ([]*generated.Membership, error) {
	return r.client.Membership.Query().AllX(ctx), nil
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context) ([]*generated.Organization, error) {
	return r.client.Organization.Query().AllX(ctx), nil
}

// Sessions is the resolver for the sessions field.
func (r *queryResolver) Sessions(ctx context.Context) ([]*generated.Session, error) {
	return r.client.Session.Query().AllX(ctx), nil
}

// Tenants is the resolver for the tenants field.
func (r *queryResolver) Tenants(ctx context.Context) ([]*generated.Tenant, error) {
	panic(fmt.Errorf("not implemented: Tenants - tenants"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*generated.User, error) {
	return r.client.User.Query().AllX(ctx), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
