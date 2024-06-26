package datumusersetting

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get user settings",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "user setting id to retrieve")
}

// get retrieves user settings from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	// if setting ID is not provided, get settings which will automatically filter by user id
	if id != "" {
		o, err := client.GetUserSettingByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetAllUserSettings(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
