package datumfeature

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing datum features",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "feature id to query")
}

// get an existing feature in the datum platform
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id := datum.Config.String("id")

	if id != "" {
		f, err := client.GetFeatureByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(f)
	}

	// get all features, will be filtered for the authorized organization(s)
	features, err := client.GetFeatures(ctx, nil)
	cobra.CheckErr(err)

	return consoleOutput(features)
}
