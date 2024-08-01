package datumorgsetting

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get organization settings",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "organization setting id to retrieve")
}

// get an existing organization setting in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	// if setting ID is not provided, get settings which will automatically filter by org id
	if id == "" {
		o, err := client.GetAllOrganizationSettings(ctx)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetOrganizationSettingByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
