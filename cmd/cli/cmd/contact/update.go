package datumcontact

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum contact",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "contact id to update")

	// command line flags for the update command
	updateCmd.Flags().StringP("name", "n", "", "full name of the contact")
	updateCmd.Flags().StringP("email", "e", "", "email address of the contact")
	updateCmd.Flags().StringP("phone", "p", "", "phone number of the contact")
	updateCmd.Flags().StringP("title", "t", "", "title of the contact")
	updateCmd.Flags().StringP("company", "c", "", "company of the contact")
}

// updateValidation validates the required fields for the command
func updateValidation() (id string, input datumclient.UpdateContactInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("contact id")
	}

	// validation of required fields for the update command
	name := datum.Config.String("name")
	if name != "" {
		input.FullName = &name
	}

	email := datum.Config.String("email")
	if email != "" {
		input.Email = &email
	}

	phone := datum.Config.String("phone")
	if phone != "" {
		input.PhoneNumber = &phone
	}

	title := datum.Config.String("title")
	if title != "" {
		input.Title = &title
	}

	company := datum.Config.String("company")
	if company != "" {
		input.Company = &company
	}

	return id, input, nil
}

// update an existing contact in the datum platform
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateContact(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
