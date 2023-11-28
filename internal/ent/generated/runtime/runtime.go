// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/refreshtoken"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/datumforge/datum/internal/ent/schema"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	entitlementMixin := schema.Entitlement{}.Mixin()
	entitlementMixinHooks0 := entitlementMixin[0].Hooks()
	entitlement.Hooks[0] = entitlementMixinHooks0[0]
	entitlementMixinFields0 := entitlementMixin[0].Fields()
	_ = entitlementMixinFields0
	entitlementMixinFields1 := entitlementMixin[1].Fields()
	_ = entitlementMixinFields1
	entitlementFields := schema.Entitlement{}.Fields()
	_ = entitlementFields
	// entitlementDescCreatedAt is the schema descriptor for created_at field.
	entitlementDescCreatedAt := entitlementMixinFields0[0].Descriptor()
	// entitlement.DefaultCreatedAt holds the default value on creation for the created_at field.
	entitlement.DefaultCreatedAt = entitlementDescCreatedAt.Default.(func() time.Time)
	// entitlementDescUpdatedAt is the schema descriptor for updated_at field.
	entitlementDescUpdatedAt := entitlementMixinFields0[1].Descriptor()
	// entitlement.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	entitlement.DefaultUpdatedAt = entitlementDescUpdatedAt.Default.(func() time.Time)
	// entitlement.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	entitlement.UpdateDefaultUpdatedAt = entitlementDescUpdatedAt.UpdateDefault.(func() time.Time)
	// entitlementDescCancelled is the schema descriptor for cancelled field.
	entitlementDescCancelled := entitlementFields[8].Descriptor()
	// entitlement.DefaultCancelled holds the default value on creation for the cancelled field.
	entitlement.DefaultCancelled = entitlementDescCancelled.Default.(bool)
	// entitlementDescID is the schema descriptor for id field.
	entitlementDescID := entitlementMixinFields1[0].Descriptor()
	// entitlement.DefaultID holds the default value on creation for the id field.
	entitlement.DefaultID = entitlementDescID.Default.(func() string)
	groupMixin := schema.Group{}.Mixin()
	groupMixinHooks0 := groupMixin[0].Hooks()
	group.Hooks[0] = groupMixinHooks0[0]
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
	groupMixinFields2 := groupMixin[2].Fields()
	_ = groupMixinFields2
	groupFields := schema.Group{}.Fields()
	_ = groupFields
	// groupDescCreatedAt is the schema descriptor for created_at field.
	groupDescCreatedAt := groupMixinFields0[0].Descriptor()
	// group.DefaultCreatedAt holds the default value on creation for the created_at field.
	group.DefaultCreatedAt = groupDescCreatedAt.Default.(func() time.Time)
	// groupDescUpdatedAt is the schema descriptor for updated_at field.
	groupDescUpdatedAt := groupMixinFields0[1].Descriptor()
	// group.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	group.DefaultUpdatedAt = groupDescUpdatedAt.Default.(func() time.Time)
	// group.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	group.UpdateDefaultUpdatedAt = groupDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupDescName is the schema descriptor for name field.
	groupDescName := groupFields[0].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescDescription is the schema descriptor for description field.
	groupDescDescription := groupFields[1].Descriptor()
	// group.DefaultDescription holds the default value on creation for the description field.
	group.DefaultDescription = groupDescDescription.Default.(string)
	// groupDescLogoURL is the schema descriptor for logo_url field.
	groupDescLogoURL := groupFields[2].Descriptor()
	// group.LogoURLValidator is a validator for the "logo_url" field. It is called by the builders before save.
	group.LogoURLValidator = groupDescLogoURL.Validators[0].(func(string) error)
	// groupDescDisplayName is the schema descriptor for display_name field.
	groupDescDisplayName := groupFields[3].Descriptor()
	// group.DefaultDisplayName holds the default value on creation for the display_name field.
	group.DefaultDisplayName = groupDescDisplayName.Default.(string)
	// group.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	group.DisplayNameValidator = func() func(string) error {
		validators := groupDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(display_name string) error {
			for _, fn := range fns {
				if err := fn(display_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupMixinFields2[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() string)
	groupsettingMixin := schema.GroupSetting{}.Mixin()
	groupsettingMixinHooks0 := groupsettingMixin[0].Hooks()
	groupsetting.Hooks[0] = groupsettingMixinHooks0[0]
	groupsettingMixinFields0 := groupsettingMixin[0].Fields()
	_ = groupsettingMixinFields0
	groupsettingMixinFields1 := groupsettingMixin[1].Fields()
	_ = groupsettingMixinFields1
	groupsettingFields := schema.GroupSetting{}.Fields()
	_ = groupsettingFields
	// groupsettingDescCreatedAt is the schema descriptor for created_at field.
	groupsettingDescCreatedAt := groupsettingMixinFields0[0].Descriptor()
	// groupsetting.DefaultCreatedAt holds the default value on creation for the created_at field.
	groupsetting.DefaultCreatedAt = groupsettingDescCreatedAt.Default.(func() time.Time)
	// groupsettingDescUpdatedAt is the schema descriptor for updated_at field.
	groupsettingDescUpdatedAt := groupsettingMixinFields0[1].Descriptor()
	// groupsetting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	groupsetting.DefaultUpdatedAt = groupsettingDescUpdatedAt.Default.(func() time.Time)
	// groupsetting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	groupsetting.UpdateDefaultUpdatedAt = groupsettingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupsettingDescTags is the schema descriptor for tags field.
	groupsettingDescTags := groupsettingFields[2].Descriptor()
	// groupsetting.DefaultTags holds the default value on creation for the tags field.
	groupsetting.DefaultTags = groupsettingDescTags.Default.([]string)
	// groupsettingDescSyncToSlack is the schema descriptor for sync_to_slack field.
	groupsettingDescSyncToSlack := groupsettingFields[3].Descriptor()
	// groupsetting.DefaultSyncToSlack holds the default value on creation for the sync_to_slack field.
	groupsetting.DefaultSyncToSlack = groupsettingDescSyncToSlack.Default.(bool)
	// groupsettingDescSyncToGithub is the schema descriptor for sync_to_github field.
	groupsettingDescSyncToGithub := groupsettingFields[4].Descriptor()
	// groupsetting.DefaultSyncToGithub holds the default value on creation for the sync_to_github field.
	groupsetting.DefaultSyncToGithub = groupsettingDescSyncToGithub.Default.(bool)
	// groupsettingDescID is the schema descriptor for id field.
	groupsettingDescID := groupsettingMixinFields1[0].Descriptor()
	// groupsetting.DefaultID holds the default value on creation for the id field.
	groupsetting.DefaultID = groupsettingDescID.Default.(func() string)
	integrationMixin := schema.Integration{}.Mixin()
	integrationMixinHooks0 := integrationMixin[0].Hooks()
	integration.Hooks[0] = integrationMixinHooks0[0]
	integrationMixinFields0 := integrationMixin[0].Fields()
	_ = integrationMixinFields0
	integrationMixinFields1 := integrationMixin[1].Fields()
	_ = integrationMixinFields1
	integrationFields := schema.Integration{}.Fields()
	_ = integrationFields
	// integrationDescCreatedAt is the schema descriptor for created_at field.
	integrationDescCreatedAt := integrationMixinFields0[0].Descriptor()
	// integration.DefaultCreatedAt holds the default value on creation for the created_at field.
	integration.DefaultCreatedAt = integrationDescCreatedAt.Default.(func() time.Time)
	// integrationDescUpdatedAt is the schema descriptor for updated_at field.
	integrationDescUpdatedAt := integrationMixinFields0[1].Descriptor()
	// integration.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	integration.DefaultUpdatedAt = integrationDescUpdatedAt.Default.(func() time.Time)
	// integration.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	integration.UpdateDefaultUpdatedAt = integrationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// integrationDescName is the schema descriptor for name field.
	integrationDescName := integrationFields[0].Descriptor()
	// integration.NameValidator is a validator for the "name" field. It is called by the builders before save.
	integration.NameValidator = integrationDescName.Validators[0].(func(string) error)
	// integrationDescID is the schema descriptor for id field.
	integrationDescID := integrationMixinFields1[0].Descriptor()
	// integration.DefaultID holds the default value on creation for the id field.
	integration.DefaultID = integrationDescID.Default.(func() string)
	oauthproviderMixin := schema.OauthProvider{}.Mixin()
	oauthproviderMixinHooks0 := oauthproviderMixin[0].Hooks()
	oauthprovider.Hooks[0] = oauthproviderMixinHooks0[0]
	oauthproviderMixinFields0 := oauthproviderMixin[0].Fields()
	_ = oauthproviderMixinFields0
	oauthproviderMixinFields1 := oauthproviderMixin[1].Fields()
	_ = oauthproviderMixinFields1
	oauthproviderFields := schema.OauthProvider{}.Fields()
	_ = oauthproviderFields
	// oauthproviderDescCreatedAt is the schema descriptor for created_at field.
	oauthproviderDescCreatedAt := oauthproviderMixinFields0[0].Descriptor()
	// oauthprovider.DefaultCreatedAt holds the default value on creation for the created_at field.
	oauthprovider.DefaultCreatedAt = oauthproviderDescCreatedAt.Default.(func() time.Time)
	// oauthproviderDescUpdatedAt is the schema descriptor for updated_at field.
	oauthproviderDescUpdatedAt := oauthproviderMixinFields0[1].Descriptor()
	// oauthprovider.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	oauthprovider.DefaultUpdatedAt = oauthproviderDescUpdatedAt.Default.(func() time.Time)
	// oauthprovider.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	oauthprovider.UpdateDefaultUpdatedAt = oauthproviderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// oauthproviderDescID is the schema descriptor for id field.
	oauthproviderDescID := oauthproviderMixinFields1[0].Descriptor()
	// oauthprovider.DefaultID holds the default value on creation for the id field.
	oauthprovider.DefaultID = oauthproviderDescID.Default.(func() string)
	organizationMixin := schema.Organization{}.Mixin()
	organization.Policy = privacy.NewPolicies(schema.Organization{})
	organization.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := organization.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	organizationMixinHooks0 := organizationMixin[0].Hooks()
	organizationHooks := schema.Organization{}.Hooks()

	organization.Hooks[1] = organizationMixinHooks0[0]

	organization.Hooks[2] = organizationHooks[0]
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationMixinFields1 := organizationMixin[1].Fields()
	_ = organizationMixinFields1
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationMixinFields0[0].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescUpdatedAt is the schema descriptor for updated_at field.
	organizationDescUpdatedAt := organizationMixinFields0[1].Descriptor()
	// organization.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organization.DefaultUpdatedAt = organizationDescUpdatedAt.Default.(func() time.Time)
	// organization.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organization.UpdateDefaultUpdatedAt = organizationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[0].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = func() func(string) error {
		validators := organizationDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// organizationDescDisplayName is the schema descriptor for display_name field.
	organizationDescDisplayName := organizationFields[1].Descriptor()
	// organization.DefaultDisplayName holds the default value on creation for the display_name field.
	organization.DefaultDisplayName = organizationDescDisplayName.Default.(string)
	// organization.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	organization.DisplayNameValidator = func() func(string) error {
		validators := organizationDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(display_name string) error {
			for _, fn := range fns {
				if err := fn(display_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields1[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() string)
	organizationsettingMixin := schema.OrganizationSetting{}.Mixin()
	organizationsettingMixinHooks0 := organizationsettingMixin[0].Hooks()
	organizationsetting.Hooks[0] = organizationsettingMixinHooks0[0]
	organizationsettingMixinFields0 := organizationsettingMixin[0].Fields()
	_ = organizationsettingMixinFields0
	organizationsettingMixinFields1 := organizationsettingMixin[1].Fields()
	_ = organizationsettingMixinFields1
	organizationsettingFields := schema.OrganizationSetting{}.Fields()
	_ = organizationsettingFields
	// organizationsettingDescCreatedAt is the schema descriptor for created_at field.
	organizationsettingDescCreatedAt := organizationsettingMixinFields0[0].Descriptor()
	// organizationsetting.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationsetting.DefaultCreatedAt = organizationsettingDescCreatedAt.Default.(func() time.Time)
	// organizationsettingDescUpdatedAt is the schema descriptor for updated_at field.
	organizationsettingDescUpdatedAt := organizationsettingMixinFields0[1].Descriptor()
	// organizationsetting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organizationsetting.DefaultUpdatedAt = organizationsettingDescUpdatedAt.Default.(func() time.Time)
	// organizationsetting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organizationsetting.UpdateDefaultUpdatedAt = organizationsettingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// organizationsettingDescSSOCert is the schema descriptor for sso_cert field.
	organizationsettingDescSSOCert := organizationsettingFields[1].Descriptor()
	// organizationsetting.DefaultSSOCert holds the default value on creation for the sso_cert field.
	organizationsetting.DefaultSSOCert = organizationsettingDescSSOCert.Default.(string)
	// organizationsettingDescSSOEntrypoint is the schema descriptor for sso_entrypoint field.
	organizationsettingDescSSOEntrypoint := organizationsettingFields[2].Descriptor()
	// organizationsetting.DefaultSSOEntrypoint holds the default value on creation for the sso_entrypoint field.
	organizationsetting.DefaultSSOEntrypoint = organizationsettingDescSSOEntrypoint.Default.(string)
	// organizationsettingDescSSOIssuer is the schema descriptor for sso_issuer field.
	organizationsettingDescSSOIssuer := organizationsettingFields[3].Descriptor()
	// organizationsetting.DefaultSSOIssuer holds the default value on creation for the sso_issuer field.
	organizationsetting.DefaultSSOIssuer = organizationsettingDescSSOIssuer.Default.(string)
	// organizationsettingDescBillingContact is the schema descriptor for billing_contact field.
	organizationsettingDescBillingContact := organizationsettingFields[4].Descriptor()
	// organizationsetting.BillingContactValidator is a validator for the "billing_contact" field. It is called by the builders before save.
	organizationsetting.BillingContactValidator = organizationsettingDescBillingContact.Validators[0].(func(string) error)
	// organizationsettingDescBillingEmail is the schema descriptor for billing_email field.
	organizationsettingDescBillingEmail := organizationsettingFields[5].Descriptor()
	// organizationsetting.BillingEmailValidator is a validator for the "billing_email" field. It is called by the builders before save.
	organizationsetting.BillingEmailValidator = organizationsettingDescBillingEmail.Validators[0].(func(string) error)
	// organizationsettingDescBillingPhone is the schema descriptor for billing_phone field.
	organizationsettingDescBillingPhone := organizationsettingFields[6].Descriptor()
	// organizationsetting.BillingPhoneValidator is a validator for the "billing_phone" field. It is called by the builders before save.
	organizationsetting.BillingPhoneValidator = organizationsettingDescBillingPhone.Validators[0].(func(string) error)
	// organizationsettingDescBillingAddress is the schema descriptor for billing_address field.
	organizationsettingDescBillingAddress := organizationsettingFields[7].Descriptor()
	// organizationsetting.BillingAddressValidator is a validator for the "billing_address" field. It is called by the builders before save.
	organizationsetting.BillingAddressValidator = organizationsettingDescBillingAddress.Validators[0].(func(string) error)
	// organizationsettingDescTags is the schema descriptor for tags field.
	organizationsettingDescTags := organizationsettingFields[9].Descriptor()
	// organizationsetting.DefaultTags holds the default value on creation for the tags field.
	organizationsetting.DefaultTags = organizationsettingDescTags.Default.([]string)
	// organizationsettingDescID is the schema descriptor for id field.
	organizationsettingDescID := organizationsettingMixinFields1[0].Descriptor()
	// organizationsetting.DefaultID holds the default value on creation for the id field.
	organizationsetting.DefaultID = organizationsettingDescID.Default.(func() string)
	personalaccesstokenMixin := schema.PersonalAccessToken{}.Mixin()
	personalaccesstokenMixinHooks0 := personalaccesstokenMixin[0].Hooks()
	personalaccesstoken.Hooks[0] = personalaccesstokenMixinHooks0[0]
	personalaccesstokenMixinFields0 := personalaccesstokenMixin[0].Fields()
	_ = personalaccesstokenMixinFields0
	personalaccesstokenMixinFields1 := personalaccesstokenMixin[1].Fields()
	_ = personalaccesstokenMixinFields1
	personalaccesstokenFields := schema.PersonalAccessToken{}.Fields()
	_ = personalaccesstokenFields
	// personalaccesstokenDescCreatedAt is the schema descriptor for created_at field.
	personalaccesstokenDescCreatedAt := personalaccesstokenMixinFields0[0].Descriptor()
	// personalaccesstoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	personalaccesstoken.DefaultCreatedAt = personalaccesstokenDescCreatedAt.Default.(func() time.Time)
	// personalaccesstokenDescUpdatedAt is the schema descriptor for updated_at field.
	personalaccesstokenDescUpdatedAt := personalaccesstokenMixinFields0[1].Descriptor()
	// personalaccesstoken.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	personalaccesstoken.DefaultUpdatedAt = personalaccesstokenDescUpdatedAt.Default.(func() time.Time)
	// personalaccesstoken.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	personalaccesstoken.UpdateDefaultUpdatedAt = personalaccesstokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// personalaccesstokenDescLastUsedAt is the schema descriptor for last_used_at field.
	personalaccesstokenDescLastUsedAt := personalaccesstokenFields[5].Descriptor()
	// personalaccesstoken.UpdateDefaultLastUsedAt holds the default value on update for the last_used_at field.
	personalaccesstoken.UpdateDefaultLastUsedAt = personalaccesstokenDescLastUsedAt.UpdateDefault.(func() time.Time)
	// personalaccesstokenDescID is the schema descriptor for id field.
	personalaccesstokenDescID := personalaccesstokenMixinFields1[0].Descriptor()
	// personalaccesstoken.DefaultID holds the default value on creation for the id field.
	personalaccesstoken.DefaultID = personalaccesstokenDescID.Default.(func() string)
	refreshtokenMixin := schema.RefreshToken{}.Mixin()
	refreshtokenMixinFields0 := refreshtokenMixin[0].Fields()
	_ = refreshtokenMixinFields0
	refreshtokenFields := schema.RefreshToken{}.Fields()
	_ = refreshtokenFields
	// refreshtokenDescClientID is the schema descriptor for client_id field.
	refreshtokenDescClientID := refreshtokenFields[0].Descriptor()
	// refreshtoken.ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	refreshtoken.ClientIDValidator = refreshtokenDescClientID.Validators[0].(func(string) error)
	// refreshtokenDescNonce is the schema descriptor for nonce field.
	refreshtokenDescNonce := refreshtokenFields[2].Descriptor()
	// refreshtoken.NonceValidator is a validator for the "nonce" field. It is called by the builders before save.
	refreshtoken.NonceValidator = refreshtokenDescNonce.Validators[0].(func(string) error)
	// refreshtokenDescClaimsUserID is the schema descriptor for claims_user_id field.
	refreshtokenDescClaimsUserID := refreshtokenFields[3].Descriptor()
	// refreshtoken.ClaimsUserIDValidator is a validator for the "claims_user_id" field. It is called by the builders before save.
	refreshtoken.ClaimsUserIDValidator = refreshtokenDescClaimsUserID.Validators[0].(func(string) error)
	// refreshtokenDescClaimsUsername is the schema descriptor for claims_username field.
	refreshtokenDescClaimsUsername := refreshtokenFields[4].Descriptor()
	// refreshtoken.ClaimsUsernameValidator is a validator for the "claims_username" field. It is called by the builders before save.
	refreshtoken.ClaimsUsernameValidator = refreshtokenDescClaimsUsername.Validators[0].(func(string) error)
	// refreshtokenDescClaimsEmail is the schema descriptor for claims_email field.
	refreshtokenDescClaimsEmail := refreshtokenFields[5].Descriptor()
	// refreshtoken.ClaimsEmailValidator is a validator for the "claims_email" field. It is called by the builders before save.
	refreshtoken.ClaimsEmailValidator = refreshtokenDescClaimsEmail.Validators[0].(func(string) error)
	// refreshtokenDescConnectorID is the schema descriptor for connector_id field.
	refreshtokenDescConnectorID := refreshtokenFields[9].Descriptor()
	// refreshtoken.ConnectorIDValidator is a validator for the "connector_id" field. It is called by the builders before save.
	refreshtoken.ConnectorIDValidator = refreshtokenDescConnectorID.Validators[0].(func(string) error)
	// refreshtokenDescLastUsed is the schema descriptor for last_used field.
	refreshtokenDescLastUsed := refreshtokenFields[13].Descriptor()
	// refreshtoken.DefaultLastUsed holds the default value on creation for the last_used field.
	refreshtoken.DefaultLastUsed = refreshtokenDescLastUsed.Default.(func() time.Time)
	// refreshtokenDescID is the schema descriptor for id field.
	refreshtokenDescID := refreshtokenMixinFields0[0].Descriptor()
	// refreshtoken.DefaultID holds the default value on creation for the id field.
	refreshtoken.DefaultID = refreshtokenDescID.Default.(func() string)
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinHooks0 := sessionMixin[0].Hooks()
	session.Hooks[0] = sessionMixinHooks0[0]
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
	sessionMixinFields1 := sessionMixin[1].Fields()
	_ = sessionMixinFields1
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for created_at field.
	sessionDescCreatedAt := sessionMixinFields0[0].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the created_at field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updated_at field.
	sessionDescUpdatedAt := sessionMixinFields0[1].Descriptor()
	// session.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	session.DefaultUpdatedAt = sessionDescUpdatedAt.Default.(func() time.Time)
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// sessionDescToken is the schema descriptor for token field.
	sessionDescToken := sessionFields[2].Descriptor()
	// session.DefaultToken holds the default value on creation for the token field.
	session.DefaultToken = sessionDescToken.Default.(func() string)
	// session.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	session.TokenValidator = sessionDescToken.Validators[0].(func(string) error)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionMixinFields1[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() string)
	userMixin := schema.User{}.Mixin()
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userMixinHooks0 := userMixin[0].Hooks()

	user.Hooks[1] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields2 := userMixin[2].Fields()
	_ = userMixinFields2
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[1].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = func() func(string) error {
		validators := userDescFirstName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(first_name string) error {
			for _, fn := range fns {
				if err := fn(first_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[2].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = func() func(string) error {
		validators := userDescLastName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(last_name string) error {
			for _, fn := range fns {
				if err := fn(last_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescDisplayName is the schema descriptor for display_name field.
	userDescDisplayName := userFields[3].Descriptor()
	// user.DefaultDisplayName holds the default value on creation for the display_name field.
	user.DefaultDisplayName = userDescDisplayName.Default.(string)
	// user.DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	user.DisplayNameValidator = func() func(string) error {
		validators := userDescDisplayName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(display_name string) error {
			for _, fn := range fns {
				if err := fn(display_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAvatarRemoteURL is the schema descriptor for avatar_remote_url field.
	userDescAvatarRemoteURL := userFields[4].Descriptor()
	// user.AvatarRemoteURLValidator is a validator for the "avatar_remote_url" field. It is called by the builders before save.
	user.AvatarRemoteURLValidator = func() func(string) error {
		validators := userDescAvatarRemoteURL.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(avatar_remote_url string) error {
			for _, fn := range fns {
				if err := fn(avatar_remote_url); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescAvatarLocalFile is the schema descriptor for avatar_local_file field.
	userDescAvatarLocalFile := userFields[5].Descriptor()
	// user.AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	user.AvatarLocalFileValidator = userDescAvatarLocalFile.Validators[0].(func(string) error)
	// userDescAvatarUpdatedAt is the schema descriptor for avatar_updated_at field.
	userDescAvatarUpdatedAt := userFields[6].Descriptor()
	// user.UpdateDefaultAvatarUpdatedAt holds the default value on update for the avatar_updated_at field.
	user.UpdateDefaultAvatarUpdatedAt = userDescAvatarUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescLastSeen is the schema descriptor for last_seen field.
	userDescLastSeen := userFields[7].Descriptor()
	// user.UpdateDefaultLastSeen holds the default value on update for the last_seen field.
	user.UpdateDefaultLastSeen = userDescLastSeen.UpdateDefault.(func() time.Time)
	// userDescOauth is the schema descriptor for oauth field.
	userDescOauth := userFields[10].Descriptor()
	// user.DefaultOauth holds the default value on creation for the oauth field.
	user.DefaultOauth = userDescOauth.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields2[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
	usersettingMixin := schema.UserSetting{}.Mixin()
	usersettingMixinHooks0 := usersettingMixin[0].Hooks()
	usersetting.Hooks[0] = usersettingMixinHooks0[0]
	usersettingMixinFields0 := usersettingMixin[0].Fields()
	_ = usersettingMixinFields0
	usersettingMixinFields1 := usersettingMixin[1].Fields()
	_ = usersettingMixinFields1
	usersettingFields := schema.UserSetting{}.Fields()
	_ = usersettingFields
	// usersettingDescCreatedAt is the schema descriptor for created_at field.
	usersettingDescCreatedAt := usersettingMixinFields0[0].Descriptor()
	// usersetting.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersetting.DefaultCreatedAt = usersettingDescCreatedAt.Default.(func() time.Time)
	// usersettingDescUpdatedAt is the schema descriptor for updated_at field.
	usersettingDescUpdatedAt := usersettingMixinFields0[1].Descriptor()
	// usersetting.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usersetting.DefaultUpdatedAt = usersettingDescUpdatedAt.Default.(func() time.Time)
	// usersetting.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usersetting.UpdateDefaultUpdatedAt = usersettingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// usersettingDescLocked is the schema descriptor for locked field.
	usersettingDescLocked := usersettingFields[0].Descriptor()
	// usersetting.DefaultLocked holds the default value on creation for the locked field.
	usersetting.DefaultLocked = usersettingDescLocked.Default.(bool)
	// usersettingDescPermissions is the schema descriptor for permissions field.
	usersettingDescPermissions := usersettingFields[6].Descriptor()
	// usersetting.DefaultPermissions holds the default value on creation for the permissions field.
	usersetting.DefaultPermissions = usersettingDescPermissions.Default.([]string)
	// usersettingDescEmailConfirmed is the schema descriptor for email_confirmed field.
	usersettingDescEmailConfirmed := usersettingFields[7].Descriptor()
	// usersetting.DefaultEmailConfirmed holds the default value on creation for the email_confirmed field.
	usersetting.DefaultEmailConfirmed = usersettingDescEmailConfirmed.Default.(bool)
	// usersettingDescTags is the schema descriptor for tags field.
	usersettingDescTags := usersettingFields[8].Descriptor()
	// usersetting.DefaultTags holds the default value on creation for the tags field.
	usersetting.DefaultTags = usersettingDescTags.Default.([]string)
	// usersettingDescID is the schema descriptor for id field.
	usersettingDescID := usersettingMixinFields1[0].Descriptor()
	// usersetting.DefaultID holds the default value on creation for the id field.
	usersetting.DefaultID = usersettingDescID.Default.(func() string)
}

const (
	Version = "v0.12.5"                                         // Version of ent codegen.
	Sum     = "h1:KREM5E4CSoej4zeGa88Ou/gfturAnpUv0mzAjch1sj4=" // Sum of ent codegen.
)
