package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
)

var orgDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgDeleteCmd)

	orgDeleteCmd.Flags().StringP("id", "i", "", "org id to delete")
}

func deleteOrg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	var s []byte

	oID := datum.Config.String("id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	o, err := client.DeleteOrganization(ctx, oID)
	cobra.CheckErr(err)

	s, err = json.Marshal(o)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
