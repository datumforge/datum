package datumsubscribers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	subscribersUpdateCmd.Flags().StringP("phone-number", "p", "", "phone number to add or update on the subscriber")
}

func subscriberUpdate(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	email := datum.Config.String("email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	phone := datum.Config.String("phone-number")

	input := datumclient.UpdateSubscriberInput{
		PhoneNumber: &phone,
	}

	var s []byte

	sub, err := client.UpdateSubscriber(ctx, email, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(sub)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
