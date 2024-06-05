package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateUser(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().StringP("id", "i", "", "user id to update")
	datum.ViperBindFlag("user.update.id", userUpdateCmd.Flags().Lookup("id"))

	userUpdateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	datum.ViperBindFlag("user.update.first-name", userUpdateCmd.Flags().Lookup("first-name"))

	userUpdateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	datum.ViperBindFlag("user.update.last-name", userUpdateCmd.Flags().Lookup("last-name"))

	userUpdateCmd.Flags().StringP("display-name", "d", "", "display name of the user")
	datum.ViperBindFlag("user.update.display-name", userUpdateCmd.Flags().Lookup("display-name"))

	userUpdateCmd.Flags().StringP("email", "e", "", "email of the user")
	datum.ViperBindFlag("user.update.email", userUpdateCmd.Flags().Lookup("email"))
}

func updateUser(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	userID := viper.GetString("user.update.id")
	if userID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	firstName := viper.GetString("user.update.first-name")
	lastName := viper.GetString("user.update.last-name")
	displayName := viper.GetString("user.update.display-name")
	email := viper.GetString("user.update.email")

	input := datumclient.UpdateUserInput{}

	if firstName != "" {
		input.FirstName = &firstName
	}

	if lastName != "" {
		input.LastName = &lastName
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if email != "" {
		input.Email = &email
	}

	// TODO: allow updates to user settings

	o, err := client.UpdateUser(ctx, userID, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
