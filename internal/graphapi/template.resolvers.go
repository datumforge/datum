package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateTemplate is the resolver for the createTemplate field.
func (r *mutationResolver) CreateTemplate(ctx context.Context, input generated.CreateTemplateInput) (*TemplateCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateTemplate - createTemplate"))
}

// UpdateTemplate is the resolver for the updateTemplate field.
func (r *mutationResolver) UpdateTemplate(ctx context.Context, id string, input generated.UpdateTemplateInput) (*TemplateUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateTemplate - updateTemplate"))
}

// DeleteTemplate is the resolver for the deleteTemplate field.
func (r *mutationResolver) DeleteTemplate(ctx context.Context, id string) (*TemplateDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteTemplate - deleteTemplate"))
}

// Template is the resolver for the template field.
func (r *queryResolver) Template(ctx context.Context, id string) (*generated.Template, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}
