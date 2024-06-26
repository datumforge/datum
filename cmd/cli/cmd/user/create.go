package datumuser

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new datum user",
	Run: func(cmd *cobra.Command, args []string) {
		err := create(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("email", "e", "", "email of the user")
	createCmd.Flags().StringP("password", "p", "", "password of the user")
	createCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	createCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	createCmd.Flags().StringP("display-name", "d", "", "first name of the user")
}

// createValidation validates the required fields for the command
func createValidation() (input datumclient.CreateUserInput, err error) {
	input.Email = datum.Config.String("email")
	if input.Email == "" {
		return input, datum.NewRequiredFieldMissingError("email")
	}

	firstName := datum.Config.String("first-name")
	if firstName == "" {
		return input, datum.NewRequiredFieldMissingError("first name")
	}

	input.FirstName = &firstName

	lastName := datum.Config.String("last-name")
	if lastName == "" {
		return input, datum.NewRequiredFieldMissingError("last name")
	}

	input.LastName = &lastName

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = displayName
	}

	password := datum.Config.String("password")
	if password != "" {
		input.Password = &password
	}

	return input, nil
}

// create a new datum user
func create(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	input, err := createValidation()
	cobra.CheckErr(err)

	o, err := client.CreateUser(ctx, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
