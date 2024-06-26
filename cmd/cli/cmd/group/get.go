package datumgroup

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing new datum group",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "group id to query")
}

// get an existing group in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	// if an group ID is provided, filter on that group, otherwise get all
	if id != "" {
		o, err := client.GetGroupByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	// get all o, will be filtered for the authorized organization(s)
	o, err := client.GetAllGroups(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
