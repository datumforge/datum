// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// The APITokenFunc type is an adapter to allow the use of ordinary
// function as APIToken mutator.
type APITokenFunc func(context.Context, *generated.APITokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f APITokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.APITokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.APITokenMutation", m)
}

// The ContactFunc type is an adapter to allow the use of ordinary
// function as Contact mutator.
type ContactFunc func(context.Context, *generated.ContactMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ContactFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ContactMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ContactMutation", m)
}

// The ContactHistoryFunc type is an adapter to allow the use of ordinary
// function as ContactHistory mutator.
type ContactHistoryFunc func(context.Context, *generated.ContactHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f ContactHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.ContactHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.ContactHistoryMutation", m)
}

// The DocumentDataFunc type is an adapter to allow the use of ordinary
// function as DocumentData mutator.
type DocumentDataFunc func(context.Context, *generated.DocumentDataMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f DocumentDataFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.DocumentDataMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.DocumentDataMutation", m)
}

// The DocumentDataHistoryFunc type is an adapter to allow the use of ordinary
// function as DocumentDataHistory mutator.
type DocumentDataHistoryFunc func(context.Context, *generated.DocumentDataHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f DocumentDataHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.DocumentDataHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.DocumentDataHistoryMutation", m)
}

// The EmailVerificationTokenFunc type is an adapter to allow the use of ordinary
// function as EmailVerificationToken mutator.
type EmailVerificationTokenFunc func(context.Context, *generated.EmailVerificationTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EmailVerificationTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EmailVerificationTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EmailVerificationTokenMutation", m)
}

// The EntitlementFunc type is an adapter to allow the use of ordinary
// function as Entitlement mutator.
type EntitlementFunc func(context.Context, *generated.EntitlementMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementMutation", m)
}

// The EntitlementHistoryFunc type is an adapter to allow the use of ordinary
// function as EntitlementHistory mutator.
type EntitlementHistoryFunc func(context.Context, *generated.EntitlementHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementHistoryMutation", m)
}

// The EntitlementPlanFunc type is an adapter to allow the use of ordinary
// function as EntitlementPlan mutator.
type EntitlementPlanFunc func(context.Context, *generated.EntitlementPlanMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementPlanFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementPlanMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementPlanMutation", m)
}

// The EntitlementPlanFeatureFunc type is an adapter to allow the use of ordinary
// function as EntitlementPlanFeature mutator.
type EntitlementPlanFeatureFunc func(context.Context, *generated.EntitlementPlanFeatureMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementPlanFeatureFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementPlanFeatureMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementPlanFeatureMutation", m)
}

// The EntitlementPlanFeatureHistoryFunc type is an adapter to allow the use of ordinary
// function as EntitlementPlanFeatureHistory mutator.
type EntitlementPlanFeatureHistoryFunc func(context.Context, *generated.EntitlementPlanFeatureHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementPlanFeatureHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementPlanFeatureHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementPlanFeatureHistoryMutation", m)
}

// The EntitlementPlanHistoryFunc type is an adapter to allow the use of ordinary
// function as EntitlementPlanHistory mutator.
type EntitlementPlanHistoryFunc func(context.Context, *generated.EntitlementPlanHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntitlementPlanHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntitlementPlanHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntitlementPlanHistoryMutation", m)
}

// The EntityFunc type is an adapter to allow the use of ordinary
// function as Entity mutator.
type EntityFunc func(context.Context, *generated.EntityMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntityFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntityMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntityMutation", m)
}

// The EntityHistoryFunc type is an adapter to allow the use of ordinary
// function as EntityHistory mutator.
type EntityHistoryFunc func(context.Context, *generated.EntityHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntityHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntityHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntityHistoryMutation", m)
}

// The EntityTypeFunc type is an adapter to allow the use of ordinary
// function as EntityType mutator.
type EntityTypeFunc func(context.Context, *generated.EntityTypeMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntityTypeFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntityTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntityTypeMutation", m)
}

// The EntityTypeHistoryFunc type is an adapter to allow the use of ordinary
// function as EntityTypeHistory mutator.
type EntityTypeHistoryFunc func(context.Context, *generated.EntityTypeHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EntityTypeHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EntityTypeHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EntityTypeHistoryMutation", m)
}

// The EventFunc type is an adapter to allow the use of ordinary
// function as Event mutator.
type EventFunc func(context.Context, *generated.EventMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EventFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EventMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EventMutation", m)
}

// The EventHistoryFunc type is an adapter to allow the use of ordinary
// function as EventHistory mutator.
type EventHistoryFunc func(context.Context, *generated.EventHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f EventHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.EventHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.EventHistoryMutation", m)
}

// The FeatureFunc type is an adapter to allow the use of ordinary
// function as Feature mutator.
type FeatureFunc func(context.Context, *generated.FeatureMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f FeatureFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.FeatureMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.FeatureMutation", m)
}

// The FeatureHistoryFunc type is an adapter to allow the use of ordinary
// function as FeatureHistory mutator.
type FeatureHistoryFunc func(context.Context, *generated.FeatureHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f FeatureHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.FeatureHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.FeatureHistoryMutation", m)
}

