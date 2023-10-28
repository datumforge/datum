// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/tenant"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	groupMixin := schema.Group{}.Mixin()
	groupMixinHooks0 := groupMixin[0].Hooks()
	group.Hooks[0] = groupMixinHooks0[0]
	groupMixinFields0 := groupMixin[0].Fields()
	_ = groupMixinFields0
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
	groupDescName := groupFields[1].Descriptor()
	// group.NameValidator is a validator for the "name" field. It is called by the builders before save.
	group.NameValidator = groupDescName.Validators[0].(func(string) error)
	// groupDescDescription is the schema descriptor for description field.
	groupDescDescription := groupFields[2].Descriptor()
	// group.DefaultDescription holds the default value on creation for the description field.
	group.DefaultDescription = groupDescDescription.Default.(string)
	// groupDescLogoURL is the schema descriptor for logo_url field.
	groupDescLogoURL := groupFields[3].Descriptor()
	// group.LogoURLValidator is a validator for the "logo_url" field. It is called by the builders before save.
	group.LogoURLValidator = groupDescLogoURL.Validators[0].(func(string) error)
	// groupDescID is the schema descriptor for id field.
	groupDescID := groupFields[0].Descriptor()
	// group.DefaultID holds the default value on creation for the id field.
	group.DefaultID = groupDescID.Default.(func() uuid.UUID)
	groupsettingsMixin := schema.GroupSettings{}.Mixin()
	groupsettingsMixinHooks0 := groupsettingsMixin[0].Hooks()
	groupsettings.Hooks[0] = groupsettingsMixinHooks0[0]
	groupsettingsMixinFields0 := groupsettingsMixin[0].Fields()
	_ = groupsettingsMixinFields0
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
	groupsettingsDescID := groupsettingsFields[0].Descriptor()
	// groupsettings.DefaultID holds the default value on creation for the id field.
	groupsettings.DefaultID = groupsettingsDescID.Default.(func() uuid.UUID)
	integrationMixin := schema.Integration{}.Mixin()
	integrationMixinHooks0 := integrationMixin[0].Hooks()
	integration.Hooks[0] = integrationMixinHooks0[0]
	integrationMixinFields0 := integrationMixin[0].Fields()
	_ = integrationMixinFields0
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
	// integrationDescID is the schema descriptor for id field.
	integrationDescID := integrationFields[0].Descriptor()
	// integration.DefaultID holds the default value on creation for the id field.
	integration.DefaultID = integrationDescID.Default.(func() uuid.UUID)
	membershipMixin := schema.Membership{}.Mixin()
	membershipMixinHooks0 := membershipMixin[0].Hooks()
	membership.Hooks[0] = membershipMixinHooks0[0]
	membershipMixinFields0 := membershipMixin[0].Fields()
	_ = membershipMixinFields0
	membershipFields := schema.Membership{}.Fields()
	_ = membershipFields
	// membershipDescCreatedAt is the schema descriptor for created_at field.
	membershipDescCreatedAt := membershipMixinFields0[0].Descriptor()
	// membership.DefaultCreatedAt holds the default value on creation for the created_at field.
	membership.DefaultCreatedAt = membershipDescCreatedAt.Default.(func() time.Time)
	// membershipDescUpdatedAt is the schema descriptor for updated_at field.
	membershipDescUpdatedAt := membershipMixinFields0[1].Descriptor()
	// membership.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	membership.DefaultUpdatedAt = membershipDescUpdatedAt.Default.(func() time.Time)
	// membership.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	membership.UpdateDefaultUpdatedAt = membershipDescUpdatedAt.UpdateDefault.(func() time.Time)
	// membershipDescCurrent is the schema descriptor for current field.
	membershipDescCurrent := membershipFields[1].Descriptor()
	// membership.DefaultCurrent holds the default value on creation for the current field.
	membership.DefaultCurrent = membershipDescCurrent.Default.(bool)
	// membershipDescID is the schema descriptor for id field.
	membershipDescID := membershipFields[0].Descriptor()
	// membership.DefaultID holds the default value on creation for the id field.
	membership.DefaultID = membershipDescID.Default.(func() uuid.UUID)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinHooks0 := organizationMixin[0].Hooks()
	organization.Hooks[0] = organizationMixinHooks0[0]
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
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
	organizationDescName := organizationFields[1].Descriptor()
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
	organizationDescID := organizationFields[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() uuid.UUID)
	sessionMixin := schema.Session{}.Mixin()
	sessionMixinHooks0 := sessionMixin[0].Hooks()
	session.Hooks[0] = sessionMixinHooks0[0]
	sessionMixinFields0 := sessionMixin[0].Fields()
	_ = sessionMixinFields0
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
	sessionDescToken := sessionFields[3].Descriptor()
	// session.DefaultToken holds the default value on creation for the token field.
	session.DefaultToken = sessionDescToken.Default.(func() string)
	// session.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	session.TokenValidator = sessionDescToken.Validators[0].(func(string) error)
	// sessionDescID is the schema descriptor for id field.
	sessionDescID := sessionFields[0].Descriptor()
	// session.DefaultID holds the default value on creation for the id field.
	session.DefaultID = sessionDescID.Default.(func() uuid.UUID)
	tenantFields := schema.Tenant{}.Fields()
	_ = tenantFields
	// tenantDescName is the schema descriptor for name field.
	tenantDescName := tenantFields[1].Descriptor()
	// tenant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tenant.NameValidator = tenantDescName.Validators[0].(func(string) error)
	// tenantDescID is the schema descriptor for id field.
	tenantDescID := tenantFields[0].Descriptor()
	// tenant.DefaultID holds the default value on creation for the id field.
	tenant.DefaultID = tenantDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinHooks0 := userMixin[0].Hooks()
	user.Hooks[0] = userMixinHooks0[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
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
	userDescFirstName := userFields[2].Descriptor()
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
	userDescLastName := userFields[3].Descriptor()
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
	userDescDisplayName := userFields[4].Descriptor()
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
	userDescLocked := userFields[5].Descriptor()
	// user.DefaultLocked holds the default value on creation for the locked field.
	user.DefaultLocked = userDescLocked.Default.(bool)
	// userDescAvatarRemoteURL is the schema descriptor for avatar_remote_url field.
	userDescAvatarRemoteURL := userFields[6].Descriptor()
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
	userDescAvatarLocalFile := userFields[7].Descriptor()
	// user.AvatarLocalFileValidator is a validator for the "avatar_local_file" field. It is called by the builders before save.
	user.AvatarLocalFileValidator = userDescAvatarLocalFile.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[1].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
