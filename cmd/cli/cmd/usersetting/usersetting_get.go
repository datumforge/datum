package datumusersetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var userSettingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get user settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return userSettings(cmd.Context())
	},
}

func init() {
	userSettingCmd.AddCommand(userSettingGetCmd)

	userSettingGetCmd.Flags().StringP("id", "i", "", "user setting id to retrieve")
}

func userSettings(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := datum.Config.String("id")

	var s []byte

	writer := tables.NewTableWriter(userSettingCmd.OutOrStdout(), "DefaultOrgName", "DefaultorgID", "2FA", "Status", "EmailConfirmed")

	// if setting ID is not provided, get settings which will automatically filter by user id
	if settingsID != "" {
		user, err := client.GetUserSettingByID(ctx, settingsID)
		cobra.CheckErr(err)

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(user.UserSetting)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(user.UserSetting.DefaultOrg.Name, user.UserSetting.DefaultOrg.ID, *user.UserSetting.IsTfaEnabled, user.UserSetting.Status, user.UserSetting.EmailConfirmed)
		writer.Render()

		return nil
	}

	settings, err := client.GetUserSettings(ctx)
	cobra.CheckErr(err)

	s, err = json.Marshal(settings.UserSettings)
	cobra.CheckErr(err)

	if datum.OutputFormat == datum.JSONOutput {
		return datum.JSONPrint(s)
	}

	for _, setting := range settings.UserSettings.Edges {
		writer.AddRow(setting.Node.DefaultOrg.Name, setting.Node.DefaultOrg.ID, *setting.Node.IsTfaEnabled, setting.Node.Status, setting.Node.EmailConfirmed)
	}

	writer.Render()

	return nil
}
