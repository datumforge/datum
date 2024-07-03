package graphapi_test

import (
	"context"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

func (suite *GraphTestSuite) TestQueryEntity() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entity := (&EntityBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		allowed  bool
		expected *ent.Entity
		errorMsg string
	}{
		{
			name:     "happy path entity",
			allowed:  true,
			queryID:  entity.ID,
			expected: entity,
		},
		{
			name:     "no access",
			allowed:  false,
			queryID:  entity.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.GetEntityByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Entity)
		})
	}

	// delete created org and entity
	(&EntityCleanup{client: suite.client, ID: entity.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryEntities() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&EntityBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&EntityBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			name:            "another user, no entities should be returned",
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllEntities(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.Entities.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateEntity() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		request     datumclient.CreateEntityInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, minimal input",
			request: datumclient.CreateEntityInput{
				Name: "fraser fir",
			},
			allowed: true,
		},
		{
			name: "happy path, all input",
			request: datumclient.CreateEntityInput{
				Name:        "mitb",
				DisplayName: lo.ToPtr("fraser fir"),
				Description: lo.ToPtr("the pine trees of appalachia"),
			},
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntityInput{
				Name: "test-entity",
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entity",
		},
		{
			name: "missing required field, name",
			request: datumclient.CreateEntityInput{
				DisplayName: lo.ToPtr("fraser firs"),
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

			resp, err := suite.client.datum.CreateEntity(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.Name, resp.CreateEntity.Entity.Name)

			if tc.request.Description == nil {
				assert.Empty(t, resp.CreateEntity.Entity.Description)
			} else {
				assert.Equal(t, *tc.request.Description, *resp.CreateEntity.Entity.Description)
			}

			// Display Name is set to the Name if not provided
			if tc.request.DisplayName == nil {
				assert.Equal(t, tc.request.Name, resp.CreateEntity.Entity.DisplayName)
			} else {
				assert.Equal(t, *tc.request.DisplayName, resp.CreateEntity.Entity.DisplayName)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateEntity() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entity := (&EntityBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateEntityInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update display name",
			request: datumclient.UpdateEntityInput{
				DisplayName: lo.ToPtr("blue spruce"),
			},
			allowed: true,
		},
		{
			name: "update description",
			request: datumclient.UpdateEntityInput{
				Description: lo.ToPtr("the pine tree with blue-green colored needles"),
			},
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntityInput{
				Description: lo.ToPtr("pine trees of the west"),
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entity",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateEntity(reqCtx, entity.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			if tc.request.Description != nil {
				assert.Equal(t, *tc.request.Description, *resp.UpdateEntity.Entity.Description)
			}

			if tc.request.DisplayName != nil {
				assert.Equal(t, *tc.request.DisplayName, resp.UpdateEntity.Entity.DisplayName)
			}
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteEntity() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entity := (&EntityBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  entity.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entity",
		},
		{
			name:        "happy path, delete entity",
			idToDelete:  entity.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "entity already deleted, not found",
			idToDelete:  entity.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entity not found",
		},
		{
			name:        "unknown entity, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "entity not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if entity exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteEntity(reqCtx, entity.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, entity.ID, resp.DeleteEntity.DeletedID)
		})
	}
}
