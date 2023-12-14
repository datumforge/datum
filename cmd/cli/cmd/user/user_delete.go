package datumuser

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteUser(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)

	userDeleteCmd.Flags().StringP("id", "i", "", "user id to delete")
	datum.ViperBindFlag("user.delete.id", userDeleteCmd.Flags().Lookup("id"))
}

func deleteUser(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	var s []byte

	userID := viper.GetString("user.delete.id")
	if userID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	o, err := cli.Client.DeleteUser(ctx, userID, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
