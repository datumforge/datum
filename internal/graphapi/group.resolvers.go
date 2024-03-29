package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input generated.CreateGroupInput) (*GroupCreatePayload, error) {
	group, err := withTransactionalMutation(ctx).Group.Create().SetInput(input).Save(ctx)
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
			return nil, newPermissionDeniedError(ActionCreate, "group")
		}

		r.logger.Errorw("failed to create group", "error", err)

		return nil, ErrInternalServerError
	}

	return &GroupCreatePayload{Group: group}, nil
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id string, input generated.UpdateGroupInput) (*GroupUpdatePayload, error) {
	// setup view context
	v := viewer.UserViewer{
		GroupID: id,
	}

	ctx = viewer.NewContext(ctx, v)

	group, err := withTransactionalMutation(ctx).Group.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get group on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "group")
		}

		r.logger.Errorw("failed to get group", "error", err)
		return nil, ErrInternalServerError
	}

	group, err = group.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update group", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "group")
		}

		r.logger.Errorw("failed to update group", "error", err)
		return nil, ErrInternalServerError
	}

	return &GroupUpdatePayload{Group: group}, nil
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id string) (*GroupDeletePayload, error) {
	// setup view context
	v := viewer.UserViewer{
		GroupID: id,
	}

	ctx = viewer.NewContext(ctx, v)

	if err := withTransactionalMutation(ctx).Group.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "group")
		}

		r.logger.Errorw("failed to delete group", "error", err)
		return nil, err
	}

	if err := generated.GroupEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &GroupDeletePayload{DeletedID: id}, nil
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id string) (*generated.Group, error) {
	// setup view context
	v := viewer.UserViewer{
		GroupID: id,
	}

	ctx = viewer.NewContext(ctx, v)

	group, err := withTransactionalMutation(ctx).Group.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get group", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "group")
		}

		return nil, ErrInternalServerError
	}

	return group, nil
}
