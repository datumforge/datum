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
	datumlogin "github.com/datumforge/datum/cmd/cli/cmd/login"
	"github.com/datumforge/datum/internal/datumclient"
)

var orgUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgUpdateCmd)

	orgUpdateCmd.Flags().StringP("id", "i", "", "org id to update")
	datum.ViperBindFlag("org.update.id", orgUpdateCmd.Flags().Lookup("id"))

	orgUpdateCmd.Flags().StringP("name", "n", "", "name of the organization")
	datum.ViperBindFlag("org.update.name", orgUpdateCmd.Flags().Lookup("name"))

	orgUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the organization")
	datum.ViperBindFlag("org.update.short-name", orgUpdateCmd.Flags().Lookup("short-name"))

	orgUpdateCmd.Flags().StringP("description", "d", "", "description of the organization")
	datum.ViperBindFlag("org.update.description", orgUpdateCmd.Flags().Lookup("description"))
}

func updateOrg(ctx context.Context) error {
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	token, err := datumlogin.GetTokenFromKeyring(ctx)
	if err != nil {
		return err
	}

	accessToken := token.AccessToken

	// if not stored, try the env var
	if accessToken == "" {
		accessToken = os.Getenv("DATUM_ACCESS_TOKEN")
	}

	i := datumclient.WithAccessToken(accessToken)
	// new client with params
	c := datumclient.NewClient(h, datum.GraphAPIHost, opt, i)

	var s []byte

	oID := viper.GetString("org.update.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	name := viper.GetString("org.update.name")
	displayName := viper.GetString("org.update.short-name")
	description := viper.GetString("org.update.description")

	input := datumclient.UpdateOrganizationInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.UpdateOrganization(ctx, oID, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
