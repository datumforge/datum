package datumorg

import (
	"context"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
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

	orgGetCmd.Flags().StringP("id", "i", "", "org id to query")
	datum.ViperBindFlag("org.get.id", orgGetCmd.Flags().Lookup("id"))
}

func orgs(ctx context.Context) error {
	// setup datum http client
	cli, err := datum.GetClient(ctx)
	if err != nil {
		return err
	}

	// filter options
	oID := viper.GetString("org.get.id")

	var s []byte

	// if an org ID is provided, filter on that organization, otherwise get all
	if oID == "" {
		orgs, err := cli.Client.GetAllOrganizations(ctx, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(orgs)
		if err != nil {
			return err
		}
	} else {
		org, err := cli.Client.GetOrganizationByID(ctx, oID, cli.Interceptor)
		if err != nil {
			return err
		}

		s, err = json.Marshal(org)
		if err != nil {
			return err
		}
	}

	return datum.JSONPrint(s)
}
