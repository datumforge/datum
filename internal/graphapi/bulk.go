package graphapi

// THIS CODE IS REGENERATED BY github.com/datumforge/datum/pkg/gqlplugin. DO NOT EDIT.

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// bulkCreateAPIToken uses the CreateBulk function to create multiple APIToken entities
func (r *mutationResolver) bulkCreateAPIToken(ctx context.Context, input []*generated.CreateAPITokenInput) (*APITokenBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.APITokenCreate, len(input))
	for i, data := range input {
		builders[i] = c.APIToken.Create().SetInput(*data)
	}

	res, err := c.APIToken.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "apitoken"}, r.logger)
	}

	// return response
	return &APITokenBulkCreatePayload{
		APITokens: res,
	}, nil
}

// bulkCreateDocumentData uses the CreateBulk function to create multiple DocumentData entities
func (r *mutationResolver) bulkCreateDocumentData(ctx context.Context, input []*generated.CreateDocumentDataInput) (*DocumentDataBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.DocumentDataCreate, len(input))
	for i, data := range input {
		builders[i] = c.DocumentData.Create().SetInput(*data)
	}

	res, err := c.DocumentData.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "documentdata"}, r.logger)
	}

	// return response
	return &DocumentDataBulkCreatePayload{
		DocumentData: res,
	}, nil
}

// bulkCreateEntitlement uses the CreateBulk function to create multiple Entitlement entities
func (r *mutationResolver) bulkCreateEntitlement(ctx context.Context, input []*generated.CreateEntitlementInput) (*EntitlementBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.EntitlementCreate, len(input))
	for i, data := range input {
		builders[i] = c.Entitlement.Create().SetInput(*data)
	}

	res, err := c.Entitlement.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "entitlement"}, r.logger)
	}

	// return response
	return &EntitlementBulkCreatePayload{
		Entitlements: res,
	}, nil
}

// bulkCreateEvent uses the CreateBulk function to create multiple Event entities
func (r *mutationResolver) bulkCreateEvent(ctx context.Context, input []*generated.CreateEventInput) (*EventBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.EventCreate, len(input))
	for i, data := range input {
		builders[i] = c.Event.Create().SetInput(*data)
	}

	res, err := c.Event.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "event"}, r.logger)
	}

	// return response
	return &EventBulkCreatePayload{
		Events: res,
	}, nil
}

// bulkCreateFeature uses the CreateBulk function to create multiple Feature entities
func (r *mutationResolver) bulkCreateFeature(ctx context.Context, input []*generated.CreateFeatureInput) (*FeatureBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.FeatureCreate, len(input))
	for i, data := range input {
		builders[i] = c.Feature.Create().SetInput(*data)
	}

	res, err := c.Feature.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "feature"}, r.logger)
	}

	// return response
	return &FeatureBulkCreatePayload{
		Features: res,
	}, nil
}

// bulkCreateFile uses the CreateBulk function to create multiple File entities
func (r *mutationResolver) bulkCreateFile(ctx context.Context, input []*generated.CreateFileInput) (*FileBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.FileCreate, len(input))
	for i, data := range input {
		builders[i] = c.File.Create().SetInput(*data)
	}

	res, err := c.File.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "file"}, r.logger)
	}

	// return response
	return &FileBulkCreatePayload{
		Files: res,
	}, nil
}

// bulkCreateGroup uses the CreateBulk function to create multiple Group entities
func (r *mutationResolver) bulkCreateGroup(ctx context.Context, input []*generated.CreateGroupInput) (*GroupBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.GroupCreate, len(input))
	for i, data := range input {
		builders[i] = c.Group.Create().SetInput(*data)
	}

	res, err := c.Group.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "group"}, r.logger)
	}

	// return response
	return &GroupBulkCreatePayload{
		Groups: res,
	}, nil
}

// bulkCreateGroupMembership uses the CreateBulk function to create multiple GroupMembership entities
func (r *mutationResolver) bulkCreateGroupMembership(ctx context.Context, input []*generated.CreateGroupMembershipInput) (*GroupMembershipBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.GroupMembershipCreate, len(input))
	for i, data := range input {
		builders[i] = c.GroupMembership.Create().SetInput(*data)
	}

	res, err := c.GroupMembership.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "groupmembership"}, r.logger)
	}

	// return response
	return &GroupMembershipBulkCreatePayload{
		GroupMemberships: res,
	}, nil
}

