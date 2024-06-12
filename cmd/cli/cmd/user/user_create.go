package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createUser(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().StringP("email", "e", "", "email of the user")
	userCreateCmd.Flags().StringP("password", "p", "", "password of the user")
	userCreateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	userCreateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	userCreateCmd.Flags().StringP("display-name", "d", "", "first name of the user")
}

func createUser(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	email := datum.Config.String("email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	firstName := datum.Config.String("first-name")
	if firstName == "" {
		return datum.NewRequiredFieldMissingError("first name")
	}

	lastName := datum.Config.String("last-name")
	if lastName == "" {
		return datum.NewRequiredFieldMissingError("last name")
	}

	displayName := datum.Config.String("display-name")

	password := datum.Config.String("password")

	input := datumclient.CreateUserInput{
		Email:     email,
		FirstName: &firstName,
		LastName:  &lastName,
	}

	// if a displayName is not provided, the email is used
	if displayName != "" {
		input.DisplayName = displayName
	}

	if password != "" {
		input.Password = &password
	}

	u, err := client.CreateUser(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(u)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
