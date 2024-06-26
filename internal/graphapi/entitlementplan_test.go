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

func (suite *GraphTestSuite) TestQueryEntitlementPlan() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	plan := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: plan.ID,
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

			resp, err := suite.client.datum.GetEntitlementPlanByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.queryID, resp.EntitlementPlan.ID)
			assert.NotEmpty(t, resp.EntitlementPlan.Name)
			assert.NotEmpty(t, resp.EntitlementPlan.Version)
			assert.NotEmpty(t, resp.EntitlementPlan.Description)
			assert.NotEmpty(t, resp.EntitlementPlan.DisplayName)
		})
	}
}

func (suite *GraphTestSuite) TestQueryEntitlementPlans() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			name:            "another user, no plans should be returned",
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllEntitlementPlans(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.EntitlementPlans.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateEntitlementPlan() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		request     datumclient.CreateEntitlementPlanInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateEntitlementPlanInput{
				Name:    "test-plan",
				Version: "v1",
			},
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateEntitlementPlanInput{
				Name:        "mitb",
				Version:     "v1",
				DisplayName: lo.ToPtr("Matt is the Best"),
				Description: lo.ToPtr("Matt is the best plan, hands down!"),
			},
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntitlementPlanInput{
				Name: "test-plan",
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitlementplan",
		},
		{
			name: "missing required field, version",
			request: datumclient.CreateEntitlementPlanInput{
				Name: "Matt is the Best",
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

			resp, err := suite.client.datum.CreateEntitlementPlan(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.Name, resp.CreateEntitlementPlan.EntitlementPlan.Name)
			assert.Equal(t, tc.request.Version, resp.CreateEntitlementPlan.EntitlementPlan.Version)

			if tc.request.Description == nil {
				assert.Empty(t, resp.CreateEntitlementPlan.EntitlementPlan.Description)
			} else {
				assert.Equal(t, *tc.request.Description, *resp.CreateEntitlementPlan.EntitlementPlan.Description)
			}

			// Display Name is set to the Name if not provided
			if tc.request.DisplayName == nil {
				assert.Equal(t, tc.request.Name, *resp.CreateEntitlementPlan.EntitlementPlan.DisplayName)
			} else {
				assert.Equal(t, *tc.request.DisplayName, *resp.CreateEntitlementPlan.EntitlementPlan.DisplayName)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateEntitlementPlan() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	plan := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateEntitlementPlanInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update display name",
			request: datumclient.UpdateEntitlementPlanInput{
				DisplayName: lo.ToPtr("test-plan"),
			},
			allowed: true,
		},
		{
			name: "update description",
			request: datumclient.UpdateEntitlementPlanInput{
				Description: lo.ToPtr("To infinity and beyond!"),
			},
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntitlementPlanInput{
				Description: lo.ToPtr("Howdy, partner!"),
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitlementplan",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateEntitlementPlan(reqCtx, plan.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.Description != nil {
				assert.Equal(t, *tc.request.Description, *resp.UpdateEntitlementPlan.EntitlementPlan.Description)
			}

			if tc.request.DisplayName != nil {
				assert.Equal(t, *tc.request.DisplayName, *resp.UpdateEntitlementPlan.EntitlementPlan.DisplayName)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteEntitlementPlan() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	plan := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  plan.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitlementplan",
		},
		{
			name:        "happy path, delete plan",
			idToDelete:  plan.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "plan already deleted, not found",
			idToDelete:  plan.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "plan not found",
		},
		{
			name:        "unknown plan, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "plan not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if plan exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteEntitlementPlan(reqCtx, plan.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, plan.ID, resp.DeleteEntitlementPlan.DeletedID)
		})
	}
}
