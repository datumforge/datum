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
		client   *datumclient.DatumClient
		ctx      context.Context
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: plan.ID,
			client:  suite.client.datum,
			ctx:     reqCtx,
		},
		{
			name:    "happy path, using api token",
			queryID: plan.ID,
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
		},
		{
			name:    "happy path, using personal access token",
			queryID: plan.ID,
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
		},
		{
			name:     "not found",
			queryID:  "notfound",
			client:   suite.client.datum,
			ctx:      reqCtx,
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errorMsg == "" {
				mock_fga.CheckAny(t, suite.client.fga, true)
			}

			resp, err := tc.client.GetEntitlementPlanByID(tc.ctx, tc.queryID)

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
		client          *datumclient.DatumClient
		ctx             context.Context
		expectedResults int
	}{
		{
			name:            "happy path",
			client:          suite.client.datum,
			ctx:             reqCtx,
			expectedResults: 2,
		},
		{
			name:            "happy path, using api token",
			client:          suite.client.datumWithAPIToken,
			ctx:             context.Background(),
			expectedResults: 2,
		},
		{
			name:            "happy path, using pat",
			client:          suite.client.datumWithPAT,
			ctx:             context.Background(),
			expectedResults: 2,
		},
		{
			name:            "another user, no plans should be returned",
			client:          suite.client.datum,
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := tc.client.GetAllEntitlementPlans(tc.ctx)
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
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateEntitlementPlanInput{
				Name:    "test-plan",
				Version: "v1",
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "happy path, using api token",
			request: datumclient.CreateEntitlementPlanInput{
				Name:    "test-plan-with-api-token",
				Version: "v1",
			},
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "happy path, using personal access token",
			request: datumclient.CreateEntitlementPlanInput{
				OwnerID: &testPersonalOrgID,
				Name:    "test-plan-with-pat",
				Version: "v1",
			},
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
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
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntitlementPlanInput{
				Name: "test-plan",
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitlementplan",
		},
		{
			name: "missing required field, version",
			request: datumclient.CreateEntitlementPlanInput{
				Name: "Matt is the Best",
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := tc.client.CreateEntitlementPlan(tc.ctx, tc.request)
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
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update display name",
			request: datumclient.UpdateEntitlementPlanInput{
				DisplayName: lo.ToPtr("test-plan"),
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "update description, using api token",
			request: datumclient.UpdateEntitlementPlanInput{
				Description: lo.ToPtr("To infinity and beyond!"),
			},
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "update description again, using personal access token",
			request: datumclient.UpdateEntitlementPlanInput{
				Description: lo.ToPtr("To infinity and beyond, and beyond!"),
			},
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntitlementPlanInput{
				Description: lo.ToPtr("Howdy, partner!"),
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitlementplan",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := tc.client.UpdateEntitlementPlan(tc.ctx, plan.ID, tc.request)
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

	plan1 := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)
	plan2 := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)
	plan3 := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  plan1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitlementplan",
		},
		{
			name:        "happy path, delete plan",
			idToDelete:  plan1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "plan already deleted, not found",
			idToDelete:  plan1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: false,
			allowed:     true,
			expectedErr: "plan not found",
		},
		{
			name:        "happy path, delete plan using api token",
			idToDelete:  plan2.ID,
			client:      suite.client.datumWithAPIToken,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "happy path, delete plan using personal access token",
			idToDelete:  plan3.ID,
			client:      suite.client.datumWithPAT,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "unknown plan, not found",
			idToDelete:  ulids.New().String(),
			client:      suite.client.datum,
			ctx:         reqCtx,
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

			resp, err := tc.client.DeleteEntitlementPlan(tc.ctx, tc.idToDelete)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.idToDelete, resp.DeleteEntitlementPlan.DeletedID)
		})
	}
}
