package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

// CreateFile is the resolver for the createFile field
func (r *mutationResolver) CreateFile(ctx context.Context, input generated.CreateFileInput) (*FileCreatePayload, error) {
	t, err := withTransactionalMutation(ctx).File.Create().SetInput(input).Save(ctx)
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
			return nil, newPermissionDeniedError(ActionCreate, "file")
		}

		r.logger.Errorw("failed to create file", "error", err)

		return nil, err
	}

	return &FileCreatePayload{File: t}, nil
}

// UpdateFile is the resolver for the updateFile field
func (r *mutationResolver) UpdateFile(ctx context.Context, id string, input generated.UpdateFileInput) (*FileUpdatePayload, error) {
	file, err := withTransactionalMutation(ctx).File.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get file on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "file")
		}

		r.logger.Errorw("failed to get file", "error", err)
		return nil, ErrInternalServerError
	}

	file, err = file.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update file", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "file")
		}

		r.logger.Errorw("failed to update file", "error", err)
		return nil, ErrInternalServerError
	}

	return &FileUpdatePayload{File: file}, nil
}

// DeleteFile is the resolver for the deleteFile field
func (r *mutationResolver) DeleteFile(ctx context.Context, id string) (*FileDeletePayload, error) {
	if err := withTransactionalMutation(ctx).File.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "file")
		}

		r.logger.Errorw("failed to delete file", "error", err)
		return nil, err
	}

	if err := generated.FileEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &FileDeletePayload{DeletedID: id}, nil
}

// File is the resolver for the file field
func (r *queryResolver) File(ctx context.Context, id string) (*generated.File, error) {
	file, err := withTransactionalMutation(ctx).File.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get file", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "file")
		}

		return nil, ErrInternalServerError
	}

	return file, nil
}
