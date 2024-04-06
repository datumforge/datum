package datumusersetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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

	// if setting ID is not provided, get settings which will automatically filter by user id
	if settingsID == "" {
		settings, err := cli.Client.GetUserSettings(ctx, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(settings)
		if err != nil {
			return err
		}
	} else {
		user, err := cli.Client.GetUserSettingByID(ctx, settingsID, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(user)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
