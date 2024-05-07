package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var subscribersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing subscribers of a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscribers(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersGetCmd)

	subscribersGetCmd.Flags().BoolP("active", "a", true, "filter on active subscribers")
	datum.ViperBindFlag("subscribers.get.active", subscribersGetCmd.Flags().Lookup("active"))
}

func subscribers(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	where := datumclient.SubscriberWhereInput{}

	active := viper.GetBool("subscribers.get.active")
	if active {
		where.Active = &active
	}

	var s []byte

	subs, err := cli.Client.Subscribers(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(subs)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}