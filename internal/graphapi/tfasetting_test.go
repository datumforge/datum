package graphapi_test

import (
	"testing"

	mock_fga "github.com/datumforge/fgax/mockery"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/datumclient"
)

func TestQueryTFASetting(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user1 := (&UserBuilder{client: client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user1.ID)
	require.NoError(t, err)

	(&TFASettingBuilder{client: client}).MustNew(reqCtx, t, user1.ID)

	user2 := (&UserBuilder{client: client}).MustNew(ctx, t)

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
			errorMsg: "tfa_settings not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			reqCtx, err := userContextWithID(tc.userID)
			require.NoError(t, err)

			resp, err := client.datum.GetTFASettings(reqCtx)

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

	(&UserCleanup{client: client, UserID: user1.ID}).MustDelete(reqCtx, t)
	(&UserCleanup{client: client, UserID: user2.ID}).MustDelete(reqCtx, t)
}

func TestMutationCreateTFASetting(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user := (&UserBuilder{client: client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		input  datumclient.CreateTFASettingsInput
		errMsg string
	}{
		{
			name: "happy path",
			input: datumclient.CreateTFASettingsInput{
				TotpAllowed: lo.ToPtr(true),
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Create "+tc.name, func(t *testing.T) {
			// update user
			resp, err := client.datum.CreateTFASettings(reqCtx, tc.input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.CreateTFASettings.TfaSettings)

			// Make sure provided values match
			assert.Equal(t, tc.input.TotpAllowed, resp.CreateTFASettings.TfaSettings.TotpAllowed)
			assert.Empty(t, resp.CreateTFASettings.TfaSettings.RecoveryCodes)

			// make sure user setting was not updated
			userSetting, err := user.Setting(ctx)
			require.NoError(t, err)

			assert.False(t, userSetting.IsTfaEnabled)
		})
	}
}

func TestMutationUpdateTFASetting(t *testing.T) {
	client := setupTest(t)
	defer client.db.Close()

	// setup user context
	ctx, err := userContext()
	require.NoError(t, err)

	user := (&UserBuilder{client: client}).MustNew(ctx, t)
	require.NoError(t, err)

	// setup valid user context
	reqCtx, err := userContextWithID(user.ID)
	require.NoError(t, err)

	(&TFASettingBuilder{client: client}).MustNew(reqCtx, t, user.ID)

	recoveryCodes := []string{}

	testCases := []struct {
		name   string
		input  datumclient.UpdateTFASettingsInput
		errMsg string
	}{
		{
			name: "update verify",
			input: datumclient.UpdateTFASettingsInput{
				Verified: lo.ToPtr(true),
			},
		},
		{
			name: "regen codes",
			input: datumclient.UpdateTFASettingsInput{
				RegenBackupCodes: lo.ToPtr(true),
			},
		},
		{
			name: "regen codes - false",
			input: datumclient.UpdateTFASettingsInput{
				RegenBackupCodes: lo.ToPtr(false),
			},
		},
	}

	for _, tc := range testCases {
		t.Run("Update "+tc.name, func(t *testing.T) {
			defer mock_fga.ClearMocks(client.fga)

			// update user
			resp, err := client.datum.UpdateTFASettings(reqCtx, tc.input)

			if tc.errMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.UpdateTFASettings.TfaSettings)

			// Make sure provided values match
			assert.NotEmpty(t, resp.UpdateTFASettings.TfaSettings.RecoveryCodes)

			// backup codes should only be regenerated on explicit request
			if tc.input.RegenBackupCodes != nil {
				if *tc.input.RegenBackupCodes {
					assert.NotEqual(t, recoveryCodes, resp.UpdateTFASettings.TfaSettings.RecoveryCodes)
				} else {
					assert.Equal(t, recoveryCodes, resp.UpdateTFASettings.TfaSettings.RecoveryCodes)
				}
			}

			// make sure user setting was not updated
			// list objects is called to get the default org, but we don't care about that here
			listObjects := []string{"organization:test"}
			mock_fga.ListAny(t, client.fga, listObjects)

			userSettings, err := client.datum.GetUserSettings(reqCtx)
			require.NoError(t, err)
			require.Len(t, userSettings.UserSettings.Edges, 1)

			if resp.UpdateTFASettings.TfaSettings.Verified {
				assert.True(t, *userSettings.UserSettings.Edges[0].Node.IsTfaEnabled)
			}

			// set at the end so we can compare later
			recoveryCodes = resp.UpdateTFASettings.TfaSettings.RecoveryCodes
		})
	}
}
