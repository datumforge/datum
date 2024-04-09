package graphapi_test

import (
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/pkg/datumclient"
)

func (suite *GraphTestSuite) TestQueryTFASetting() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user1.ID)
	require.NoError(t, err)

	(&TFASettingBuilder{client: suite.client}).MustNew(reqCtx, t, user1.ID)

	user2 := (&UserBuilder{client: suite.client}).MustNew(ctx, t)

	testCases := []struct {
		name     string
		userID   string
		errorMsg string
	}{
		{
			name:   "happy path user",
			userID: user1.ID,
		},
		{
			name:     "valid user, but not auth",
			userID:   user2.ID,
			errorMsg: "tfa_setting not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			reqCtx, err := userContextWithID(tc.userID)
			require.NoError(t, err)

			resp, err := suite.client.datum.GetTFASetting(reqCtx)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
		})
	}

	(&UserCleanup{client: suite.client, UserID: user1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: suite.client, UserID: user2.ID}).MustDelete(reqCtx, t)
}

func (suite *GraphTestSuite) TestMutationCreateTFASetting() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		input  datumclient.CreateTFASettingInput
		errMsg string
	}{
		{
			name: "happy path",
			input: datumclient.CreateTFASettingInput{
				TotpAllowed: lo.ToPtr(true),
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			// update user
			resp, err := suite.client.datum.CreateTFASetting(reqCtx, tc.input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateTFASetting.TfaSetting)

			// Make sure provided values match
			assert.Equal(t, tc.input.TotpAllowed, resp.CreateTFASetting.TfaSetting.TotpAllowed)
			assert.Empty(t, resp.CreateTFASetting.TfaSetting.RecoveryCodes)

			// make sure user setting was not updated
			userSetting, err := user.Setting(ctx)
			require.NoError(t, err)

			assert.False(t, userSetting.IsTfaEnabled)
		})
	}
}

func (suite *GraphTestSuite) TestMutationUpdateTFASetting() {
	t := suite.T()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user := (&UserBuilder{client: suite.client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	(&TFASettingBuilder{client: suite.client}).MustNew(reqCtx, t, user.ID)

	recoveryCodes := []string{}

	testCases := []struct {
		name   string
		input  datumclient.UpdateTFASettingInput
		errMsg string
	}{
		{
			name: "update verify",
			input: datumclient.UpdateTFASettingInput{
				Verified: lo.ToPtr(true),
			},
		},
		{
			name: "regen codes",
			input: datumclient.UpdateTFASettingInput{
				RegenBackupCodes: lo.ToPtr(true),
			},
		},
		{
			name: "regen codes - false",
			input: datumclient.UpdateTFASettingInput{
				RegenBackupCodes: lo.ToPtr(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(suite.client.fga)

			// update user
			resp, err := suite.client.datum.UpdateTFASetting(reqCtx, tc.input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateTFASetting.TfaSetting)

			// Make sure provided values match
			assert.NotEmpty(t, resp.UpdateTFASetting.TfaSetting.RecoveryCodes)

			// backup codes should only be regenerated on explicit request
			if tc.input.RegenBackupCodes != nil {
				if *tc.input.RegenBackupCodes {
					assert.NotEqual(t, recoveryCodes, resp.UpdateTFASetting.TfaSetting.RecoveryCodes)
				} else {
					assert.Equal(t, recoveryCodes, resp.UpdateTFASetting.TfaSetting.RecoveryCodes)
				}
			}

			// make sure user setting was not updated
			// list objects is called to get the default org, but we don't care about that here
			listObjects := []string{"organization:test"}
			mock_fga.ListAny(t, suite.client.fga, listObjects)

			userSettings, err := suite.client.datum.GetUserSettings(reqCtx)
			require.NoError(t, err)
			require.Len(t, userSettings.UserSettings.Edges, 1)

			if resp.UpdateTFASetting.TfaSetting.Verified {
				assert.True(t, *userSettings.UserSettings.Edges[0].Node.IsTfaEnabled)
			}

			// set at the end so we can compare later
			recoveryCodes = resp.UpdateTFASetting.TfaSetting.RecoveryCodes
		})
	}
}
