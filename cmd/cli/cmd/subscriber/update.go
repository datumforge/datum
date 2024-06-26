package datumsubscribers

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update subscriber details",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("email", "e", "", "email address of the subscriber to update")
	updateCmd.Flags().StringP("phone-number", "p", "", "phone number to add or update on the subscriber")
}

// updateValidation validates the input flags provided by the user
func updateValidation() (email string, input datumclient.UpdateSubscriberInput, err error) {
	email = datum.Config.String("email")
	if email == "" {
		return email, input, datum.NewRequiredFieldMissingError("email")
	}

	phone := datum.Config.String("phone-number")

	input.PhoneNumber = &phone

	return email, input, nil
}

// update a subscriber details
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	email, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateSubscriber(ctx, email, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
