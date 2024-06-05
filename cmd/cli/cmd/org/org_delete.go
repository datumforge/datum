package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	datum.ViperBindFlag("org.delete.id", orgDeleteCmd.Flags().Lookup("id"))
}

func deleteOrg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	oID := viper.GetString("org.delete.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	o, err := client.DeleteOrganization(ctx, oID, client.Config().Interceptors...)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
