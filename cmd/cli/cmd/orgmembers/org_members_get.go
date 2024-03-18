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
	cli, err := datum.GetGraphClient(ctx)
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
		OrganizationID: &oID,
	}

	var s []byte

	org, err := cli.Client.GetOrgMembersByOrgID(ctx, &where, cli.Interceptor)
	if err != nil {
		return err
	}

	if viper.GetString("output.format") == "json" {
		s, err = json.Marshal(org)
		if err != nil {
			return err
		}

		return datum.JSONPrint(s)
	}

	return orgMembersTablePrint(*org)
}

func orgMembersTablePrint(om datumclient.GetOrgMembersByOrgID) error {
	// check if there are any org members, otherwise we have nothing to print
	if len(om.OrgMemberships.Edges) > 0 {
		// get the headers for the table for each struct and substruct
		header := datum.GetHeaders(om.OrgMemberships.Edges[0].Node, "")
		subHeaders := datum.GetHeaders(om.OrgMemberships.Edges[0].Node.User, "User.")

		// combine the headers
		header = append(header, subHeaders...)

		data := [][]string{}

		// get the field values for each struct and substruct per row
		for _, v := range om.OrgMemberships.Edges {
			fields := datum.GetFields(*v.Node)
			subfields := datum.GetFields(v.Node.User)

			// combine the fields
			fields = append(fields, subfields...)

			// append the fields to the data slice
			data = append(data, fields)
		}

		// print ze data
		return datum.TablePrint(header, data)
	}

	return nil
}
