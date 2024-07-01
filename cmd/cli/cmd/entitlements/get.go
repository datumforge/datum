package datumentitlement

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing datum entitlements",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "entitlement id to query")
}

// get retrieves entitlements from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	// if an entitlement ID is provided, filter on that entitlement, otherwise get all
	if id == "" {
		// get all entitlements, will be filtered for the authorized organization(s)
		e, err := client.GetEntitlements(ctx, nil)
		cobra.CheckErr(err)

		return consoleOutput(e)
	}

	e, err := client.GetEntitlementByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(e)
}
