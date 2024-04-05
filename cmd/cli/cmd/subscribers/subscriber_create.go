package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
)

var subscribersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add subscribers to a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscriberCreate(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersCreateCmd)

	subscribersCreateCmd.Flags().StringP("email", "e", "", "email address of the subscriber")
	datum.ViperBindFlag("subscribers.create.email", subscribersCreateCmd.Flags().Lookup("email"))
}

func subscriberCreate(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	email := viper.GetString("subscribers.create.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	input := datumclient.CreateSubscriberInput{
		Email: email,
	}

	var s []byte

	sub, err := cli.Client.CreateSubscriber(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(sub)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
