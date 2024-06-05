package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var subscribersUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update subscriber details",
	RunE: func(cmd *cobra.Command, args []string) error {
		return subscriberUpdate(cmd.Context())
	},
}

func init() {
	subscribersCmd.AddCommand(subscribersUpdateCmd)

	subscribersUpdateCmd.Flags().StringP("email", "e", "", "email address of the subscriber to update")
	datum.ViperBindFlag("subscribers.update.email", subscribersUpdateCmd.Flags().Lookup("email"))

	subscribersUpdateCmd.Flags().StringP("phonenumber", "p", "", "phone number to add or update on the subscriber")
	datum.ViperBindFlag("subscribers.update.phone", subscribersUpdateCmd.Flags().Lookup("phonenumber"))
}

func subscriberUpdate(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	email := viper.GetString("subscribers.update.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	phone := viper.GetString("subscribers.update.phone")

	input := datumclient.UpdateSubscriberInput{
		PhoneNumber: &phone,
	}

	var s []byte

	sub, err := client.UpdateSubscriber(ctx, email, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(sub)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
