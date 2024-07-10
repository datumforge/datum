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

func (suite *GraphTestSuite) TestQueryEntityType() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entityType := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		allowed  bool
		expected *ent.EntityType
		errorMsg string
	}{
		{
			name:     "happy path entityType",
			allowed:  true,
			queryID:  entityType.ID,
			expected: entityType,
		},
		{
			name:     "no access",
			allowed:  false,
			queryID:  entityType.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.GetEntityTypeByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.EntityType)
		})
	}

	// delete created org and entityType
	(&EntityTypeCleanup{client: suite.client, ID: entityType.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryEntityTypes() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	_ = (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)
	_ = (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)

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

			resp, err := suite.client.datum.GetAllEntityTypes(tc.ctx)
			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.EntityTypes.Edges, tc.expectedResults)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateEntityType() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		request     datumclient.CreateEntityTypeInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, all input",
			request: datumclient.CreateEntityTypeInput{
				Name: "cats",
			},
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntityTypeInput{
				Name: "dogs",
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitytype",
		},
		{
			name:        "missing required field, name",
			request:     datumclient.CreateEntityTypeInput{},
			allowed:     true,
			expectedErr: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.CreateEntityType(reqCtx, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Equal(t, tc.request.Name, resp.CreateEntityType.EntityType.Name)
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateEntityType() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entityType := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		request     datumclient.UpdateEntityTypeInput
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update name",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("maine coons"),
			},
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("dogs"),
			},
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitytype",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.UpdateEntityType(reqCtx, entityType.ID, tc.request)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, *tc.request.Name, resp.UpdateEntityType.EntityType.Name)
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteEntityType() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	entityType := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name        string
		idToDelete  string
		allowed     bool
		checkAccess bool
		expectedErr string
	}{
		{
			name:        "not allowed to delete",
			idToDelete:  entityType.ID,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitytype",
		},
		{
			name:        "happy path, delete entitytype",
			idToDelete:  entityType.ID,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "entityType already deleted, not found",
			idToDelete:  entityType.ID,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entity_type not found",
		},
		{
			name:        "unknown entitytype, not found",
			idToDelete:  ulids.New().String(),
			checkAccess: false,
			allowed:     true,
			expectedErr: "entity_type not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Delete "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization if entityType exists
			if tc.checkAccess {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			resp, err := suite.client.datum.DeleteEntityType(reqCtx, entityType.ID)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, entityType.ID, resp.DeleteEntityType.DeletedID)
		})
	}
}
