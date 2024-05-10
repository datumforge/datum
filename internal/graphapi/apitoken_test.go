package graphapi_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/datumclient"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func (suite *GraphTestSuite) TestQueryApiToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	apiToken := (&APITokenTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		errorMsg string
	}{
		{
			name:    "happy path",
			queryID: apiToken.ID,
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

				// mock a call to check orgs
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", testPersonalOrgID)})
			}

			resp, err := suite.client.datum.GetAPITokenByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.APIToken)
			assert.Equal(t, redacted, resp.APIToken.Token)
			assert.Equal(t, testPersonalOrgID, resp.APIToken.Owner.ID)
		})
	}
}

func (suite *GraphTestSuite) TestQueryAPITokens() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	(&APITokenTokenBuilder{client: suite.client}).MustNew(reqCtx, t)
	(&APITokenTokenBuilder{client: suite.client, Scopes: []string{"read", "write"}}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		errorMsg string
	}{
		{
			name: "happy path, all api tokens",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", testPersonalOrgID)})
			}

			resp, err := suite.client.datum.GetAllAPITokens(reqCtx)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			assert.Len(t, resp.APITokens.Edges, 2)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreateAPIToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	tokenDescription := gofakeit.Sentence(5)
	expiration30Days := time.Now().Add(time.Hour * 24 * 30)

	testCases := []struct {
		name     string
		input    datumclient.CreateAPITokenInput
		errorMsg string
	}{
		{
			name: "happy path",
			input: datumclient.CreateAPITokenInput{
				Name:        "forthethingz",
				Description: &tokenDescription,
				Scopes:      []string{"read", "write"},
			},
		},
		{
			name: "happy path, set expire",
			input: datumclient.CreateAPITokenInput{
				Name:        "forthethingz",
				Description: &tokenDescription,
				ExpiresAt:   &expiration30Days,
			},
		},
		{
			name: "happy path, set org",
			input: datumclient.CreateAPITokenInput{
				Name:        "forthethingz",
				Description: &tokenDescription,
				ExpiresAt:   &expiration30Days,
			},
		},
		{
			name: "happy path, name only",
			input: datumclient.CreateAPITokenInput{
				Name: "forthethingz",
			},
		},
		{
			name: "empty name",
			input: datumclient.CreateAPITokenInput{
				Description: &tokenDescription,
			},
			errorMsg: "value is less than the required length",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, true)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", testPersonalOrgID)})
				mock_fga.WriteAny(t, suite.client.fga)
			}

			resp, err := suite.client.datum.CreateAPIToken(reqCtx, tc.input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateAPIToken.APIToken)

			assert.Equal(t, tc.input.Name, resp.CreateAPIToken.APIToken.Name)
			assert.Equal(t, tc.input.Description, resp.CreateAPIToken.APIToken.Description)
			assert.Equal(t, tc.input.Scopes, resp.CreateAPIToken.APIToken.Scopes)

			// check expiration if set
			if tc.input.ExpiresAt != nil {
				assert.Equal(t, &expiration30Days, tc.input.ExpiresAt)
			}

			// ensure the owner is the org set in the request
			assert.Equal(t, testPersonalOrgID, resp.CreateAPIToken.APIToken.Owner.ID)

			// token should not be redacted on create
			assert.NotEqual(t, redacted, resp.CreateAPIToken.APIToken.Token)

			// ensure the token is prefixed
			assert.Contains(t, resp.CreateAPIToken.APIToken.Token, "dtma_")
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateAPIToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	token := (&APITokenTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	tokenDescription := gofakeit.Sentence(5)
	tokenName := gofakeit.Word()

	testCases := []struct {
		name     string
		tokenID  string
		input    datumclient.UpdateAPITokenInput
		errorMsg string
	}{
		{
			name:    "happy path, update name ",
			tokenID: token.ID,
			input: datumclient.UpdateAPITokenInput{
				Name: &tokenName,
			},
		},
		{
			name:    "happy path, update description",
			tokenID: token.ID,
			input: datumclient.UpdateAPITokenInput{
				Description: &tokenDescription,
			},
		},
		{
			name:    "happy path, add scope",
			tokenID: token.ID,
			input: datumclient.UpdateAPITokenInput{
				Scopes: []string{"write"},
			},
		},
		{
			name:    "invalid token id",
			tokenID: "notvalidtoken",
			input: datumclient.UpdateAPITokenInput{
				Description: &tokenDescription,
			},
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.errorMsg == "" {
				mock_fga.CheckAny(t, suite.client.fga, true)
				// mock a call to check orgs
				mock_fga.ListAny(t, suite.client.fga, []string{fmt.Sprintf("organization:%s", testPersonalOrgID)})
			}

			resp, err := suite.client.datum.UpdateAPIToken(reqCtx, tc.tokenID, tc.input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateAPIToken.APIToken)

			if tc.input.Name != nil {
				assert.Equal(t, resp.UpdateAPIToken.APIToken.Name, *tc.input.Name)
			}

			if tc.input.Description != nil {
				assert.Equal(t, resp.UpdateAPIToken.APIToken.Description, tc.input.Description)
			}

			// Ensure its added
			if tc.input.Scopes != nil {
				assert.Len(t, resp.UpdateAPIToken.APIToken.Scopes, 1)
			}

			assert.Equal(t, testPersonalOrgID, resp.UpdateAPIToken.APIToken.Owner.ID)

			// token should be redacted on update
			assert.Equal(t, redacted, resp.UpdateAPIToken.APIToken.Token)
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeleteAPIToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create user to make tokens
	user := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	user2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	orgID := user.Edges.Setting.Edges.DefaultOrg.ID
	orgID2 := user2.Edges.Setting.Edges.DefaultOrg.ID

	reqCtx, err = auth.NewTestContextWithOrgID(user.ID, orgID)
	require.NoError(t, err)

	token := (&APITokenTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx2, err := auth.NewTestContextWithOrgID(user2.ID, orgID2)
	require.NoError(t, err)

	token2 := (&APITokenTokenBuilder{client: suite.client}).MustNew(reqCtx2, t)

	testCases := []struct {
		name     string
		tokenID  string
		errorMsg string
		allowed  bool
	}{
		{
			name:    "happy path, delete token",
			tokenID: token.ID,
			allowed: true,
		},
		{
			name:     "delete someone else's token, no go",
			tokenID:  token2.ID,
			errorMsg: "not authorized",
			allowed:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			mock_fga.CheckAny(t, suite.client.fga, tc.allowed)

			resp, err := suite.client.datum.DeleteAPIToken(reqCtx, tc.tokenID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.tokenID, resp.DeleteAPIToken.DeletedID)
		})
	}
}
