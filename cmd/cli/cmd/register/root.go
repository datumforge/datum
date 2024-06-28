package register

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/models"
)

var cmd = &cobra.Command{
	Use:   "register",
	Short: "register a new datum user",
	Run: func(cmd *cobra.Command, args []string) {
		err := register(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	datum.RootCmd.AddCommand(cmd)

	cmd.Flags().StringP("email", "e", "", "email of the user")
	cmd.Flags().StringP("password", "p", "", "password of the user")
	cmd.Flags().StringP("first-name", "f", "", "first name of the user")
	cmd.Flags().StringP("last-name", "l", "", "last name of the user")
}

// validateRegister validates the required fields for the command
func validateRegister() (*models.RegisterRequest, error) {
	email := datum.Config.String("email")
	if email == "" {
		return nil, datum.NewRequiredFieldMissingError("email")
	}

	firstName := datum.Config.String("first-name")
	if firstName == "" {
		return nil, datum.NewRequiredFieldMissingError("first name")
	}

	lastName := datum.Config.String("last-name")
	if lastName == "" {
		return nil, datum.NewRequiredFieldMissingError("last name")
	}

	password := datum.Config.String("password")
	if password == "" {
		return nil, datum.NewRequiredFieldMissingError("password")
	}

	return &models.RegisterRequest{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}, nil
}

// register registers a new user in the datum platform
func register(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	cobra.CheckErr(err)

	var s []byte

	input, err := validateRegister()
	cobra.CheckErr(err)

	registration, err := client.Register(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(registration)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
