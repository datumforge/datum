package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgMembersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add user to a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return addOrgMember(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersCreateCmd)

	orgMembersCreateCmd.Flags().StringP("org-id", "o", "", "organization id")
	orgMembersCreateCmd.Flags().StringP("user-id", "u", "", "user id")
	orgMembersCreateCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

func addOrgMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	oID := datum.Config.String("org-id")

	uID := datum.Config.String("user-id")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := datum.Config.String("role")

	r, err := datum.GetRoleEnum(role)
	if err != nil {
		return err
	}

	input := datumclient.CreateOrgMembershipInput{
		UserID: uID,
		Role:   &r,
	}

	if oID != "" {
		input.OrganizationID = oID
	}

	var s []byte

	orgMember, err := client.AddUserToOrgWithRole(ctx, input)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
