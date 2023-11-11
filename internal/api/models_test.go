package api_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/datumforge/datum/internal/ent/generated"
)

type OrganizationBuilder struct {
	Name        string
	DisplayName string
	Description *string
}

func (o *OrganizationBuilder) MustNew(ctx context.Context) *generated.Organization {
	if o.Name == "" {
		o.Name = gofakeit.AppName()
	}

	if o.DisplayName == "" {
		o.Name = gofakeit.LetterN(40)
	}

	if o.Description == nil {
		desc := gofakeit.HipsterSentence(10)
		o.Description = &desc
	}

	return EntClient.Organization.Create().SetName(o.Name).SetDescription(*o.Description).SaveX(ctx)
}
