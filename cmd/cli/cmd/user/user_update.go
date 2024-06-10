package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

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
	userUpdateCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	userUpdateCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	userUpdateCmd.Flags().StringP("display-name", "d", "", "display name of the user")
	userUpdateCmd.Flags().StringP("email", "e", "", "email of the user")
}

func updateUser(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	userID := datum.Config.String("id")
	if userID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	firstName := datum.Config.String("first-name")
	lastName := datum.Config.String("last-name")
	displayName := datum.Config.String("display-name")
	email := datum.Config.String("email")

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
