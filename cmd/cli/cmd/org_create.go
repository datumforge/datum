package datum

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	viperBindFlag("org.create.name", orgCreateCmd.Flags().Lookup("name"))

	orgCreateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	viperBindFlag("org.create.short-name", orgCreateCmd.Flags().Lookup("short-name"))

	orgCreateCmd.Flags().StringP("description", "d", "", "description of the organization")
	viperBindFlag("org.create.description", orgCreateCmd.Flags().Lookup("description"))

	orgCreateCmd.Flags().StringP("parent-org-id", "p", "", "parent organization id, leave empty to create a root org")
	viperBindFlag("org.create.parent-org-id", orgCreateCmd.Flags().Lookup("parent-org-id"))
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
	c := datumclient.NewClient(h, host, opt, i)

	var s []byte

	name := viper.GetString("org.create.name")
	if name == "" {
		return ErrOrgNameRequired
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

	return jsonPrint(s)
}
