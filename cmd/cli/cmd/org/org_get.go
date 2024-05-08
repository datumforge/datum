package datumorg

import (
	"context"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/utils/cli/tables"
)

var orgGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of existing datum orgs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return orgs(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgGetCmd)

	orgGetCmd.Flags().StringP("id", "i", "", "get a specific organization by ID")
	datum.ViperBindFlag("org.get.id", orgGetCmd.Flags().Lookup("id"))
}

func orgs(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetGraphClient(ctx)
	if err != nil {
		return err
	}

	// save session cookies on function exit
	client, _ := cli.Client.(*datumclient.Client)
	defer datum.StoreSessionCookies(client)

	// filter options
	oID := viper.GetString("org.get.id")

	var s []byte

	writer := tables.NewTableWriter(orgCmd.OutOrStdout(), "ID", "Name", "Description", "PersonalOrg", "Children", "Members")

	// if an org ID is provided, filter on that organization, otherwise get all
	if oID != "" {
		org, err := cli.Client.GetOrganizationByID(ctx, oID, cli.Interceptor)
		if err != nil {
			return err
		}

		if datum.OutputFormat == "json" {
			s, err := json.Marshal(org.Organization)
			if err != nil {
				return err
			}

			return datum.JSONPrint(s)
		}

		writer.AddRow(org.Organization.ID, org.Organization.Name, *org.Organization.Description, *org.Organization.PersonalOrg, len(org.Organization.Children.Edges), len(org.Organization.Members))
		writer.Render()

		return nil
	}

	orgs, err := cli.Client.GetAllOrganizations(ctx, cli.Interceptor)
	if err != nil {
		return err
	}

	s, err = json.Marshal(orgs.Organizations)
	if err != nil {
		return err
	}

	if datum.OutputFormat == "json" {
		return datum.JSONPrint(s)
	}

	for _, o := range orgs.Organizations.Edges {
		writer.AddRow(o.Node.ID, o.Node.Name, *o.Node.Description, *o.Node.PersonalOrg, len(o.Node.Children.Edges), len(o.Node.Members))
	}

	writer.Render()

	return nil
}
