package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateOrgSettings is the resolver for the createOrgSettings field.
func (r *createOrganizationInputResolver) CreateOrgSettings(ctx context.Context, obj *generated.CreateOrganizationInput, data *generated.CreateOrganizationSettingInput) error {
	c := withTransactionalMutation(ctx)

	orgSettings, err := c.OrganizationSetting.Create().SetInput(*data).Save(ctx)
	if err != nil {
		return err
	}

	obj.SettingID = &orgSettings.ID

	return nil
}

// AddOrgMembers is the resolver for the addOrgMembers field.
func (r *updateOrganizationInputResolver) AddOrgMembers(ctx context.Context, obj *generated.UpdateOrganizationInput, data []*generated.CreateOrgMembershipInput) error {
	opCtx := graphql.GetOperationContext(ctx)
	orgID, ok := opCtx.Variables["updateOrganizationId"]
	if !ok {
		r.logger.Errorw("unable to get org from context")

		return ErrInternalServerError
	}

	c := withTransactionalMutation(ctx)
	builders := make([]*generated.OrgMembershipCreate, len(data))
	for i := range data {
		input := *data[i]
		input.OrganizationID = orgID.(string)
		builders[i] = c.OrgMembership.Create().SetInput(input)
	}

	_, err := c.OrgMembership.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// UpdateOrgSettings is the resolver for the updateOrgSettings field.
func (r *updateOrganizationInputResolver) UpdateOrgSettings(ctx context.Context, obj *generated.UpdateOrganizationInput, data *generated.UpdateOrganizationSettingInput) error {
	opCtx := graphql.GetOperationContext(ctx)
	orgID, ok := opCtx.Variables["updateOrganizationId"]
	if !ok {
		r.logger.Errorw("unable to get org from context")

		return ErrInternalServerError
	}

	c := withTransactionalMutation(ctx)

	// get setting ID to Update
	settingID := obj.SettingID
	if settingID == nil {
		org, err := c.Organization.Get(ctx, orgID.(string))
		if err != nil {
			return err
		}

		setting, err := org.Setting(ctx)
		if err != nil {
			return err
		}

		settingID = &setting.ID
	}

	_, err := c.OrganizationSetting.UpdateOneID(*settingID).SetInput(*data).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
