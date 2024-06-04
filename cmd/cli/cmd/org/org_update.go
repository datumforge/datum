package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgUpdateCmd)

	orgUpdateCmd.Flags().StringP("id", "i", "", "org id to update")
	orgUpdateCmd.Flags().StringP("name", "n", "", "name of the organization")
	orgUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	orgUpdateCmd.Flags().StringP("description", "d", "", "description of the organization")
}

func updateOrg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	oID := datum.Config.String("id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	name := datum.Config.String("name")
	displayName := datum.Config.String("short-name")
	description := datum.Config.String("description")

	input := datumclient.UpdateOrganizationInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := client.UpdateOrganization(ctx, oID, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
