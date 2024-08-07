package graphapi_test

import (
	"context"
	"testing"
	"time"

	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/ulids"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestQueryEntitlement() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entitlement := (&EntitlementBuilder{client: suite.client, OrganizationID: testOrgID}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		client   *datumclient.DatumClient
		ctx      context.Context
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: entitlement.ID,
			client:  suite.client.datum,
			ctx:     reqCtx,
		},
		{
			name:    "happy path, using api token",
			queryID: entitlement.ID,
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
		},
		{
			name:    "happy path, using personal access token",
			queryID: entitlement.ID,
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

			resp, err := tc.client.GetEntitlementByID(tc.ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.queryID, resp.Entitlement.ID)

			require.NotEmpty(t, resp.Entitlement.GetPlan())
			assert.NotEmpty(t, resp.Entitlement.Plan.ID)

			require.NotEmpty(t, resp.Entitlement.GetOrganization())
			assert.NotEmpty(t, resp.Entitlement.Organization.ID)
		})
	}
}

func (suite *GraphTestSuite) TestQueryEntitlements() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&EntitlementBuilder{client: suite.client, OrganizationID: testOrgID}).MustNew(reqCtx, t)

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
			expectedResults: 1,
		},
		{
			name:            "happy path, using api token",
			client:          suite.client.datumWithAPIToken,
			ctx:             context.Background(),
			expectedResults: 1,
		},
		{
			name:            "happy path, using pat",
			client:          suite.client.datumWithPAT,
			ctx:             context.Background(),
			expectedResults: 1,
		},
		{
			name:            "another user, no entitlements should be returned",
			client:          suite.client.datum,
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := tc.client.GetAllEntitlements(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.Entitlements.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateEntitlement() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// setup for entitlement creation
	org1 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)
	plan := (&EntitlementPlanBuilder{client: suite.client}).MustNew(reqCtx, t)

	expiresAt := time.Now().Add(time.Hour * 24 * 365) // 1 year from now

	testCases := []struct {
		name        string
		request     datumclient.CreateEntitlementInput
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateEntitlementInput{
				PlanID:         plan.ID,
				OrganizationID: org1.ID,
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateEntitlementInput{
				PlanID:                 plan.ID,
				OrganizationID:         org2.ID,
				ExternalCustomerID:     lo.ToPtr("customer-123"),
				ExternalSubscriptionID: lo.ToPtr("sub-123"),
				Cancelled:              lo.ToPtr(false),
				ExpiresAt:              &expiresAt,
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "organization already has active entitlement",
			request: datumclient.CreateEntitlementInput{
				PlanID:         plan.ID,
				OrganizationID: org1.ID,
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     true,
			expectedErr: "entitlement already exists",
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntitlementInput{
				PlanID:         plan.ID,
				OrganizationID: org1.ID,
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitlement",
		},
		{
			name: "missing required field, organization",
			request: datumclient.CreateEntitlementInput{
				PlanID: plan.ID,
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
		{
			name: "missing required field, plan",
			request: datumclient.CreateEntitlementInput{
				OrganizationID: org1.ID,
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

			resp, err := tc.client.CreateEntitlement(tc.ctx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			require.NotEmpty(t, resp.CreateEntitlement.Entitlement.GetPlan())
			assert.Equal(t, tc.request.PlanID, resp.CreateEntitlement.Entitlement.Plan.ID)

			if tc.request.ExpiresAt != nil {
				assert.WithinDuration(t, *tc.request.ExpiresAt, *resp.CreateEntitlement.Entitlement.ExpiresAt, time.Second)
				assert.True(t, resp.CreateEntitlement.Entitlement.Expires)
			} else {
				assert.False(t, resp.CreateEntitlement.Entitlement.Expires)
			}

			if tc.request.ExternalCustomerID != nil {
				assert.Equal(t, *tc.request.ExternalCustomerID, *resp.CreateEntitlement.Entitlement.ExternalCustomerID)
			}

			if tc.request.ExternalSubscriptionID != nil {
				assert.Equal(t, *tc.request.ExternalSubscriptionID, *resp.CreateEntitlement.Entitlement.ExternalSubscriptionID)
			}

			if tc.request.Cancelled != nil {
				assert.Equal(t, *tc.request.Cancelled, resp.CreateEntitlement.Entitlement.Cancelled)
			} else {
				assert.False(t, resp.CreateEntitlement.Entitlement.Cancelled)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateEntitlement() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entitlement := (&EntitlementBuilder{client: suite.client}).MustNew(reqCtx, t)

	expiresAt := time.Now().Add(time.Hour * 24 * 30) // 30 days from now

	testCases := []struct {
		name        string
		request     datumclient.UpdateEntitlementInput
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update external customer id",
			request: datumclient.UpdateEntitlementInput{
				ExternalCustomerID: lo.ToPtr("customer-123"),
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "happy path, update external customer id using api token",
			request: datumclient.UpdateEntitlementInput{
				ExternalSubscriptionID: lo.ToPtr("sub-123"),
			},
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "happy path, expire entitlement using pat",
			request: datumclient.UpdateEntitlementInput{
				ExpiresAt: &expiresAt,
			},
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "cancel entitlement",
			request: datumclient.UpdateEntitlementInput{
				Cancelled: lo.ToPtr(true),
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntitlementInput{
				Cancelled: lo.ToPtr(false),
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitlement",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := tc.client.UpdateEntitlement(tc.ctx, entitlement.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.Cancelled != nil {
				assert.Equal(t, *tc.request.Cancelled, resp.UpdateEntitlement.Entitlement.GetCancelled())
			}

			if tc.request.ExternalCustomerID != nil {
				assert.Equal(t, *tc.request.ExternalCustomerID, *resp.UpdateEntitlement.Entitlement.GetExternalCustomerID())
			}

			if tc.request.ExternalSubscriptionID != nil {
				assert.Equal(t, *tc.request.ExternalSubscriptionID, *resp.UpdateEntitlement.Entitlement.GetExternalSubscriptionID())
			}

			if tc.request.ExpiresAt != nil {
				assert.WithinDuration(t, *tc.request.ExpiresAt, *resp.UpdateEntitlement.Entitlement.GetExpiresAt(), time.Second)
				assert.True(t, resp.UpdateEntitlement.Entitlement.GetExpires())
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteEntitlement() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entitlement1 := (&EntitlementBuilder{client: suite.client}).MustNew(reqCtx, t)
	entitlement2 := (&EntitlementBuilder{client: suite.client}).MustNew(reqCtx, t)
	entitlement3 := (&EntitlementBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			idToDelete:  entitlement1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitlement",
		},
		{
			name:        "happy path, delete entitlement",
			idToDelete:  entitlement1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "entitlement already deleted, not found",
			idToDelete:  entitlement1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entitlement not found",
		},
		{
			name:        "happy path, delete entitlement using api token",
			idToDelete:  entitlement2.ID,
			client:      suite.client.datumWithAPIToken,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "happy path, delete entitlement using pat",
			idToDelete:  entitlement3.ID,
			client:      suite.client.datumWithPAT,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "unknown entitlement, not found",
			idToDelete:  ulids.New().String(),
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entitlement not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if entitlement exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := tc.client.DeleteEntitlement(tc.ctx, tc.idToDelete)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.idToDelete, resp.DeleteEntitlement.DeletedID)
		})
	}
}
