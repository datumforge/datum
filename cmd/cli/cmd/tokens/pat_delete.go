package datumtokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
)

var patDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum personal access token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deletePat(cmd.Context())
	},
}

func init() {
	patCmd.AddCommand(patDeleteCmd)

	patDeleteCmd.Flags().StringP("id", "i", "", "pat id to delete")
	datum.ViperBindFlag("pat.delete.id", patDeleteCmd.Flags().Lookup("id"))
}

func deletePat(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	var s []byte

	oID := viper.GetString("pat.delete.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	o, err := cli.Client.DeletePersonalAccessToken(ctx, oID, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
