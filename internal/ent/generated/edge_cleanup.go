// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/session"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

func EntitlementEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func GroupSettingEdgeCleanup(ctx context.Context, id string) error {

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

func OrganizationEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).Group.Query().Where((group.HasOwnerWith(organization.ID(id)))).Exist(ctx); err != nil && exists {
		if groupCount, err := FromContext(ctx).Group.Delete().Where(group.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting group", "count", groupCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).Integration.Query().Where((integration.HasOwnerWith(organization.ID(id)))).Exist(ctx); err != nil && exists {
		if integrationCount, err := FromContext(ctx).Integration.Delete().Where(integration.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting integration", "count", integrationCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).OrganizationSetting.Query().Where((organizationsetting.HasOrganizationWith(organization.ID(id)))).Exist(ctx); err != nil && exists {
		if organizationsettingCount, err := FromContext(ctx).OrganizationSetting.Delete().Where(organizationsetting.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting organizationsetting", "count", organizationsettingCount, "err", err)
			return err
		}
	}

	return nil
}

func OrganizationSettingEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func PermissionEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func PersonalAccessTokenEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func RoleEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func RolePermissionEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func SessionEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func UserEdgeCleanup(ctx context.Context, id string) error {

	if exists, err := FromContext(ctx).Session.Query().Where((session.HasOwnerWith(user.ID(id)))).Exist(ctx); err != nil && exists {
		if sessionCount, err := FromContext(ctx).Session.Delete().Where(session.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting session", "count", sessionCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).PersonalAccessToken.Query().Where((personalaccesstoken.HasOwnerWith(user.ID(id)))).Exist(ctx); err != nil && exists {
		if personalaccesstokenCount, err := FromContext(ctx).PersonalAccessToken.Delete().Where(personalaccesstoken.HasOwnerWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting personalaccesstoken", "count", personalaccesstokenCount, "err", err)
			return err
		}
	}

	if exists, err := FromContext(ctx).UserSetting.Query().Where((usersetting.HasUserWith(user.ID(id)))).Exist(ctx); err != nil && exists {
		if usersettingCount, err := FromContext(ctx).UserSetting.Delete().Where(usersetting.HasUserWith(user.ID(id))).Exec(ctx); err != nil {
			FromContext(ctx).Logger.Debugw("deleting usersetting", "count", usersettingCount, "err", err)
			return err
		}
	}

	return nil
}

func UserRoleEdgeCleanup(ctx context.Context, id string) error {

	return nil
}

func UserSettingEdgeCleanup(ctx context.Context, id string) error {

	return nil
}
