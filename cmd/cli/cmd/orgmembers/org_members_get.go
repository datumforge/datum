package datumorgmembers

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
)

var orgMembersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing members of a datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return orgMembers(cmd.Context())
	},
}

func init() {
	orgMembersCmd.AddCommand(orgMembersGetCmd)

	orgMembersGetCmd.Flags().StringP("org-id", "o", "", "org id to query")
	datum.ViperBindFlag("orgmember.get.id", orgMembersGetCmd.Flags().Lookup("org-id"))
}

func orgMembers(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	oID := viper.GetString("orgmember.get.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	where := datumclient.OrgMembershipWhereInput{
		OrgID: &oID,
	}

	var s []byte

	org, err := cli.Client.GetOrgMembersByOrgID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(org)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
