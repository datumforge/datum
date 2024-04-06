package datumuser

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
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
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

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
