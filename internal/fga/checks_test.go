package fga

import (
	"context"
	"testing"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/datumforge/datum/internal/echox"
	mock_client "github.com/datumforge/datum/internal/fga/mocks"
)

func Test_CheckDirectUser(t *testing.T) {
	ec, err := echox.NewTestContextWithValidUser("nano-id-of-member")
	if err != nil {
		t.Fatal()
	}

	echoContext := *ec

	ctx := context.WithValue(echoContext.Request().Context(), echox.EchoContextKey, echoContext)

	echoContext.SetRequest(echoContext.Request().WithContext(ctx))

	testCases := []struct {
		name        string
		relation    string
		object      string
		expectedRes bool
		errRes      string
	}{
		{
			name:        "happy path, valid tuple",
			relation:    "member",
			object:      "organization:datum",
			expectedRes: true,
			errRes:      "",
		},
		{
			name:        "tuple does not exist",
			relation:    "member",
			object:      "organization:cat-friends",
			expectedRes: false,
			errRes:      "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup mock controller
			mockCtrl := gomock.NewController(t)
			c := mock_client.NewMockSdkClient(mockCtrl)

			fc, err := newTestFGAClient(t, mockCtrl, c)
			if err != nil {
				t.Fatal()
			}

			// mock response for input
			body := ofgaclient.ClientCheckRequest{
				User:     "user:nano-id-of-member",
				Relation: tc.relation,
				Object:   tc.object,
			}

			mockCheck(mockCtrl, c, ctx, body, tc.expectedRes)

			// do request
			valid, err := fc.CheckDirectUser(ctx, tc.relation, tc.object)

			if tc.errRes != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.errRes)
				assert.Equal(t, tc.expectedRes, valid)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRes, valid)
		})
	}
}

func mockCheck(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, body ofgaclient.ClientCheckRequest, allowed bool) {
	mockExecute := mock_client.NewMockSdkClientCheckRequestInterface(mockCtrl)

	resp := ofgaclient.ClientCheckResponse{
		CheckResponse: openfga.CheckResponse{
			Allowed: openfga.PtrBool(allowed),
		},
	}

	mockExecute.EXPECT().Execute().Return(&resp, nil)

	mockBody := mock_client.NewMockSdkClientCheckRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(body).Return(mockExecute)

	c.EXPECT().Check(ctx).Return(mockBody)
}
