package api_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/mock/gomock"

	"github.com/datumforge/datum/internal/ent/generated"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
)

type OrganizationBuilder struct {
	Name        string
	DisplayName string
	Description *string
}

func (o *OrganizationBuilder) MustNew(ctx context.Context, mockCtrl *gomock.Controller, mc *mock_client.MockSdkClient) *generated.Organization {
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

	mockWriteTuplesAny(mockCtrl, mc, ctx, nil)

	return EntClient.Organization.Create().SetName(o.Name).SetDescription(*o.Description).SaveX(ctx)
}
