package graphapi_test

import (
	"context"
	"testing"

	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/ulids"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestQueryFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	feature := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: feature.ID,
		},
		{
			name:     "not found",
			queryID:  "notfound",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errorMsg == "" {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := suite.client.datum.GetFeatureByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.queryID, resp.Feature.ID)
			assert.NotEmpty(t, resp.Feature.Name)
			assert.NotEmpty(t, resp.Feature.Description)
			assert.NotEmpty(t, resp.Feature.DisplayName)
		})
	}
}

func (suite *GraphTestSuite) TestQueryFeatures() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	otherUser := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	otherCtx, err := userContextWithID(otherUser.ID)
	require.NoError(t, err)

	testCases := []struct {
		name            string
		ctx             context.Context
		expectedResults int
	}{
		{
			name:            "happy path",
			ctx:             reqCtx,
			expectedResults: 2,
		},
		{
			name:            "another user, no features should be returned",
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllFeatures(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.Features.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		request     datumclient.CreateFeatureInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateFeatureInput{
				Name: "test-feature",
			},
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateFeatureInput{
				Name:        "mitb",
				DisplayName: lo.ToPtr("Matt is the Best"),
				Enabled:     lo.ToPtr(true),
				Description: lo.ToPtr("Matt is the best feature, hands down!"),
			},
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateFeatureInput{
				Name: "test-feature",
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on feature",
		},
		{
			name: "missing required field",
			request: datumclient.CreateFeatureInput{
				DisplayName: lo.ToPtr("Matt is the Best"),
			},
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.CreateFeature(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.Name, resp.CreateFeature.Feature.Name)

			if tc.request.Enabled == nil {
				assert.False(t, resp.CreateFeature.Feature.Enabled)
			} else {
				assert.Equal(t, *tc.request.Enabled, resp.CreateFeature.Feature.Enabled)
			}

			if tc.request.Description == nil {
				assert.Nil(t, resp.CreateFeature.Feature.Description)
			} else {
				assert.Equal(t, *tc.request.Description, *resp.CreateFeature.Feature.Description)
			}

			// Display Name is set to the Name if not provided
			if tc.request.DisplayName == nil {
				assert.Equal(t, tc.request.Name, *resp.CreateFeature.Feature.DisplayName)
			} else {
				assert.Equal(t, *tc.request.DisplayName, *resp.CreateFeature.Feature.DisplayName)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	feature := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateFeatureInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update display name",
			request: datumclient.UpdateFeatureInput{
				DisplayName: lo.ToPtr("test-feature"),
			},
			allowed: true,
		},
		{
			name: "enable feature",
			request: datumclient.UpdateFeatureInput{
				Enabled: lo.ToPtr(true),
			},
			allowed: true,
		},
		{
			name: "update description",
			request: datumclient.UpdateFeatureInput{
				Description: lo.ToPtr("To infinity and beyond!"),
			},
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateFeatureInput{
				Enabled: lo.ToPtr(false),
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on feature",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateFeature(reqCtx, feature.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.Description != nil {
				assert.Equal(t, *tc.request.Description, *resp.UpdateFeature.Feature.Description)
			}

			if tc.request.DisplayName != nil {
				assert.Equal(t, *tc.request.DisplayName, *resp.UpdateFeature.Feature.DisplayName)
			}

			if tc.request.Enabled != nil {
				assert.Equal(t, *tc.request.Enabled, resp.UpdateFeature.Feature.Enabled)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	feature := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  feature.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on feature",
		},
		{
			name:        "happy path, delete feature",
			idToDelete:  feature.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "feature already deleted, not found",
			idToDelete:  feature.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "feature not found",
		},
		{
			name:        "unknown feature, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "feature not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if feature exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteFeature(reqCtx, feature.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, feature.ID, resp.DeleteFeature.DeletedID)
		})
	}
}
