package datumtokens

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
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
}

func deletePat(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	oID := datum.Config.String("id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("token id")
	}

	o, err := client.DeletePersonalAccessToken(ctx, oID)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
