package datuminvite

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get organization invitation(s)",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().StringP("id", "i", "", "invite id to query")
}

// get an organization invitation(s)
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	// filter options
	id := datum.Config.String("id")

	if id != "" {
		o, err := client.GetInviteByID(ctx, id)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	o, err := client.GetAllInvites(ctx)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
