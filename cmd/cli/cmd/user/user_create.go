package datumuser

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	datum.ViperBindFlag("user.create.email", userCreateCmd.Flags().Lookup("email"))

	userCreateCmd.Flags().StringP("password", "p", "", "password of the user")
	datum.ViperBindFlag("user.create.password", userCreateCmd.Flags().Lookup("password"))

	userCreateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	datum.ViperBindFlag("user.create.first-name", userCreateCmd.Flags().Lookup("first-name"))

	userCreateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	datum.ViperBindFlag("user.create.last-name", userCreateCmd.Flags().Lookup("last-name"))

	userCreateCmd.Flags().StringP("display-name", "d", "", "first name of the user")
	datum.ViperBindFlag("user.create.display-name", userCreateCmd.Flags().Lookup("display-name"))
}

func createUser(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	email := viper.GetString("user.create.email")
	if email == "" {
		return datum.NewRequiredFieldMissingError("email")
	}

	firstName := viper.GetString("user.create.first-name")
	if firstName == "" {
		return datum.NewRequiredFieldMissingError("first name")
	}

	lastName := viper.GetString("user.create.last-name")
	if lastName == "" {
		return datum.NewRequiredFieldMissingError("last name")
	}

	displayName := viper.GetString("user.create.display-name")
	if displayName == "" {
		// set a default display name if not set
		displayName = strings.ToLower(fmt.Sprintf("%s.%s", firstName, lastName))
	}

	password := viper.GetString("user.create.password")

	input := datumclient.CreateUserInput{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if password != "" {
		input.Password = &password
	}

	u, err := cli.Client.CreateUser(ctx, input, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(u)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
