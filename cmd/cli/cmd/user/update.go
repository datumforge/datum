package datumuser

import (
	"context"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing datum user",
	Run: func(cmd *cobra.Command, args []string) {
		err := update(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("id", "i", "", "user id to update")
	updateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	updateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	updateCmd.Flags().StringP("display-name", "d", "", "display name of the user")
	updateCmd.Flags().StringP("email", "e", "", "email of the user")
}

// updateValidation validates the input flags provided by the user
func updateValidation() (id string, input datumclient.UpdateUserInput, err error) {
	id = datum.Config.String("id")
	if id == "" {
		return id, input, datum.NewRequiredFieldMissingError("user id")
	}

	firstName := datum.Config.String("first-name")
	if firstName != "" {
		input.FirstName = &firstName
	}

	lastName := datum.Config.String("last-name")
	if lastName != "" {
		input.LastName = &lastName
	}

	displayName := datum.Config.String("display-name")
	if displayName != "" {
		input.DisplayName = &displayName
	}

	email := datum.Config.String("email")
	if email != "" {
		input.Email = &email
	}

	// TODO: allow updates to user settings
	return id, input, nil
}

// update an existing datum user
func update(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	id, input, err := updateValidation()
	cobra.CheckErr(err)

	o, err := client.UpdateUser(ctx, id, input)
	cobra.CheckErr(err)

	return consoleOutput(o)
}
