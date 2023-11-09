// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettings"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/refreshtoken"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
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
	entitlementDescCancelled := entitlementFields[4].Descriptor()
	// entitlement.DefaultCancelled holds the default value on creation for the cancelled field.
	entitlement.DefaultCancelled = entitlementDescCancelled.Default.(bool)
	// entitlementDescID is the schema descriptor for id field.
	entitlementDescID := entitlementMixinFields1[0].Descriptor()
	// entitlement.DefaultID holds the default value on creation for the id field.
	entitlement.DefaultID = entitlementDescID.Default.(func() string)
	groupMixin := schema.Group{}.Mixin()
	group.Policy = privacy.NewPolicies(groupMixin[1], schema.Group{})
	group.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := group.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	groupMixinHooks0 := groupMixin[0].Hooks()

	group.Hooks[1] = groupMixinHooks0[0]
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
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupMixinFields2[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() string)
	groupsettingsMixin := schema.GroupSettings{}.Mixin()
	groupsettingsMixinHooks0 := groupsettingsMixin[0].Hooks()
	groupsettings.Hooks[0] = groupsettingsMixinHooks0[0]
	groupsettingsMixinFields0 := groupsettingsMixin[0].Fields()
	_ = groupsettingsMixinFields0
	groupsettingsMixinFields1 := groupsettingsMixin[1].Fields()
	_ = groupsettingsMixinFields1
	groupsettingsFields := schema.GroupSettings{}.Fields()
	_ = groupsettingsFields
	// groupsettingsDescCreatedAt is the schema descriptor for created_at field.
	groupsettingsDescCreatedAt := groupsettingsMixinFields0[0].Descriptor()
	// groupsettings.DefaultCreatedAt holds the default value on creation for the created_at field.
	groupsettings.DefaultCreatedAt = groupsettingsDescCreatedAt.Default.(func() time.Time)
	// groupsettingsDescUpdatedAt is the schema descriptor for updated_at field.
	groupsettingsDescUpdatedAt := groupsettingsMixinFields0[1].Descriptor()
	// groupsettings.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	groupsettings.DefaultUpdatedAt = groupsettingsDescUpdatedAt.Default.(func() time.Time)
	// groupsettings.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	groupsettings.UpdateDefaultUpdatedAt = groupsettingsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// groupsettingsDescID is the schema descriptor for id field.
	groupsettingsDescID := groupsettingsMixinFields1[0].Descriptor()
	// groupsettings.DefaultID holds the default value on creation for the id field.
	groupsettings.DefaultID = groupsettingsDescID.Default.(func() string)
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
	organizationMixinHooks0 := organizationMixin[0].Hooks()
	organization.Hooks[0] = organizationMixinHooks0[0]
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
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields1[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() string)
	organizationsettingsMixin := schema.OrganizationSettings{}.Mixin()
	organizationsettingsMixinHooks0 := organizationsettingsMixin[0].Hooks()
	organizationsettings.Hooks[0] = organizationsettingsMixinHooks0[0]
	organizationsettingsMixinFields0 := organizationsettingsMixin[0].Fields()
	_ = organizationsettingsMixinFields0
	organizationsettingsMixinFields1 := organizationsettingsMixin[1].Fields()
	_ = organizationsettingsMixinFields1
	organizationsettingsFields := schema.OrganizationSettings{}.Fields()
	_ = organizationsettingsFields
	// organizationsettingsDescCreatedAt is the schema descriptor for created_at field.
	organizationsettingsDescCreatedAt := organizationsettingsMixinFields0[0].Descriptor()
	// organizationsettings.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationsettings.DefaultCreatedAt = organizationsettingsDescCreatedAt.Default.(func() time.Time)
	// organizationsettingsDescUpdatedAt is the schema descriptor for updated_at field.
	organizationsettingsDescUpdatedAt := organizationsettingsMixinFields0[1].Descriptor()
	// organizationsettings.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	organizationsettings.DefaultUpdatedAt = organizationsettingsDescUpdatedAt.Default.(func() time.Time)
	// organizationsettings.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	organizationsettings.UpdateDefaultUpdatedAt = organizationsettingsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// organizationsettingsDescSSOCert is the schema descriptor for sso_cert field.
	organizationsettingsDescSSOCert := organizationsettingsFields[1].Descriptor()
	// organizationsettings.DefaultSSOCert holds the default value on creation for the sso_cert field.
	organizationsettings.DefaultSSOCert = organizationsettingsDescSSOCert.Default.(string)
	// organizationsettingsDescSSOEntrypoint is the schema descriptor for sso_entrypoint field.
	organizationsettingsDescSSOEntrypoint := organizationsettingsFields[2].Descriptor()
	// organizationsettings.DefaultSSOEntrypoint holds the default value on creation for the sso_entrypoint field.
	organizationsettings.DefaultSSOEntrypoint = organizationsettingsDescSSOEntrypoint.Default.(string)
	// organizationsettingsDescSSOIssuer is the schema descriptor for sso_issuer field.
	organizationsettingsDescSSOIssuer := organizationsettingsFields[3].Descriptor()
	// organizationsettings.DefaultSSOIssuer holds the default value on creation for the sso_issuer field.
	organizationsettings.DefaultSSOIssuer = organizationsettingsDescSSOIssuer.Default.(string)
	// organizationsettingsDescBillingContact is the schema descriptor for billing_contact field.
	organizationsettingsDescBillingContact := organizationsettingsFields[4].Descriptor()
	// organizationsettings.BillingContactValidator is a validator for the "billing_contact" field. It is called by the builders before save.
	organizationsettings.BillingContactValidator = organizationsettingsDescBillingContact.Validators[0].(func(string) error)
	// organizationsettingsDescBillingEmail is the schema descriptor for billing_email field.
	organizationsettingsDescBillingEmail := organizationsettingsFields[5].Descriptor()
	// organizationsettings.BillingEmailValidator is a validator for the "billing_email" field. It is called by the builders before save.
	organizationsettings.BillingEmailValidator = organizationsettingsDescBillingEmail.Validators[0].(func(string) error)
	// organizationsettingsDescBillingPhone is the schema descriptor for billing_phone field.
	organizationsettingsDescBillingPhone := organizationsettingsFields[6].Descriptor()
	// organizationsettings.BillingPhoneValidator is a validator for the "billing_phone" field. It is called by the builders before save.
	organizationsettings.BillingPhoneValidator = organizationsettingsDescBillingPhone.Validators[0].(func(string) error)
	// organizationsettingsDescBillingAddress is the schema descriptor for billing_address field.
	organizationsettingsDescBillingAddress := organizationsettingsFields[7].Descriptor()
	// organizationsettings.BillingAddressValidator is a validator for the "billing_address" field. It is called by the builders before save.
	organizationsettings.BillingAddressValidator = organizationsettingsDescBillingAddress.Validators[0].(func(string) error)
	// organizationsettingsDescID is the schema descriptor for id field.
	organizationsettingsDescID := organizationsettingsMixinFields1[0].Descriptor()
	// organizationsettings.DefaultID holds the default value on creation for the id field.
	organizationsettings.DefaultID = organizationsettingsDescID.Default.(func() string)
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
	user.Policy = privacy.NewPolicies(userMixin[1], schema.User{})
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
	// userDescLocked is the schema descriptor for locked field.
	userDescLocked := userFields[4].Descriptor()
	// user.DefaultLocked holds the default value on creation for the locked field.
	user.DefaultLocked = userDescLocked.Default.(bool)
	// userDescAvatarRemoteURL is the schema descriptor for avatar_remote_url field.
	userDescAvatarRemoteURL := userFields[5].Descriptor()
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
	userDescAvatarLocalFile := userFields[6].Descriptor()
	// user.AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	user.AvatarLocalFileValidator = userDescAvatarLocalFile.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields2[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
