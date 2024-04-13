package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/tfasetting"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
)

// CreateTFASetting is the resolver for the createTFASetting field.
func (r *mutationResolver) CreateTFASetting(ctx context.Context, input generated.CreateTFASettingInput) (*TFASettingCreatePayload, error) {
	// setup view context
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	input.OwnerID = &userID

	settings, err := withTransactionalMutation(ctx).TFASetting.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &TFASettingCreatePayload{TfaSetting: settings}, nil
}

// UpdateTFASetting is the resolver for the updateTFASetting field.
func (r *mutationResolver) UpdateTFASetting(ctx context.Context, input generated.UpdateTFASettingInput) (*TFASettingUpdatePayload, error) {
	// setup view context
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	settings, err := withTransactionalMutation(ctx).TFASetting.Query().Where(tfasetting.OwnerID(userID)).Only(ctx)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied
		}

		r.logger.Errorw("failed to get tfa settings", "error", err)
		return nil, ErrInternalServerError
	}

	settings, err = settings.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			ve := err.(*generated.ValidationError)

			return nil, rout.InvalidField(ve.Name)
		}

		r.logger.Errorw("failed to update tfa settings", "error", err)
		return nil, err
	}

	return &TFASettingUpdatePayload{TfaSetting: settings}, nil
}

// TfaSetting is the resolver for the tfaSettings field.
func (r *queryResolver) TfaSetting(ctx context.Context, id *string) (*generated.TFASetting, error) {
	// setup view context
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var (
		settings *generated.TFASetting
	)

	if id != nil && *id != "" {
		settings, err = withTransactionalMutation(ctx).TFASetting.Get(ctx, *id)
		if err != nil {
			return nil, err
		}
	} else {
		settings, err = withTransactionalMutation(ctx).TFASetting.Query().Where(tfasetting.OwnerID(userID)).Only(ctx)
		if err != nil {
			return nil, err
		}
	}

	return settings, nil
}

// RegenBackupCodes is the resolver for the regenBackupCodes field.
func (r *updateTFASettingInputResolver) RegenBackupCodes(ctx context.Context, obj *generated.UpdateTFASettingInput, data *bool) error {
	return nil
}