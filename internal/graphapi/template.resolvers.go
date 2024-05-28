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
)

// CreateTemplate is the resolver for the createTemplate field.
func (r *mutationResolver) CreateTemplate(ctx context.Context, input generated.CreateTemplateInput) (*TemplateCreatePayload, error) {
	t, err := withTransactionalMutation(ctx).Template.Create().SetInput(input).Save(ctx)
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
			return nil, newPermissionDeniedError(ActionCreate, "template")
		}

		r.logger.Errorw("failed to create template", "error", err)

		return nil, err
	}

	return &TemplateCreatePayload{Template: t}, nil
}

// CreateBulkTemplate is the resolver for the createBulkTemplate field.
func (r *mutationResolver) CreateBulkTemplate(ctx context.Context, input []*generated.CreateTemplateInput) (*TemplateBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkTemplate - createBulkTemplate"))
}

// UpdateTemplate is the resolver for the updateTemplate field.
func (r *mutationResolver) UpdateTemplate(ctx context.Context, id string, input generated.UpdateTemplateInput) (*TemplateUpdatePayload, error) {
	template, err := withTransactionalMutation(ctx).Template.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get template on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "template")
		}

		r.logger.Errorw("failed to get template", "error", err)
		return nil, ErrInternalServerError
	}

	template, err = template.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update template", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "template")
		}

		r.logger.Errorw("failed to update template", "error", err)
		return nil, ErrInternalServerError
	}

	return &TemplateUpdatePayload{Template: template}, nil
}

// DeleteTemplate is the resolver for the deleteTemplate field.
func (r *mutationResolver) DeleteTemplate(ctx context.Context, id string) (*TemplateDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Template.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "template")
		}

		r.logger.Errorw("failed to delete template", "error", err)
		return nil, err
	}

	if err := generated.TemplateEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &TemplateDeletePayload{DeletedID: id}, nil
}

// Template is the resolver for the template field.
func (r *queryResolver) Template(ctx context.Context, id string) (*generated.Template, error) {
	template, err := withTransactionalMutation(ctx).Template.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get template", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "template")
		}

		return nil, ErrInternalServerError
	}

	return template, nil
}
