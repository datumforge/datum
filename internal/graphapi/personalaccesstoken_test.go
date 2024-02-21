package graphapi_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
)

const (
	notFoundErrorMsg = "personal_access_token not found"
	redacted         = "*****************************"
)

func TestQueryPersonalAccessToken(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create user to get tokens
	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	token := (&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, client.fga, []string{"organization:test"})
			}

			resp, err := client.datum.GetPersonalAccessTokenByID(reqCtx, tc.queryID)

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

func TestQueryPersonalAccessTokens(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	(&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

	// create user to get tokens
	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	(&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)
	(&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, client.fga, []string{"organization:test"})
			}

			resp, err := client.datum.GetAllPersonalAccessTokens(reqCtx)

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

func TestMutationCreatePersonalAccessToken(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// create user to get tokens
	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)
	user2 := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	org := (&OrganizationBuilder{client: client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

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
				OrganizationIDs: []string{org.ID},
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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, client.fga, []string{fmt.Sprintf("organization:%s", org.ID)})
			}

			resp, err := client.datum.CreatePersonalAccessToken(reqCtx, tc.input)

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
				assert.Nil(t, resp.CreatePersonalAccessToken.PersonalAccessToken.Organizations)
			}

			// ensure the owner is the user that made the request
			assert.Equal(t, user.ID, resp.CreatePersonalAccessToken.PersonalAccessToken.Owner.ID)

			// token should not be redacted on create
			assert.NotEqual(t, redacted, resp.CreatePersonalAccessToken.PersonalAccessToken.Token)
		})
	}
}

func TestMutationUpdatePersonalAccessToken(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// token for another user
	tokenOther := (&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

	// create user
	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	org := (&OrganizationBuilder{client: client}).MustNew(reqCtx, t)

	token := (&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

	tokenDescription := gofakeit.Sentence(5)
	tokenName := gofakeit.Word()

	testCases := []struct {
		name     string
		tokenID  string
		input    datumclient.UpdatePersonalAccessTokenInput
		errorMsg string
	}{
		{
			name:    "happy path, update name ",
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
			defer mock_fga.ClearMocks(client.fga)

			if tc.errorMsg == "" {
				// mock a call to check orgs
				mock_fga.ListAny(t, client.fga, []string{fmt.Sprintf("organization:%s", org.ID)})
			}

			resp, err := client.datum.UpdatePersonalAccessToken(reqCtx, tc.tokenID, tc.input)

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
			assert.Len(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Organizations, len(tc.input.AddOrganizationIDs))

			// Ensure its removed
			if tc.input.RemoveOrganizationIDs != nil {
				assert.Len(t, resp.UpdatePersonalAccessToken.PersonalAccessToken.Organizations, 0)
			}

			assert.Equal(t, user.ID, resp.UpdatePersonalAccessToken.PersonalAccessToken.Owner.ID)

			// token should be redacted on update
			assert.Equal(t, redacted, resp.UpdatePersonalAccessToken.PersonalAccessToken.Token)
		})
	}
}

func TestMutationDeletePersonalAccessToken(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	// token for another user
	tokenOther := (&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

	// create user
	user := (&UserBuilder{client: client}).MustNew(reqCtx, t)

	reqCtx, err = userContextWithID(user.ID)
	require.NoError(t, err)

	token := (&PersonalAccessTokenBuilder{client: client}).MustNew(reqCtx, t)

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
			resp, err := client.datum.DeletePersonalAccessToken(reqCtx, tc.tokenID)

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
