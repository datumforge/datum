package fga

import (
	"context"
	"errors"
	"testing"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mock_client "github.com/datumforge/datum/internal/fga/mocks"
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
			// setup mock controller
			mockCtrl := gomock.NewController(t)
			c := mock_client.NewMockSdkClient(mockCtrl)

			fc, err := NewTestFGAClient(t, mockCtrl, c)
			if err != nil {
				t.Fatal()
			}

			// mock response for input
			body := []string{
				"organization:datum",
			}

			mockListAny(mockCtrl, c, context.Background(), body, tc.errRes)

			// do request
			resp, err := fc.ListObjectsRequest(
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

func mockListAny(mockCtrl *gomock.Controller, c *mock_client.MockSdkClient, ctx context.Context, allowedObjects []string, err error) {
	mockExecute := mock_client.NewMockSdkClientListObjectsRequestInterface(mockCtrl)

	resp := ofgaclient.ClientListObjectsResponse{
		Objects: allowedObjects,
	}

	if err == nil {
		mockExecute.EXPECT().Execute().Return(&resp, nil)
	} else {
		mockExecute.EXPECT().Execute().Return(nil, err)
	}

	mockBody := mock_client.NewMockSdkClientListObjectsRequestInterface(mockCtrl)

	mockBody.EXPECT().Body(gomock.Any()).Return(mockExecute)

	c.EXPECT().ListObjects(gomock.Any()).Return(mockBody)
}
