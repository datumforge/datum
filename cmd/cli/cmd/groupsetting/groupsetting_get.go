package datumgroupsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var groupSettingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get group settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return groupSettings(cmd.Context())
	},
}

func init() {
	groupSettingCmd.AddCommand(groupSettingGetCmd)

	groupSettingGetCmd.Flags().StringP("id", "i", "", "group setting id to retrieve")
	datum.ViperBindFlag("groupsetting.get.id", groupSettingGetCmd.Flags().Lookup("id"))
}

func groupSettings(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := viper.GetString("groupsetting.get.id")

	var s []byte

	// if setting ID is not provided, get settings which will automatically filter by group id
	if settingsID == "" {
		settings, err := cli.Client.GetGroupSettings(ctx, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(settings)
		if err != nil {
			return err
		}
	} else {
		group, err := cli.Client.GetGroupSettingByID(ctx, settingsID, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(group)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
