package datumusersetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	datum.ViperBindFlag("usersetting.get.id", userSettingGetCmd.Flags().Lookup("id"))
}

func userSettings(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := viper.GetString("usersetting.get.id")

	var s []byte

	writer := tables.NewTableWriter(userSettingCmd.OutOrStdout(), "DefaultOrgName", "DefaultorgID", "2FA", "Status", "EmailConfirmed")

	// if setting ID is not provided, get settings which will automatically filter by user id
	if settingsID != "" {
		user, err := cli.Client.GetUserSettingByID(ctx, settingsID, cli.Interceptor)
		if err != nil {
			return err
		}

		if datum.OutputFormat == "json" {
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

	settings, err := cli.Client.GetUserSettings(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(settings.UserSettings)
	if err != nil {
		return err
	}

	if datum.OutputFormat == "json" {
		return datum.JSONPrint(s)
	}

	for _, setting := range settings.UserSettings.Edges {
		writer.AddRow(setting.Node.DefaultOrg.Name, setting.Node.DefaultOrg.ID, *setting.Node.IsTfaEnabled, setting.Node.Status, setting.Node.EmailConfirmed)
	}

	writer.Render()

	return nil
}
