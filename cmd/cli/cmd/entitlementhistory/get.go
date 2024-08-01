package datumentitlementhistory

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get an existing datum entitlementHistory",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)
	getCmd.Flags().StringP("id", "i", "", "id to query")
}

// get an existing entitlementHistory in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")
	if id != "" {
		o, err := client.GetEntitlementHistories(ctx, &datumclient.EntitlementHistoryWhereInput{
			Ref: &id,
		})
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	// get all will be filtered for the authorized organization(s)
	o, err := client.GetAllEntitlementHistories(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
