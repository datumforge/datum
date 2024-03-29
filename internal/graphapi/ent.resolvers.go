package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// AuthStyle is the resolver for the authStyle field.
func (r *oauthProviderResolver) AuthStyle(ctx context.Context, obj *generated.OauthProvider) (int, error) {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]generated.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Entitlements is the resolver for the entitlements field.
func (r *queryResolver) Entitlements(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.EntitlementWhereInput) (*generated.EntitlementConnection, error) {
	return withTransactionalMutation(ctx).Entitlement.Query().Paginate(ctx, after, first, before, last, generated.WithEntitlementFilter(where.Filter))
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.GroupOrder, where *generated.GroupWhereInput) (*generated.GroupConnection, error) {
	return withTransactionalMutation(ctx).Group.Query().Paginate(ctx, after, first, before, last, generated.WithGroupOrder(orderBy), generated.WithGroupFilter(where.Filter))
}

// GroupMemberships is the resolver for the groupMemberships field.
func (r *queryResolver) GroupMemberships(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.GroupMembershipWhereInput) (*generated.GroupMembershipConnection, error) {
	return withTransactionalMutation(ctx).GroupMembership.Query().Paginate(ctx, after, first, before, last, generated.WithGroupMembershipFilter(where.Filter))
}

// GroupSettings is the resolver for the groupSettings field.
func (r *queryResolver) GroupSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.GroupSettingWhereInput) (*generated.GroupSettingConnection, error) {
	return withTransactionalMutation(ctx).GroupSetting.Query().Paginate(ctx, after, first, before, last, generated.WithGroupSettingFilter(where.Filter))
}

// Integrations is the resolver for the integrations field.
func (r *queryResolver) Integrations(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.IntegrationOrder, where *generated.IntegrationWhereInput) (*generated.IntegrationConnection, error) {
	return withTransactionalMutation(ctx).Integration.Query().Paginate(ctx, after, first, before, last, generated.WithIntegrationOrder(orderBy), generated.WithIntegrationFilter(where.Filter))
}

// Invites is the resolver for the invites field.
func (r *queryResolver) Invites(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.InviteWhereInput) (*generated.InviteConnection, error) {
	return r.client.Invite.Query().Paginate(ctx, after, first, before, last, generated.WithInviteFilter(where.Filter))
}

// OauthProviders is the resolver for the oauthProviders field.
func (r *queryResolver) OauthProviders(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OauthProviderWhereInput) (*generated.OauthProviderConnection, error) {
	panic(fmt.Errorf("not implemented: OauthProviders - oauthProviders"))
}

// OhAuthTooTokens is the resolver for the ohAuthTooTokens field.
func (r *queryResolver) OhAuthTooTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OhAuthTooTokenWhereInput) (*generated.OhAuthTooTokenConnection, error) {
	panic(fmt.Errorf("not implemented: OhAuthTooTokens - ohAuthTooTokens"))
}

// OrgMemberships is the resolver for the orgMemberships field.
func (r *queryResolver) OrgMemberships(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OrgMembershipWhereInput) (*generated.OrgMembershipConnection, error) {
	return withTransactionalMutation(ctx).OrgMembership.Query().Paginate(ctx, after, first, before, last, generated.WithOrgMembershipFilter(where.Filter))
}

// Organizations is the resolver for the organizations field.
func (r *queryResolver) Organizations(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.OrganizationOrder, where *generated.OrganizationWhereInput) (*generated.OrganizationConnection, error) {
	return withTransactionalMutation(ctx).Organization.Query().Paginate(ctx, after, first, before, last, generated.WithOrganizationOrder(orderBy), generated.WithOrganizationFilter(where.Filter))
}

// OrganizationSettings is the resolver for the organizationSettings field.
func (r *queryResolver) OrganizationSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.OrganizationSettingWhereInput) (*generated.OrganizationSettingConnection, error) {
	return withTransactionalMutation(ctx).OrganizationSetting.Query().Paginate(ctx, after, first, before, last, generated.WithOrganizationSettingFilter(where.Filter))
}

// PersonalAccessTokens is the resolver for the personalAccessTokens field.
func (r *queryResolver) PersonalAccessTokens(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.PersonalAccessTokenWhereInput) (*generated.PersonalAccessTokenConnection, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	return withTransactionalMutation(ctx).PersonalAccessToken.Query().Paginate(ctx, after, first, before, last, generated.WithPersonalAccessTokenFilter(where.Filter))
}

// Subscribers is the resolver for the subscribers field.
func (r *queryResolver) Subscribers(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.SubscriberWhereInput) (*generated.SubscriberConnection, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	return withTransactionalMutation(ctx).Subscriber.Query().Paginate(ctx, after, first, before, last, generated.WithSubscriberFilter(where.Filter))
}

