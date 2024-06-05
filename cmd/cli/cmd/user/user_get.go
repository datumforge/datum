package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return users(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	userGetCmd.Flags().StringP("id", "i", "", "user id to query")
	datum.ViperBindFlag("user.get.id", userGetCmd.Flags().Lookup("id"))
}

func users(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	// filter options
	userID := viper.GetString("user.get.id")

	var s []byte

	writer := tables.NewTableWriter(userCmd.OutOrStdout(), "ID", "Email", "FirstName", "LastName", "AuthProvider")

	// if a user ID is provided, filter on that user, otherwise get all
	if userID != "" {
		user, err := client.GetUserByID(ctx, userID, client.Config().Interceptors...)
		if err != nil {
			return err
		}

		if datum.OutputFormat == datum.JSONOutput {
			s, err := json.Marshal(user.User)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(user.User.ID, user.User.Email, *user.User.FirstName, *user.User.LastName, user.User.AuthProvider)

		writer.Render()

		return nil
	}

	users, err := client.GetAllUsers(ctx, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	s, err = json.Marshal(users)
	if err != nil {
		return err
	}

	if datum.OutputFormat == datum.JSONOutput {
		return datum.JSONPrint(s)
	}

	for _, u := range users.Users.Edges {
		writer.AddRow(u.Node.ID, u.Node.Email, *u.Node.FirstName, *u.Node.LastName, u.Node.AuthProvider)
	}

	writer.Render()

	return nil
}
