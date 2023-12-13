package datumorg

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
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
	datum.ViperBindFlag("org.create.name", orgCreateCmd.Flags().Lookup("name"))

	orgCreateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	datum.ViperBindFlag("org.create.short-name", orgCreateCmd.Flags().Lookup("short-name"))

	orgCreateCmd.Flags().StringP("description", "d", "", "description of the organization")
	datum.ViperBindFlag("org.create.description", orgCreateCmd.Flags().Lookup("description"))

	orgCreateCmd.Flags().StringP("parent-org-id", "p", "", "parent organization id, leave empty to create a root org")
	datum.ViperBindFlag("org.create.parent-org-id", orgCreateCmd.Flags().Lookup("parent-org-id"))
}

func createOrg(ctx context.Context) error {
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	token := os.Getenv("DATUM_ACCESS_TOKEN")

	i := datumclient.WithAccessToken(token)

	// new client with params
	c := datumclient.NewClient(h, datum.GraphAPIHost, opt, i)

	var s []byte

	name := viper.GetString("org.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("organization name")
	}

	displayName := viper.GetString("org.create.short-name")
	description := viper.GetString("org.create.description")
	parentOrgID := viper.GetString("org.create.parent-org-id")

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

	o, err := c.CreateOrganization(ctx, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
