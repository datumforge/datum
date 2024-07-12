package graphapi_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/pkg/datumclient"

	"github.com/datumforge/datum/pkg/testutils"
)

const (
	notFoundErrorMsg = "personal_access_token not found"
	redacted         = "*****************************"
)

func (suite *GraphTestSuite) TestQueryPersonalAccessToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create user to get tokens
	user := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	token := (&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		queryID  string
		errorMsg string
	}{
		{
			name:    "happy path pat",
			queryID: token.ID,
		},
		{
			name:     "not found",
			queryID:  "notfound",
			errorMsg: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetPersonalAccessTokenByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.PersonalAccessToken)
			assert.Equal(t, redacted, resp.PersonalAccessToken.Token)
		})
	}
}

func (suite *GraphTestSuite) TestQueryPersonalAccessTokens() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	(&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	// create user to get tokens
	user := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	(&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)
	(&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		errorMsg string
	}{
		{
			name: "happy path, all pats",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.GetAllPersonalAccessTokens(reqCtx)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)

			assert.Len(t, resp.PersonalAccessTokens.Edges, 2)
		})
	}
}

func (suite *GraphTestSuite) TestMutationCreatePersonalAccessToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create user to get tokens
	user2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	tokenDescription := gofakeit.Sentence(5)
	expiration30Days := time.Now().Add(time.Hour * 24 * 30)

	testCases := []struct {
		name     string
		input    datumclient.CreatePersonalAccessTokenInput
		errorMsg string
	}{
		{
			name: "happy path",
			input: datumclient.CreatePersonalAccessTokenInput{
				Name:        "forthethingz",
				Description: &tokenDescription,
			},
		},
		{
			name: "happy path, set expire",
			input: datumclient.CreatePersonalAccessTokenInput{
				Name:        "forthethingz",
				Description: &tokenDescription,
				ExpiresAt:   expiration30Days,
			},
		},
		{
			name: "happy path, set org",
			input: datumclient.CreatePersonalAccessTokenInput{
				Name:            "forthethingz",
				Description:     &tokenDescription,
				ExpiresAt:       expiration30Days,
				OrganizationIDs: []string{org.ID, testPersonalOrgID},
			},
		},
		{
			name: "happy path, name only",
			input: datumclient.CreatePersonalAccessTokenInput{
				Name: "forthethingz",
			},
		},
		{
			name: "empty name",
			input: datumclient.CreatePersonalAccessTokenInput{
				Description: &tokenDescription,
			},
			errorMsg: "value is less than the required length",
		},
		{
			name: "setting other user id",
			input: datumclient.CreatePersonalAccessTokenInput{
				OwnerID:     user2.ID, // this should get ignored
				Name:        "forthethingz",
				Description: &tokenDescription,
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.CreatePersonalAccessToken(reqCtx, tc.input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreatePersonalAccessToken.PersonalAccessToken)
			assert.Equal(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Name, tc.input.Name)
			assert.Equal(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Description, tc.input.Description)

			// check expiration if set
			if !tc.input.ExpiresAt.IsZero() {
				assert.Equal(t, expiration30Days, tc.input.ExpiresAt)
			}

			// check organization is set if provided
			if tc.input.OrganizationIDs != nil {
				assert.Len(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Organizations, len(tc.input.OrganizationIDs))

				for _, orgID := range resp.CreatePersonalAccessToken.PersonalAccessToken.Organizations {
					assert.Contains(t, tc.input.OrganizationIDs, orgID.ID)
				}
			} else {
				assert.Len(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Organizations, 0)
			}

			// ensure the owner is the user that made the request
			assert.Equal(t, testUser.ID, resp.CreatePersonalAccessToken.PersonalAccessToken.Owner.ID)

			// token should not be redacted on create
			assert.NotEqual(t, redacted, resp.CreatePersonalAccessToken.PersonalAccessToken.Token)

			// ensure the token is prefixed
			assert.Contains(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Token, "dtmp_")
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdatePersonalAccessToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(reqCtx, t)

	// setup a token for another user
	user2 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	regCtx2, err := userContextWithID(user2.ID)
	require.NoError(t, err)
	tokenOther := (&PersonalAccessTokenBuilder{
		client:  suite.client,
		OwnerID: user2.ID}).
		MustNew(regCtx2, t)

	token := (&PersonalAccessTokenBuilder{
		client:          suite.client,
		OwnerID:         testUser.ID,
		OrganizationIDs: []string{testPersonalOrgID}}).
		MustNew(reqCtx, t)

	tokenDescription := gofakeit.Sentence(5)
	tokenName := gofakeit.Word()

	testCases := []struct {
		name     string
		tokenID  string
		input    datumclient.UpdatePersonalAccessTokenInput
		errorMsg string
	}{
		{
			name:    "happy path, update name",
			tokenID: token.ID,
			input: datumclient.UpdatePersonalAccessTokenInput{
				Name: &tokenName,
			},
		},
		{
			name:    "happy path, update description",
			tokenID: token.ID,
			input: datumclient.UpdatePersonalAccessTokenInput{
				Description: &tokenDescription,
			},
		},
		{
			name:    "happy path, add org",
			tokenID: token.ID,
			input: datumclient.UpdatePersonalAccessTokenInput{
				AddOrganizationIDs: []string{org.ID},
			},
		},
		{
			name:    "happy path, remove org",
			tokenID: token.ID,
			input: datumclient.UpdatePersonalAccessTokenInput{
				RemoveOrganizationIDs: []string{org.ID},
			},
		},
		{
			name:    "invalid token id",
			tokenID: "notvalidtoken",
			input: datumclient.UpdatePersonalAccessTokenInput{
				Description: &tokenDescription,
			},
			errorMsg: notFoundErrorMsg,
		},
		{
			name:    "not authorized",
			tokenID: tokenOther.ID,
			input: datumclient.UpdatePersonalAccessTokenInput{
				Description: &tokenDescription,
			},
			errorMsg: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			resp, err := suite.client.datum.UpdatePersonalAccessToken(reqCtx, tc.tokenID, tc.input)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdatePersonalAccessToken.PersonalAccessToken)

			if tc.input.Name != nil {
				assert.Equal(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Name, *tc.input.Name)
			}

			if tc.input.Description != nil {
				assert.Equal(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Description, tc.input.Description)
			}

			// make sure these fields did not get updated
			assert.WithinDuration(t, *token.ExpiresAt, resp.UpdatePersonalAccessToken.PersonalAccessToken.ExpiresAt, 1*time.Second)
			assert.Len(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Organizations, len(tc.input.AddOrganizationIDs)+1)

			// Ensure its removed
			if tc.input.RemoveOrganizationIDs != nil {
				assert.Len(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Organizations, 1)
			}

			assert.Equal(t, testUser.ID, resp.UpdatePersonalAccessToken.PersonalAccessToken.Owner.ID)

			// token should be redacted on update
			assert.Equal(t, redacted, resp.UpdatePersonalAccessToken.PersonalAccessToken.Token)
		})
	}
}

