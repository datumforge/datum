package datumgroupsetting

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
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
}

func groupSettings(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	settingsID := datum.Config.String("id")

	var s []byte

	// if setting ID is not provided, get settings which will automatically filter by group id
	if settingsID == "" {
		settings, err := client.GetGroupSettings(ctx)
		if err != nil {
			return err
		}

		s, err = json.Marshal(settings)
		if err != nil {
			return err
		}
	} else {
		group, err := client.GetGroupSettingByID(ctx, settingsID)
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
