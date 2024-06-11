package datumgroup

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var groupDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupDeleteCmd)

	groupDeleteCmd.Flags().StringP("id", "i", "", "group id to delete")
}

func deleteGroup(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	gID := datum.Config.String("id")
	if gID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	o, err := client.DeleteGroup(ctx, gID)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