// bulkCreateGroupSetting uses the CreateBulk function to create multiple GroupSetting entities
func (r *mutationResolver) bulkCreateGroupSetting(ctx context.Context, input []*generated.CreateGroupSettingInput) (*GroupSettingBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.GroupSettingCreate, len(input))
	for i, data := range input {
		builders[i] = c.GroupSetting.Create().SetInput(*data)
	}

	res, err := c.GroupSetting.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "groupsetting"}, r.logger)
	}

	// return response
	return &GroupSettingBulkCreatePayload{
		GroupSettings: res,
	}, nil
}

// bulkCreateHush uses the CreateBulk function to create multiple Hush entities
func (r *mutationResolver) bulkCreateHush(ctx context.Context, input []*generated.CreateHushInput) (*HushBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.HushCreate, len(input))
	for i, data := range input {
		builders[i] = c.Hush.Create().SetInput(*data)
	}

	res, err := c.Hush.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "hush"}, r.logger)
	}

	// return response
	return &HushBulkCreatePayload{
		Hushes: res,
	}, nil
}

// bulkCreateIntegration uses the CreateBulk function to create multiple Integration entities
func (r *mutationResolver) bulkCreateIntegration(ctx context.Context, input []*generated.CreateIntegrationInput) (*IntegrationBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.IntegrationCreate, len(input))
	for i, data := range input {
		builders[i] = c.Integration.Create().SetInput(*data)
	}

	res, err := c.Integration.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "integration"}, r.logger)
	}

	// return response
	return &IntegrationBulkCreatePayload{
		Integrations: res,
	}, nil
}

// bulkCreateInvite uses the CreateBulk function to create multiple Invite entities
func (r *mutationResolver) bulkCreateInvite(ctx context.Context, input []*generated.CreateInviteInput) (*InviteBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.InviteCreate, len(input))
	for i, data := range input {
		builders[i] = c.Invite.Create().SetInput(*data)
	}

	res, err := c.Invite.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "invite"}, r.logger)
	}

	// return response
	return &InviteBulkCreatePayload{
		Invites: res,
	}, nil
}

// bulkCreateOauthProvider uses the CreateBulk function to create multiple OauthProvider entities
func (r *mutationResolver) bulkCreateOauthProvider(ctx context.Context, input []*generated.CreateOauthProviderInput) (*OauthProviderBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OauthProviderCreate, len(input))
	for i, data := range input {
		builders[i] = c.OauthProvider.Create().SetInput(*data)
	}

	res, err := c.OauthProvider.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "oauthprovider"}, r.logger)
	}

	// return response
	return &OauthProviderBulkCreatePayload{
		OauthProviders: res,
	}, nil
}

// bulkCreateOhAuthTooToken uses the CreateBulk function to create multiple OhAuthTooToken entities
func (r *mutationResolver) bulkCreateOhAuthTooToken(ctx context.Context, input []*generated.CreateOhAuthTooTokenInput) (*OhAuthTooTokenBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OhAuthTooTokenCreate, len(input))
	for i, data := range input {
		builders[i] = c.OhAuthTooToken.Create().SetInput(*data)
	}

	res, err := c.OhAuthTooToken.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "ohauthtootoken"}, r.logger)
	}

	// return response
	return &OhAuthTooTokenBulkCreatePayload{
		OhAuthTooTokens: res,
	}, nil
}

// bulkCreateOrganization uses the CreateBulk function to create multiple Organization entities
func (r *mutationResolver) bulkCreateOrganization(ctx context.Context, input []*generated.CreateOrganizationInput) (*OrganizationBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OrganizationCreate, len(input))
	for i, data := range input {
		builders[i] = c.Organization.Create().SetInput(*data)
	}

	res, err := c.Organization.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "organization"}, r.logger)
	}

	// return response
	return &OrganizationBulkCreatePayload{
		Organizations: res,
	}, nil
}

