package graphapi

import (
	"context"
	"errors"
	"regexp"

	"github.com/stoewer/go-strcase"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func (r *mutationResolver) createOrg(ctx context.Context, input generated.CreateOrganizationInput) (*OrganizationCreatePayload, error) {
	// if this is empty generate a default org setting schema
	if input.SettingID == nil {
		// sets up default org settings using schema defaults
		orgSettingID, err := r.defaultOrganizationSettings(ctx)
		if err != nil {
			return nil, err
		}

		// add the org setting ID to the input
		input.SettingID = &orgSettingID
	}

	// default the display name to the name if not provided
	if input.DisplayName == nil {
		displayName := strcase.LowerCamelCase(input.Name)
		displayName = nonAlphanumericRegex.ReplaceAllString(displayName, "")

		input.DisplayName = &displayName
	}

	org, err := r.client.Organization.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			validationError := err.(*generated.ValidationError)

			r.logger.Debugw("validation error", "field", validationError.Name, "error", validationError.Error())

			return nil, validationError
		}

		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			r.logger.Debugw("constraint error", "error", constraintError.Error())

			return nil, constraintError
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionCreate, "organization")
		}

		r.logger.Errorw("failed to create organization", "error", err)

		return nil, ErrInternalServerError
	}

	return &OrganizationCreatePayload{Organization: org}, nil
}

// defaultOrganizationSettings creates the default organizations settings for a new org
func (r *mutationResolver) defaultOrganizationSettings(ctx context.Context) (string, error) { //nolint:unused
	input := generated.CreateOrganizationSettingInput{}

	organizationSetting, err := r.client.OrganizationSetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return "", err
	}

	return organizationSetting.ID, nil
}
