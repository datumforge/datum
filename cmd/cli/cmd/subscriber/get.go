package datumsubscribers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get existing subscribers of a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := get(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(getCmd)

	getCmd.Flags().BoolP("active", "a", true, "filter on active subscribers")
	getCmd.Flags().StringP("email", "e", "", "email address of the subscriber to get")
}

// get an existing subscriber(s) of an organization
func get(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	email := datum.Config.String("email")

	if email != "" {
		o, err := client.GetSubscriberByEmail(ctx, email)
		cobra.CheckErr(err)

		return consoleOutput(o)
	}

	// filter options
	where := datumclient.SubscriberWhereInput{}

	active := datum.Config.Bool("active")
	if active {
		where.Active = &active
	}

	o, err := client.GetSubscribers(ctx, &where)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
