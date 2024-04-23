package graphapi_test

import (
	"fmt"
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/datumclient"
)

func (suite *GraphTestSuite) TestQueryUserSetting() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	user1Setting, err := user1.Setting(ctx)
	require.NoError(t, err)

	user2 := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	user2Setting, err := user2.Setting(ctx)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user1.ID)
	require.NoError(t, err)

	testCases := []struct {
		name     string
		queryID  string
		expected *ent.UserSetting
		errorMsg string
	}{
		{
			name:     "happy path user",
			queryID:  user1Setting.ID,
			expected: user1Setting,
		},
		{
			name:     "valid user, but not auth",
			queryID:  user2Setting.ID,
			errorMsg: "not found",
		},
		{
			name:     "invalid-id",
			queryID:  "tacos-for-dinner",
			errorMsg: "not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			if tc.expected != nil {
				// placeholder until authz is enforced on org members
				// at which point this needs to be the correct organization id
				listObjects := []string{"organization:test"}
				mock_fga.ListAny(t, suite.client.fga, listObjects)
			}

			resp, err := suite.client.datum.GetUserSettingByID(reqCtx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UserSetting)
			require.Equal(t, tc.expected.Status, resp.UserSetting.Status)
		})
	}

	(&UserCleanup{client: suite.client, UserID: user1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: user2.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestQueryUserSettings() {
	t := suite.T()

	// setup user context
	reqCtx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)
	user1Setting, err := user1.Setting(reqCtx)
	require.NoError(t, err)

	// create another user to make sure we don't get their settings back
	_ = (&UserBuilder{client: suite.client}).MustNew(reqCtx, t)

	// mock list
	listObjects := []string{"organization:test"}

	t.Run("Get User Settings", func(t *testing.T) {
		defer mock_fga.ClearMocks(suite.client.fga)

		// mock list for default org on user settings
		mock_fga.ListAny(t, suite.client.fga, listObjects)

		resp, err := suite.client.datum.GetUserSettings(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.UserSettings.Edges)

		// make sure only the current user settings are returned
		assert.Equal(t, len(resp.UserSettings.Edges), 1)

		// setup valid user context
		reqCtx, err := userContextWithID(user1.ID)
		require.NoError(t, err)

		resp, err = suite.client.datum.GetUserSettings(reqCtx)

		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.UserSettings.Edges)
		require.Equal(t, user1Setting.ID, resp.UserSettings.Edges[0].Node.ID)
	})
}

func (suite *GraphTestSuite) TestMutationUpdateUserSetting() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	user1Setting, err := user.Setting(ctx)
	require.NoError(t, err)

	org := (&OrganizationBuilder{client: suite.client}).MustNew(ctx, t)
	org2 := (&OrganizationBuilder{client: suite.client}).MustNew(ctx, t)

	// mock list objects
	listObjects := []string{fmt.Sprintf("organization:%s", org.ID)}

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		updateInput datumclient.UpdateUserSettingInput
		expectedRes datumclient.UpdateUserSetting_UpdateUserSetting_UserSetting
		allowed     bool
		checkOrg    bool
		errorMsg    string
	}{
		{
			name: "update default org and tags",
			updateInput: datumclient.UpdateUserSettingInput{
				DefaultOrgID: &org.ID,
				Tags:         []string{"mitb", "datum"},
			},
			allowed:  true,
			checkOrg: true,
			expectedRes: datumclient.UpdateUserSetting_UpdateUserSetting_UserSetting{
				Status: enums.Active,
				Tags:   []string{"mitb", "datum"},
				DefaultOrg: &datumclient.UpdateUserSetting_UpdateUserSetting_UserSetting_DefaultOrg{
					ID: org.ID,
				},
			},
		},
		{
			name: "update default org to invalid org ID",
			updateInput: datumclient.UpdateUserSettingInput{
				DefaultOrgID: &org2.ID,
			},
			allowed:  false,
			checkOrg: true,
			errorMsg: "invalid or unparsable field: Organization",
		},
		{
			name: "update status to invalid",
			updateInput: datumclient.UpdateUserSettingInput{
				Status: &enums.StatusInvalid,
			},
			checkOrg: false,
			errorMsg: "INVALID is not a valid UserSettingUserStatus",
		},
		{
			name: "update status to suspended",
			updateInput: datumclient.UpdateUserSettingInput{
				Status: &enums.Suspended,
			},
			checkOrg: false,
			expectedRes: datumclient.UpdateUserSetting_UpdateUserSetting_UserSetting{
				Status: enums.Suspended,
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// when attempting to update default org, we do a check
			if tc.checkOrg {
				mock_fga.CheckAny(t, suite.client.fga, tc.allowed)
			}

			// on successful update, we list the default org
			if tc.errorMsg == "" {
				mock_fga.ListAny(t, suite.client.fga, listObjects)
			}

			// update user
			resp, err := suite.client.datum.UpdateUserSetting(reqCtx, user1Setting.ID, tc.updateInput)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateUserSetting.UserSetting)

			// Make sure provided values match
			assert.Equal(t, tc.expectedRes.Status, resp.UpdateUserSetting.UserSetting.Status)
			assert.ElementsMatch(t, tc.expectedRes.Tags, resp.UpdateUserSetting.UserSetting.Tags)

			if tc.updateInput.DefaultOrgID != nil {
				assert.Equal(t, tc.expectedRes.DefaultOrg.ID, resp.UpdateUserSetting.UserSetting.DefaultOrg.ID)
			}
		})
	}
}