// bulkCreateOrganizationSetting uses the CreateBulk function to create multiple OrganizationSetting entities
func (r *mutationResolver) bulkCreateOrganizationSetting(ctx context.Context, input []*generated.CreateOrganizationSettingInput) (*OrganizationSettingBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OrganizationSettingCreate, len(input))
	for i, data := range input {
		builders[i] = c.OrganizationSetting.Create().SetInput(*data)
	}

	res, err := c.OrganizationSetting.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "organizationsetting"}, r.logger)
	}

	// return response
	return &OrganizationSettingBulkCreatePayload{
		OrganizationSettings: res,
	}, nil
}

// bulkCreateOrgMembership uses the CreateBulk function to create multiple OrgMembership entities
func (r *mutationResolver) bulkCreateOrgMembership(ctx context.Context, input []*generated.CreateOrgMembershipInput) (*OrgMembershipBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OrgMembershipCreate, len(input))
	for i, data := range input {
		builders[i] = c.OrgMembership.Create().SetInput(*data)
	}

	res, err := c.OrgMembership.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "orgmembership"}, r.logger)
	}

	// return response
	return &OrgMembershipBulkCreatePayload{
		OrgMemberships: res,
	}, nil
}

// bulkCreatePersonalAccessToken uses the CreateBulk function to create multiple PersonalAccessToken entities
func (r *mutationResolver) bulkCreatePersonalAccessToken(ctx context.Context, input []*generated.CreatePersonalAccessTokenInput) (*PersonalAccessTokenBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.PersonalAccessTokenCreate, len(input))
	for i, data := range input {
		builders[i] = c.PersonalAccessToken.Create().SetInput(*data)
	}

	res, err := c.PersonalAccessToken.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "personalaccesstoken"}, r.logger)
	}

	// return response
	return &PersonalAccessTokenBulkCreatePayload{
		PersonalAccessTokens: res,
	}, nil
}

// bulkCreateSubscriber uses the CreateBulk function to create multiple Subscriber entities
func (r *mutationResolver) bulkCreateSubscriber(ctx context.Context, input []*generated.CreateSubscriberInput) (*SubscriberBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.SubscriberCreate, len(input))
	for i, data := range input {
		builders[i] = c.Subscriber.Create().SetInput(*data)
	}

	res, err := c.Subscriber.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "subscriber"}, r.logger)
	}

	// return response
	return &SubscriberBulkCreatePayload{
		Subscribers: res,
	}, nil
}

// bulkCreateTemplate uses the CreateBulk function to create multiple Template entities
func (r *mutationResolver) bulkCreateTemplate(ctx context.Context, input []*generated.CreateTemplateInput) (*TemplateBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.TemplateCreate, len(input))
	for i, data := range input {
		builders[i] = c.Template.Create().SetInput(*data)
	}

	res, err := c.Template.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "template"}, r.logger)
	}

	// return response
	return &TemplateBulkCreatePayload{
		Templates: res,
	}, nil
}

// bulkCreateUser uses the CreateBulk function to create multiple User entities
func (r *mutationResolver) bulkCreateUser(ctx context.Context, input []*generated.CreateUserInput) (*UserBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.UserCreate, len(input))
	for i, data := range input {
		builders[i] = c.User.Create().SetInput(*data)
	}

	res, err := c.User.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "user"}, r.logger)
	}

	// return response
	return &UserBulkCreatePayload{
		Users: res,
	}, nil
}

// bulkCreateUserSetting uses the CreateBulk function to create multiple UserSetting entities
func (r *mutationResolver) bulkCreateUserSetting(ctx context.Context, input []*generated.CreateUserSettingInput) (*UserSettingBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.UserSettingCreate, len(input))
	for i, data := range input {
		builders[i] = c.UserSetting.Create().SetInput(*data)
	}

	res, err := c.UserSetting.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "usersetting"}, r.logger)
	}

	// return response
	return &UserSettingBulkCreatePayload{
		UserSettings: res,
	}, nil
}

// bulkCreateWebhook uses the CreateBulk function to create multiple Webhook entities
func (r *mutationResolver) bulkCreateWebhook(ctx context.Context, input []*generated.CreateWebhookInput) (*WebhookBulkCreatePayload, error) {
	c := withTransactionalMutation(ctx)
	builders := make([]*generated.WebhookCreate, len(input))
	for i, data := range input {
		builders[i] = c.Webhook.Create().SetInput(*data)
	}

	res, err := c.Webhook.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "webhook"}, r.logger)
	}

	// return response
	return &WebhookBulkCreatePayload{
		Webhooks: res,
	}, nil
}
