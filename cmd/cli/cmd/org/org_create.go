package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgCreateCmd)

	orgCreateCmd.Flags().StringP("name", "n", "", "name of the organization")
	orgCreateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	orgCreateCmd.Flags().StringP("description", "d", "", "description of the organization")
	orgCreateCmd.Flags().StringP("parent-org-id", "p", "", "parent organization id, leave empty to create a root org")

	// TODO: https://github.com/datumforge/datum/issues/734
	// remove flag once the feature is implemented
	orgCreateCmd.Flags().BoolP("dedicated-db", "D", false, "create a dedicated database for the organization")
}

func createOrg(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	var s []byte

	name := datum.Config.String("name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("organization name")
	}

	displayName := datum.Config.String("short-name")
	description := datum.Config.String("description")
	parentOrgID := datum.Config.String("parent-org-id")
	dedicatedDB := datum.Config.Bool("dedicated-db")

	input := datumclient.CreateOrganizationInput{
		Name: name,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	if parentOrgID != "" {
		input.ParentID = &parentOrgID
	}

	if dedicatedDB {
		input.DedicatedDb = &dedicatedDB
	}

	o, err := client.CreateOrganization(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
