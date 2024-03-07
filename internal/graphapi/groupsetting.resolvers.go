package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/rout"
)

// CreateGroupSetting is the resolver for the createGroupSetting field.
func (r *mutationResolver) CreateGroupSetting(ctx context.Context, input generated.CreateGroupSettingInput) (*GroupSettingCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateGroupSetting - createGroupSetting"))
}

// UpdateGroupSetting is the resolver for the updateGroupSetting field.
func (r *mutationResolver) UpdateGroupSetting(ctx context.Context, id string, input generated.UpdateGroupSettingInput) (*GroupSettingUpdatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	groupSetting, err := withTransactionalMutation(ctx).GroupSetting.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied
		}

		r.logger.Errorw("failed to get user setting", "error", err)
		return nil, ErrInternalServerError
	}

	groupSetting, err = groupSetting.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			ve := err.(*generated.ValidationError)

			return nil, rout.InvalidField(ve.Name)
		}

		r.logger.Errorw("failed to update user setting", "error", err)
		return nil, err
	}

	return &GroupSettingUpdatePayload{GroupSetting: groupSetting}, nil
}

// DeleteGroupSetting is the resolver for the deleteGroupSetting field.
func (r *mutationResolver) DeleteGroupSetting(ctx context.Context, id string) (*GroupSettingDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteGroupSetting - deleteGroupSetting"))
}

// GroupSetting is the resolver for the groupSetting field.
func (r *queryResolver) GroupSetting(ctx context.Context, id string) (*generated.GroupSetting, error) {
	// setup view context
	v := viewer.UserViewer{
		GroupID: id,
	}

	ctx = viewer.NewContext(ctx, v)

	group, err := withTransactionalMutation(ctx).GroupSetting.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get group settings", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "organization")
		}

		return nil, ErrInternalServerError
	}

	return group, nil
}
