package datumuserhistory

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing datum userHistory",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "userHistory id to query")
}

// get an existing userHistory in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// get all will be filtered for the authorized organization(s)
	o, err := client.GetAllUserHistories(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}