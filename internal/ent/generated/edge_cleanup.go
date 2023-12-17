// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
)

func AccessTokenEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func EntitlementEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func GroupEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func GroupSettingEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func IntegrationEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func OauthProviderEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func OhAuthTooTokenEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func OrganizationEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	if groupCount, err := r.Group.Delete().Where(group.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
		r.Logger.Debugw("deleting group", "count", groupCount, "err", err)
		return err
	}

	if integrationCount, err := r.Integration.Delete().Where(integration.HasOwnerWith(organization.ID(id))).Exec(ctx); err != nil {
		r.Logger.Debugw("deleting integration", "count", integrationCount, "err", err)
		return err
	}

	if organizationsettingCount, err := r.OrganizationSetting.Delete().Where(organizationsetting.HasOrganizationWith(organization.ID(id))).Exec(ctx); err != nil {
		r.Logger.Debugw("deleting organizationsetting", "count", organizationsettingCount, "err", err)
		return err
	}

	return nil
}

func OrganizationSettingEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func PersonalAccessTokenEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func RefreshTokenEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func SessionEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func UserEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}

func UserSettingEdgeCleanup(ctx context.Context, r *Client, id string) error {
	// TODO: pass in transaction so that all upstream
	// deletes can be rolled back if one fails

	return nil
}
