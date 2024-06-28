package datumentitlementplan

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing datum entitlement plans",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "entitlement plan id to query")
}

// get retrieves plans from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	// if an plan ID is provided, filter on that plan, otherwise get all
	if id == "" {
		// get all plans, will be filtered for the authorized organization(s)
		p, err := client.GetAllEntitlementPlans(ctx)
		cobra.CheckErr(err)

		return consoleOutput(p)
	}

	p, err := client.GetEntitlementPlanByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(p)
}
