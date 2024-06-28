package graphapi_test

import (
	"context"
	"testing"

	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/ulids"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestQueryEntitlementPlanFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	planFeature := (&EntitlementPlanFeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: planFeature.ID,
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

			resp, err := suite.client.datum.GetEntitlementPlanFeatureByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.queryID, resp.EntitlementPlanFeature.ID)
			require.NotEmpty(t, resp.EntitlementPlanFeature.GetFeature())
			assert.Equal(t, planFeature.FeatureID, resp.EntitlementPlanFeature.Feature.ID)
			require.NotEmpty(t, resp.EntitlementPlanFeature.GetPlan())
			assert.Equal(t, planFeature.PlanID, resp.EntitlementPlanFeature.Plan.ID)
			require.NotEmpty(t, resp.EntitlementPlanFeature.GetMetadata())
		})
	}
}

func (suite *GraphTestSuite) TestQueryEntitlementPlanFeatures() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&EntitlementPlanFeatureBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&EntitlementPlanFeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			name:            "another user, no planFeatures should be returned",
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllEntitlementPlanFeatures(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.EntitlementPlanFeatures.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateEntitlementPlanFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// setup for creation of planFeature
	plan := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)
	feature1 := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)
	feature2 := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)
	feature3 := (&FeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.CreateEntitlementPlanFeatureInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				PlanID:    plan.ID,
				FeatureID: feature1.ID,
			},
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				PlanID:    plan.ID,
				FeatureID: feature2.ID,
				Metadata: map[string]interface{}{
					"limit_type": "days",
					"limit":      "30",
				},
			},
			allowed: true,
		},
		{
			name: "already exists",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				PlanID:    plan.ID,
				FeatureID: feature2.ID,
				Metadata: map[string]interface{}{
					"limit_type": "days",
					"limit":      "30",
				},
			},
			allowed:     true,
			expectedErr: "entitlementplanfeature already exists",
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				PlanID:    plan.ID,
				FeatureID: feature3.ID,
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitlementplanfeature",
		},
		{
			name: "missing required field, feature",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				PlanID: plan.ID,
			},
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
		{
			name: "missing required field, plan",
			request: datumclient.CreateEntitlementPlanFeatureInput{
				FeatureID: feature1.ID,
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

			resp, err := suite.client.datum.CreateEntitlementPlanFeature(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.PlanID, resp.CreateEntitlementPlanFeature.EntitlementPlanFeature.Plan.GetID())
			assert.Equal(t, tc.request.FeatureID, resp.CreateEntitlementPlanFeature.EntitlementPlanFeature.Feature.GetID())

			if tc.request.Metadata != nil {
				assert.Equal(t, tc.request.Metadata, resp.CreateEntitlementPlanFeature.EntitlementPlanFeature.Metadata)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateEntitlementPlanFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	planFeature := (&EntitlementPlanFeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateEntitlementPlanFeatureInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update metadata",
			request: datumclient.UpdateEntitlementPlanFeatureInput{
				Metadata: map[string]interface{}{
					"limit_type": "days",
					"limit":      "15",
				},
			},
			allowed: true,
		},

		{
			name: "not allowed to update",
			request: datumclient.UpdateEntitlementPlanFeatureInput{
				Metadata: map[string]interface{}{
					"limit_type": "days",
					"limit":      "65",
				}},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitlementplanfeature",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateEntitlementPlanFeature(reqCtx, planFeature.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.request.Metadata, resp.UpdateEntitlementPlanFeature.EntitlementPlanFeature.GetMetadata())
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteEntitlementPlanFeature() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	planFeature := (&EntitlementPlanFeatureBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  planFeature.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitlementplanfeature",
		},
		{
			name:        "happy path, delete plan feature",
			idToDelete:  planFeature.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "plan feature already deleted, not found",
			idToDelete:  planFeature.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entitlement_plan_feature not found",
		},
		{
			name:        "unknown plan feature, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "entitlement_plan_feature not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if planFeature exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteEntitlementPlanFeature(reqCtx, planFeature.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, planFeature.ID, resp.DeleteEntitlementPlanFeature.DeletedID)
		})
	}
}
