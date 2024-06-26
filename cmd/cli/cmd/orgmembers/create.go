package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "add user to a datum organization",
	Run: func(cmd *cobra.Command, args []string) {
		err := addOrgMember(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	cmd.AddCommand(createCmd)

	createCmd.Flags().StringP("org-id", "o", "", "organization id")
	createCmd.Flags().StringP("user-id", "u", "", "user id")
	createCmd.Flags().StringP("role", "r", "member", "role to assign the user (member, admin)")
}

func addOrgMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	cobra.CheckErr(err)
	defer datum.StoreSessionCookies(client)

	oID := datum.Config.String("org-id")

	uID := datum.Config.String("user-id")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// role defaults to `member` so it is not required
	role := datum.Config.String("role")

	r, err := datum.GetRoleEnum(role)
	cobra.CheckErr(err)

	input := datumclient.CreateOrgMembershipInput{
		UserID: uID,
		Role:   &r,
	}

	if oID != "" {
		input.OrganizationID = oID
	}

	var s []byte

	orgMember, err := client.AddUserToOrgWithRole(ctx, input)
	cobra.CheckErr(err)

	s, err = json.Marshal(orgMember)
	cobra.CheckErr(err)

	return datum.JSONPrint(s)
}
