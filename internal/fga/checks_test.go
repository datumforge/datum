package fga

import (
	"context"
	"testing"

	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"

	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
)

func Test_CheckTuple(t *testing.T) {
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
			// setup mock client
			c := mock_fga.NewMockSdkClient(t)
			mc := NewMockFGAClient(t, c)

			// mock response for input
			body := ofgaclient.ClientCheckRequest{
				User:     "user:ulid-of-member",
				Relation: tc.relation,
				Object:   tc.object,
			}

			mock_fga.MockeryCheckAny(t, c, tc.expectedRes)

			// do request
			valid, err := mc.CheckTuple(context.Background(), body)

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
