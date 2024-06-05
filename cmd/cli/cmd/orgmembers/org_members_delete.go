package datumorgmembers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
)

var orgMembersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a user from a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteOrgMember(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersDeleteCmd)

	orgMembersDeleteCmd.Flags().StringP("user-id", "u", "", "user id")
	datum.ViperBindFlag("orgmember.delete.userid", orgMembersDeleteCmd.Flags().Lookup("user-id"))
}

func deleteOrgMember(ctx context.Context) error {
	// setup datum http client
	client, err := datum.SetupClientWithAuth(ctx)
	if err != nil {
		return err
	}
	defer datum.StoreSessionCookies(client)

	oID := viper.GetString("orgmember.delete.orgid")

	uID := viper.GetString("orgmember.delete.userid")
	if uID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	// get the id of the org member
	where := datumclient.OrgMembershipWhereInput{
		UserID: &uID,
	}

	if oID != "" {
		where.OrganizationID = &oID
	}

	orgMembers, err := client.GetOrgMembersByOrgID(ctx, &where)
	if err != nil {
		return err
	}

	if len(orgMembers.OrgMemberships.Edges) != 1 {
		return errors.New("error getting existing relation") //nolint:goerr113
	}

	id := orgMembers.OrgMemberships.Edges[0].Node.ID

	var s []byte

	orgMember, err := client.RemoveUserFromOrg(ctx, id)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgMember)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