// The FileFunc type is an adapter to allow the use of ordinary
// function as File mutator.
type FileFunc func(context.Context, *generated.FileMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f FileFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.FileMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.FileMutation", m)
}

// The FileHistoryFunc type is an adapter to allow the use of ordinary
// function as FileHistory mutator.
type FileHistoryFunc func(context.Context, *generated.FileHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f FileHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.FileHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.FileHistoryMutation", m)
}

// The GroupFunc type is an adapter to allow the use of ordinary
// function as Group mutator.
type GroupFunc func(context.Context, *generated.GroupMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupMutation", m)
}

// The GroupHistoryFunc type is an adapter to allow the use of ordinary
// function as GroupHistory mutator.
type GroupHistoryFunc func(context.Context, *generated.GroupHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupHistoryMutation", m)
}

// The GroupMembershipFunc type is an adapter to allow the use of ordinary
// function as GroupMembership mutator.
type GroupMembershipFunc func(context.Context, *generated.GroupMembershipMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupMembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupMembershipMutation", m)
}

// The GroupMembershipHistoryFunc type is an adapter to allow the use of ordinary
// function as GroupMembershipHistory mutator.
type GroupMembershipHistoryFunc func(context.Context, *generated.GroupMembershipHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupMembershipHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupMembershipHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupMembershipHistoryMutation", m)
}

// The GroupSettingFunc type is an adapter to allow the use of ordinary
// function as GroupSetting mutator.
type GroupSettingFunc func(context.Context, *generated.GroupSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupSettingMutation", m)
}

// The GroupSettingHistoryFunc type is an adapter to allow the use of ordinary
// function as GroupSettingHistory mutator.
type GroupSettingHistoryFunc func(context.Context, *generated.GroupSettingHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f GroupSettingHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.GroupSettingHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.GroupSettingHistoryMutation", m)
}

// The HushFunc type is an adapter to allow the use of ordinary
// function as Hush mutator.
type HushFunc func(context.Context, *generated.HushMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f HushFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.HushMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.HushMutation", m)
}

// The HushHistoryFunc type is an adapter to allow the use of ordinary
// function as HushHistory mutator.
type HushHistoryFunc func(context.Context, *generated.HushHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f HushHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.HushHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.HushHistoryMutation", m)
}

// The IntegrationFunc type is an adapter to allow the use of ordinary
// function as Integration mutator.
type IntegrationFunc func(context.Context, *generated.IntegrationMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f IntegrationFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.IntegrationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.IntegrationMutation", m)
}

// The IntegrationHistoryFunc type is an adapter to allow the use of ordinary
// function as IntegrationHistory mutator.
type IntegrationHistoryFunc func(context.Context, *generated.IntegrationHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f IntegrationHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.IntegrationHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.IntegrationHistoryMutation", m)
}

// The InviteFunc type is an adapter to allow the use of ordinary
// function as Invite mutator.
type InviteFunc func(context.Context, *generated.InviteMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f InviteFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.InviteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.InviteMutation", m)
}

// The NoteFunc type is an adapter to allow the use of ordinary
// function as Note mutator.
type NoteFunc func(context.Context, *generated.NoteMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f NoteFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.NoteMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.NoteMutation", m)
}

// The NoteHistoryFunc type is an adapter to allow the use of ordinary
// function as NoteHistory mutator.
type NoteHistoryFunc func(context.Context, *generated.NoteHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f NoteHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.NoteHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.NoteHistoryMutation", m)
}

// The OauthProviderFunc type is an adapter to allow the use of ordinary
// function as OauthProvider mutator.
type OauthProviderFunc func(context.Context, *generated.OauthProviderMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OauthProviderFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OauthProviderMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OauthProviderMutation", m)
}

// The OauthProviderHistoryFunc type is an adapter to allow the use of ordinary
// function as OauthProviderHistory mutator.
type OauthProviderHistoryFunc func(context.Context, *generated.OauthProviderHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OauthProviderHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OauthProviderHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OauthProviderHistoryMutation", m)
}

// The OhAuthTooTokenFunc type is an adapter to allow the use of ordinary
// function as OhAuthTooToken mutator.
type OhAuthTooTokenFunc func(context.Context, *generated.OhAuthTooTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OhAuthTooTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OhAuthTooTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OhAuthTooTokenMutation", m)
}

// The OrgMembershipFunc type is an adapter to allow the use of ordinary
// function as OrgMembership mutator.
type OrgMembershipFunc func(context.Context, *generated.OrgMembershipMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrgMembershipFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrgMembershipMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrgMembershipMutation", m)
}

// The OrgMembershipHistoryFunc type is an adapter to allow the use of ordinary
// function as OrgMembershipHistory mutator.
type OrgMembershipHistoryFunc func(context.Context, *generated.OrgMembershipHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrgMembershipHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrgMembershipHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrgMembershipHistoryMutation", m)
}

