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
		client   *datumclient.DatumClient
		ctx      context.Context
		allowed  bool
		expected *ent.EntityType
		errorMsg string
	}{
		{
			name:     "happy path entity type",
			client:   suite.client.datum,
			ctx:      reqCtx,
			allowed:  true,
			queryID:  entityType.ID,
			expected: entityType,
		},
		{
			name:     "happy path entity type, using api token",
			client:   suite.client.datumWithAPIToken,
			ctx:      context.Background(),
			allowed:  true,
			queryID:  entityType.ID,
			expected: entityType,
		},
		{
			name:     "happy path entity type, using pat",
			client:   suite.client.datumWithPAT,
			ctx:      context.Background(),
			allowed:  true,
			queryID:  entityType.ID,
			expected: entityType,
		},
		{
			name:     "no access",
			client:   suite.client.datum,
			ctx:      reqCtx,
			allowed:  false,
			queryID:  entityType.ID,
			errorMsg: "not authorized",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := tc.client.GetEntityTypeByID(tc.ctx, tc.queryID)

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
		client          *datumclient.DatumClient
		ctx             context.Context
		expectedResults int
	}{
		{
			name:            "happy path",
			client:          suite.client.datum,
			ctx:             reqCtx,
			expectedResults: 3, // 1 is created in the setup
		},
		{
			name:            "happy path, using api token",
			client:          suite.client.datumWithAPIToken,
			ctx:             context.Background(),
			expectedResults: 3, // 1 is created in the setup
		},
		{
			name:            "happy path, using pat",
			client:          suite.client.datumWithPAT,
			ctx:             context.Background(),
			expectedResults: 3, // 1 is created in the setup
		},
		{
			name:            "another user, no entities should be returned",
			client:          suite.client.datum,
			ctx:             otherCtx,
			expectedResults: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("List "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := tc.client.GetAllEntityTypes(tc.ctx)
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
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, all input",
			request: datumclient.CreateEntityTypeInput{
				Name: "cats",
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "happy path, all input, using api token",
			request: datumclient.CreateEntityTypeInput{
				Name: "horses",
			},
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "happy path, all input, using pat",
			request: datumclient.CreateEntityTypeInput{
				OwnerID: &testOrgID,
				Name:    "bunnies",
			},
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "do not create if not allowed",
			request: datumclient.CreateEntityTypeInput{
				Name: "dogs",
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: create on entitytype",
		},
		{
			name:        "missing required field, name",
			request:     datumclient.CreateEntityTypeInput{},
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

			resp, err := tc.client.CreateEntityType(tc.ctx, tc.request)
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
		client      *datumclient.DatumClient
		ctx         context.Context
		allowed     bool
		expectedErr string
	}{
		{
			name: "happy path, update name",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("maine coons"),
			},
			client:  suite.client.datum,
			ctx:     reqCtx,
			allowed: true,
		},
		{
			name: "happy path, update name using api token",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("sphynx"),
			},
			client:  suite.client.datumWithAPIToken,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "happy path, update name using personal access token",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("persian"),
			},
			client:  suite.client.datumWithPAT,
			ctx:     context.Background(),
			allowed: true,
		},
		{
			name: "not allowed to update",
			request: datumclient.UpdateEntityTypeInput{
				Name: lo.ToPtr("dogs"),
			},
			client:      suite.client.datum,
			ctx:         reqCtx,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: update on entitytype",
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// check for edit permissions on the organization
			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := tc.client.UpdateEntityType(tc.ctx, entityType.ID, tc.request)
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

	entityType1 := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)
	entityType2 := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)
	entityType3 := (&EntityTypeBuilder{client: suite.client}).MustNew(reqCtx, t)

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
			idToDelete:  entityType1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     false,
			expectedErr: "you are not authorized to perform this action: delete on entitytype",
		},
		{
			name:        "happy path, delete entity type",
			idToDelete:  entityType1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "entityType already deleted, not found",
			idToDelete:  entityType1.ID,
			client:      suite.client.datum,
			ctx:         reqCtx,
			checkAccess: false,
			allowed:     true,
			expectedErr: "entity_type not found",
		},
		{
			name:        "happy path, delete entity type using api token",
			idToDelete:  entityType2.ID,
			client:      suite.client.datumWithAPIToken,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "happy path, delete entity type using pat",
			idToDelete:  entityType3.ID,
			client:      suite.client.datumWithPAT,
			ctx:         context.Background(),
			checkAccess: true,
			allowed:     true,
		},
		{
			name:        "unknown entitytype, not found",
			idToDelete:  ulids.New().String(),
			client:      suite.client.datum,
			ctx:         reqCtx,
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

			resp, err := tc.client.DeleteEntityType(tc.ctx, tc.idToDelete)
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.expectedErr)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Equal(t, tc.idToDelete, resp.DeleteEntityType.DeletedID)
		})
	}
}