// TfaSettingsSlice is the resolver for the tfaSettingsSlice field.
func (r *queryResolver) TfaSettingsSlice(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.TFASettingsWhereInput) (*generated.TFASettingsConnection, error) {
	panic(fmt.Errorf("not implemented: TfaSettingsSlice - tfaSettingsSlice"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, orderBy *generated.UserOrder, where *generated.UserWhereInput) (*generated.UserConnection, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	return withTransactionalMutation(ctx).User.Query().Paginate(ctx, after, first, before, last, generated.WithUserOrder(orderBy), generated.WithUserFilter(where.Filter))
}

// UserSettings is the resolver for the userSettings field.
func (r *queryResolver) UserSettings(ctx context.Context, after *entgql.Cursor[string], first *int, before *entgql.Cursor[string], last *int, where *generated.UserSettingWhereInput) (*generated.UserSettingConnection, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	return withTransactionalMutation(ctx).UserSetting.Query().Paginate(ctx, after, first, before, last, generated.WithUserSettingFilter(where.Filter))
}

// AuthStyle is the resolver for the authStyle field.
func (r *createOauthProviderInputResolver) AuthStyle(ctx context.Context, obj *generated.CreateOauthProviderInput, data int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// AuthStyle is the resolver for the authStyle field.
func (r *oauthProviderWhereInputResolver) AuthStyle(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// AuthStyleNeq is the resolver for the authStyleNEQ field.
func (r *oauthProviderWhereInputResolver) AuthStyleNeq(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleNeq - authStyleNEQ"))
}

// AuthStyleIn is the resolver for the authStyleIn field.
func (r *oauthProviderWhereInputResolver) AuthStyleIn(ctx context.Context, obj *generated.OauthProviderWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: AuthStyleIn - authStyleIn"))
}

// AuthStyleNotIn is the resolver for the authStyleNotIn field.
func (r *oauthProviderWhereInputResolver) AuthStyleNotIn(ctx context.Context, obj *generated.OauthProviderWhereInput, data []int) error {
	panic(fmt.Errorf("not implemented: AuthStyleNotIn - authStyleNotIn"))
}

// AuthStyleGt is the resolver for the authStyleGT field.
func (r *oauthProviderWhereInputResolver) AuthStyleGt(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleGt - authStyleGT"))
}

// AuthStyleGte is the resolver for the authStyleGTE field.
func (r *oauthProviderWhereInputResolver) AuthStyleGte(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleGte - authStyleGTE"))
}

// AuthStyleLt is the resolver for the authStyleLT field.
func (r *oauthProviderWhereInputResolver) AuthStyleLt(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleLt - authStyleLT"))
}

// AuthStyleLte is the resolver for the authStyleLTE field.
func (r *oauthProviderWhereInputResolver) AuthStyleLte(ctx context.Context, obj *generated.OauthProviderWhereInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyleLte - authStyleLTE"))
}

// AuthStyle is the resolver for the authStyle field.
func (r *updateOauthProviderInputResolver) AuthStyle(ctx context.Context, obj *generated.UpdateOauthProviderInput, data *int) error {
	panic(fmt.Errorf("not implemented: AuthStyle - authStyle"))
}

// OauthProvider returns OauthProviderResolver implementation.
func (r *Resolver) OauthProvider() OauthProviderResolver { return &oauthProviderResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// CreateGroupInput returns CreateGroupInputResolver implementation.
func (r *Resolver) CreateGroupInput() CreateGroupInputResolver { return &createGroupInputResolver{r} }

// CreateOauthProviderInput returns CreateOauthProviderInputResolver implementation.
func (r *Resolver) CreateOauthProviderInput() CreateOauthProviderInputResolver {
	return &createOauthProviderInputResolver{r}
}

// CreateOrganizationInput returns CreateOrganizationInputResolver implementation.
func (r *Resolver) CreateOrganizationInput() CreateOrganizationInputResolver {
	return &createOrganizationInputResolver{r}
}

// OauthProviderWhereInput returns OauthProviderWhereInputResolver implementation.
func (r *Resolver) OauthProviderWhereInput() OauthProviderWhereInputResolver {
	return &oauthProviderWhereInputResolver{r}
}

// UpdateGroupInput returns UpdateGroupInputResolver implementation.
func (r *Resolver) UpdateGroupInput() UpdateGroupInputResolver { return &updateGroupInputResolver{r} }

// UpdateOauthProviderInput returns UpdateOauthProviderInputResolver implementation.
func (r *Resolver) UpdateOauthProviderInput() UpdateOauthProviderInputResolver {
	return &updateOauthProviderInputResolver{r}
}

// UpdateOrganizationInput returns UpdateOrganizationInputResolver implementation.
func (r *Resolver) UpdateOrganizationInput() UpdateOrganizationInputResolver {
	return &updateOrganizationInputResolver{r}
}

// UpdateTFASettingsInput returns UpdateTFASettingsInputResolver implementation.
func (r *Resolver) UpdateTFASettingsInput() UpdateTFASettingsInputResolver {
	return &updateTFASettingsInputResolver{r}
}

type oauthProviderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createGroupInputResolver struct{ *Resolver }
type createOauthProviderInputResolver struct{ *Resolver }
type createOrganizationInputResolver struct{ *Resolver }
type oauthProviderWhereInputResolver struct{ *Resolver }
type updateGroupInputResolver struct{ *Resolver }
type updateOauthProviderInputResolver struct{ *Resolver }
type updateOrganizationInputResolver struct{ *Resolver }
type updateTFASettingsInputResolver struct{ *Resolver }