// The OrganizationFunc type is an adapter to allow the use of ordinary
// function as Organization mutator.
type OrganizationFunc func(context.Context, *generated.OrganizationMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationMutation", m)
}

// The OrganizationHistoryFunc type is an adapter to allow the use of ordinary
// function as OrganizationHistory mutator.
type OrganizationHistoryFunc func(context.Context, *generated.OrganizationHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationHistoryMutation", m)
}

// The OrganizationSettingFunc type is an adapter to allow the use of ordinary
// function as OrganizationSetting mutator.
type OrganizationSettingFunc func(context.Context, *generated.OrganizationSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationSettingMutation", m)
}

// The OrganizationSettingHistoryFunc type is an adapter to allow the use of ordinary
// function as OrganizationSettingHistory mutator.
type OrganizationSettingHistoryFunc func(context.Context, *generated.OrganizationSettingHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f OrganizationSettingHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.OrganizationSettingHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.OrganizationSettingHistoryMutation", m)
}

// The PasswordResetTokenFunc type is an adapter to allow the use of ordinary
// function as PasswordResetToken mutator.
type PasswordResetTokenFunc func(context.Context, *generated.PasswordResetTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f PasswordResetTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.PasswordResetTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.PasswordResetTokenMutation", m)
}

// The PersonalAccessTokenFunc type is an adapter to allow the use of ordinary
// function as PersonalAccessToken mutator.
type PersonalAccessTokenFunc func(context.Context, *generated.PersonalAccessTokenMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f PersonalAccessTokenFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.PersonalAccessTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.PersonalAccessTokenMutation", m)
}

// The SubscriberFunc type is an adapter to allow the use of ordinary
// function as Subscriber mutator.
type SubscriberFunc func(context.Context, *generated.SubscriberMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f SubscriberFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.SubscriberMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.SubscriberMutation", m)
}

// The TFASettingFunc type is an adapter to allow the use of ordinary
// function as TFASetting mutator.
type TFASettingFunc func(context.Context, *generated.TFASettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f TFASettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.TFASettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.TFASettingMutation", m)
}

// The TemplateFunc type is an adapter to allow the use of ordinary
// function as Template mutator.
type TemplateFunc func(context.Context, *generated.TemplateMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f TemplateFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.TemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.TemplateMutation", m)
}

// The TemplateHistoryFunc type is an adapter to allow the use of ordinary
// function as TemplateHistory mutator.
type TemplateHistoryFunc func(context.Context, *generated.TemplateHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f TemplateHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.TemplateHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.TemplateHistoryMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *generated.UserMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserMutation", m)
}

// The UserHistoryFunc type is an adapter to allow the use of ordinary
// function as UserHistory mutator.
type UserHistoryFunc func(context.Context, *generated.UserHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserHistoryMutation", m)
}

// The UserSettingFunc type is an adapter to allow the use of ordinary
// function as UserSetting mutator.
type UserSettingFunc func(context.Context, *generated.UserSettingMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserSettingFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserSettingMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserSettingMutation", m)
}

// The UserSettingHistoryFunc type is an adapter to allow the use of ordinary
// function as UserSettingHistory mutator.
type UserSettingHistoryFunc func(context.Context, *generated.UserSettingHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f UserSettingHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.UserSettingHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.UserSettingHistoryMutation", m)
}

// The WebauthnFunc type is an adapter to allow the use of ordinary
// function as Webauthn mutator.
type WebauthnFunc func(context.Context, *generated.WebauthnMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f WebauthnFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.WebauthnMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.WebauthnMutation", m)
}

// The WebhookFunc type is an adapter to allow the use of ordinary
// function as Webhook mutator.
type WebhookFunc func(context.Context, *generated.WebhookMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f WebhookFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.WebhookMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.WebhookMutation", m)
}

// The WebhookHistoryFunc type is an adapter to allow the use of ordinary
// function as WebhookHistory mutator.
type WebhookHistoryFunc func(context.Context, *generated.WebhookHistoryMutation) (generated.Value, error)

// Mutate calls f(ctx, m).
func (f WebhookHistoryFunc) Mutate(ctx context.Context, m generated.Mutation) (generated.Value, error) {
	if mv, ok := m.(*generated.WebhookHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *generated.WebhookHistoryMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, generated.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m generated.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op generated.Op) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m generated.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk generated.Hook, cond Condition) generated.Hook {
	return func(next generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(ctx context.Context, m generated.Mutation) (generated.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, generated.Delete|generated.Create)
func On(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, generated.Update|generated.UpdateOne)
func Unless(hk generated.Hook, op generated.Op) generated.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) generated.Hook {
	return func(generated.Mutator) generated.Mutator {
		return generated.MutateFunc(func(context.Context, generated.Mutation) (generated.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []generated.Hook {
//		return []generated.Hook{
//			Reject(generated.Delete|generated.Update),
//		}
//	}
func Reject(op generated.Op) generated.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []generated.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...generated.Hook) Chain {
	return Chain{append([]generated.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() generated.Hook {
	return func(mutator generated.Mutator) generated.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...generated.Hook) Chain {
	newHooks := make([]generated.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
