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

	// filter options
	oID := viper.GetString("org.get.id")

	var s []byte

	// if an org ID is provided, filter on that organization, otherwise get all
	if oID == "" {
		orgs, err := c.GetAllOrganizations(ctx, i)
		if err != nil {
			return err
		}

		s, err = json.Marshal(orgs)
		if err != nil {
			return err
		}
	} else {
		org, err := c.GetOrganizationByID(ctx, oID, i)
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