func (suite *GraphTestSuite) TestMutationDeletePersonalAccessToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// token for another user
	tokenOther := (&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	// create user
	user := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	token := (&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	testCases := []struct {
		name     string
		tokenID  string
		errorMsg string
	}{
		{
			name:    "happy path, delete token",
			tokenID: token.ID,
		},
		{
			name:     "delete someone else's token, no go",
			tokenID:  tokenOther.ID,
			errorMsg: notFoundErrorMsg,
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := suite.client.datum.DeletePersonalAccessToken(reqCtx, tc.tokenID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.Equal(t, tc.tokenID, resp.DeletePersonalAccessToken.DeletedID)
		})
	}
}

func (suite *GraphTestSuite) TestLastUsedPersonalAccessToken() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create new personal access token
	token := (&PersonalAccessTokenBuilder{client: suite.client}).MustNew(reqCtx, t)

	// check that the last used is empty
	res, err := suite.client.datum.GetPersonalAccessTokenByID(reqCtx, token.ID)
	require.NoError(t, err)
	assert.Empty(t, res.PersonalAccessToken.LastUsedAt)

	// TODO: (slevine: update once we have updated the last used at field on the token when used)
	// // setup graph client using the personal access token
	authHeader := datumclient.Authorization{
		BearerToken: token.Token,
	}

	graphClient, err := testutils.DatumTestClientWithAuth(t, suite.client.db, datumclient.WithCredentials(authHeader))
	require.NoError(t, err)

	// get the token to make sure the last used is updated using the token
	out, err := graphClient.GetPersonalAccessTokenByID(reqCtx, token.ID)
	require.NoError(t, err)
	assert.NotEmpty(t, out.PersonalAccessToken.LastUsedAt)
}
