package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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

	subscribersGetCmd.Flags().StringP("org-id", "o", "", "org id to query")
	datum.ViperBindFlag("subscribers.get.orgid", subscribersGetCmd.Flags().Lookup("org-id"))

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

	oID := viper.GetString("subscribers.get.orgid")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("org id")
	}

	// filter options
	where := datumclient.SubscriberWhereInput{
		OwnerID: &oID,
	}

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
