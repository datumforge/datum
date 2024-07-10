package datumcontact

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum contact",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	// command line flags for the create command
	createCmd.Flags().StringP("name", "n", "", "full name of the contact")
	createCmd.Flags().StringP("email", "e", "", "email address of the contact")
	createCmd.Flags().StringP("phone", "p", "", "phone number of the contact")
	createCmd.Flags().StringP("title", "t", "", "title of the contact")
	createCmd.Flags().StringP("company", "c", "", "company of the contact")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateContactInput, err error) {
	// validation of required fields for the create command
	input.FullName = datum.Config.String("name")
	if input.FullName == "" {
		return input, datum.NewRequiredFieldMissingError("contact name")
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

	return input, nil
}

// create a new datum contact
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateContact(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
