// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupmembership"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

func EmailVerificationTokenEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func EntitlementEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).GroupMembership.Query().Where((groupmembership.HasGroupWith(group.ID(id)))).Exist(ctx); err == nil && exists {
		if groupmembershipCount, err := FromContext(ctx).GroupMembership.Delete().Where(groupmembership.HasGroupWith(group.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting groupmembership", "count", groupmembershipCount, "err", err)
			return err
		}
	}

	return nil
}

func GroupHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupMembershipEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupSettingEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupSettingHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func IntegrationEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OauthProviderEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OhAuthTooTokenEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OrgMembershipEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OrganizationEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).Group.Query().Where((group.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if groupCount, err := FromContext(ctx).Group.Delete().Where(group.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting group", "count", groupCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).Integration.Query().Where((integration.HasOwnerWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if integrationCount, err := FromContext(ctx).Integration.Delete().Where(integration.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting integration", "count", integrationCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).OrganizationSetting.Query().Where((organizationsetting.HasOrganizationWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if organizationsettingCount, err := FromContext(ctx).OrganizationSetting.Delete().Where(organizationsetting.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting organizationsetting", "count", organizationsettingCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).OrgMembership.Query().Where((orgmembership.HasOrgWith(organization.ID(id)))).Exist(ctx); err == nil && exists {
		if orgmembershipCount, err := FromContext(ctx).OrgMembership.Delete().Where(orgmembership.HasOrgWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting orgmembership", "count", orgmembershipCount, "err", err)
			return err
		}
	}

	return nil
}

func OrganizationHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OrganizationSettingEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func OrganizationSettingHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func PasswordResetTokenEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func PersonalAccessTokenEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func UserEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).PersonalAccessToken.Query().Where((personalaccesstoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if personalaccesstokenCount, err := FromContext(ctx).PersonalAccessToken.Delete().Where(personalaccesstoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting personalaccesstoken", "count", personalaccesstokenCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).UserSetting.Query().Where((usersetting.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if usersettingCount, err := FromContext(ctx).UserSetting.Delete().Where(usersetting.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting usersetting", "count", usersettingCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).EmailVerificationToken.Query().Where((emailverificationtoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if emailverificationtokenCount, err := FromContext(ctx).EmailVerificationToken.Delete().Where(emailverificationtoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting emailverificationtoken", "count", emailverificationtokenCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).PasswordResetToken.Query().Where((passwordresettoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if passwordresettokenCount, err := FromContext(ctx).PasswordResetToken.Delete().Where(passwordresettoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting passwordresettoken", "count", passwordresettokenCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).OrgMembership.Query().Where((orgmembership.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if orgmembershipCount, err := FromContext(ctx).OrgMembership.Delete().Where(orgmembership.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting orgmembership", "count", orgmembershipCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).GroupMembership.Query().Where((groupmembership.HasUserWith(user.ID(id)))).Exist(ctx); err == nil && exists {
		if groupmembershipCount, err := FromContext(ctx).GroupMembership.Delete().Where(groupmembership.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting groupmembership", "count", groupmembershipCount, "err", err)
			return err
		}
	}

	return nil
}

func UserHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func UserSettingEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func UserSettingHistoryEdgeCleanup(ctx context.Context, id string) error {

	return nil
}
