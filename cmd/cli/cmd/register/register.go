package register

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return registerUser(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("email", "e", "", "email of the user")
	registerCmd.Flags().StringP("password", "p", "", "password of the user")
	registerCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	registerCmd.Flags().StringP("last-name", "l", "", "last name of the user")
}

func registerUser(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	if err != nil {
		return err
	}

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

	password := datum.Config.String("password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	register := models.RegisterRequest{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}

	registration, err := client.Register(ctx, &register)
	if err != nil {
		return err
	}

	s, err = json.Marshal(registration)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
