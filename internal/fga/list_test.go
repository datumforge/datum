package fga

import (
	"context"
	"errors"
	"testing"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"

	mock_fga "github.com/datumforge/datum/internal/fga/mockery"
)

func Test_ListContains(t *testing.T) {
	testCases := []struct {
		name        string
		objectID    string
		fgaObjects  []string
		expectedRes bool
	}{
		{
			name:     "happy path, object found",
			objectID: "TbaK4knu9NDoG85DAKob0",
			fgaObjects: []string{
				"organization:TbaK4knu9NDoG85DAKob0",
				"organization:-AV6JyT7-qmedy0WPOjKM",
				"something-else:TbaK4knu9NDoG85DAKob0",
			},
			expectedRes: true,
		},
		{
			name:     "incorrect type but correct id, not found",
			objectID: "TbaK4knu9NDoG85DAKob0",
			fgaObjects: []string{
				"organization:GxSAidJu4LZzjcnHQ-KTV",
				"organization:-AV6JyT7-qmedy0WPOjKM",
				"something-else:TbaK4knu9NDoG85DAKob0",
			},
			expectedRes: false,
		},
		{
			name:     "id not found",
			objectID: "TbaK4knu9NDoG85DAKob0",
			fgaObjects: []string{
				"organization:GxSAidJu4LZzjcnHQ-KTV",
				"organization:-AV6JyT7-qmedy0WPOjKM",
			},
			expectedRes: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entityType := "organization"
			found := ListContains(entityType, tc.fgaObjects, tc.objectID)

			assert.Equal(t, tc.expectedRes, found)
		})
	}
}

func Test_ListObjectsRequest(t *testing.T) {
	objects := []string{"organization:datum"}
	testCases := []struct {
		name        string
		relation    string
		userID      string
		objectType  string
		expectedRes *ofgaclient.ClientListObjectsResponse
		errRes      error
	}{
		{
			name:       "happy path",
			relation:   "can_view",
			userID:     "ulid-of-user",
			objectType: "organization",
			expectedRes: &openfga.ListObjectsResponse{
				Objects: objects,
			},
			errRes: nil,
		},
		{
			name:        "error response",
			relation:    "can_view",
			userID:      "ulid-of-user",
			objectType:  "organization",
			expectedRes: nil,
			errRes:      errors.New("boom"), //nolint:goerr113
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup mock client
			mc := mock_fga.NewMockSdkClient(t)

			c := NewMockFGAClient(t, mc)

			// mock response for input
			body := []string{
				"organization:datum",
			}

			mock_fga.ListOnce(t, mc, body, tc.errRes)

			// do request
			resp, err := c.ListObjectsRequest(
				context.Background(),
				tc.userID,
				tc.objectType,
				tc.relation,
			)

			if tc.errRes != nil {
				assert.Error(t, err)
				assert.Equal(t, err, tc.errRes)
				assert.Equal(t, tc.expectedRes, resp)

				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedRes.GetObjects(), resp.GetObjects())
		})
	}
}
