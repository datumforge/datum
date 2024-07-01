package datumorg

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get details of existing datum organization(s)",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "get a specific organization by ID")
}

// get an organization in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	// if an org ID is provided, filter on that organization, otherwise get all
	if id != "" {
		o, err := client.GetOrganizationByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetAllOrganizations(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
