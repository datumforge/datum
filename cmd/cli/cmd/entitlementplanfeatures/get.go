package datumentitlementplanfeatures

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing datum plan features",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "feature id to query")
}

// get retrieves plan features from the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	// if an feature ID is provided, filter on that feature, otherwise get all
	if id == "" {
		// get all features, will be filtered for the authorized organization(s)
		out, err := client.GetAllEntitlementPlanFeatures(ctx)
		cobra.CheckErr(err)

		return consoleOutput(out)
	}

	out, err := client.GetEntitlementPlanFeatureByID(ctx, id)
	cobra.CheckErr(err)

	return consoleOutput(out)
}
