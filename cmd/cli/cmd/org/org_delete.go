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

var orgDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteOrg(cmd.Context())
	},
}

func init() {
	orgCmd.AddCommand(orgDeleteCmd)

	orgDeleteCmd.Flags().StringP("id", "i", "", "org id to delete")
	datum.ViperBindFlag("org.delete.id", orgDeleteCmd.Flags().Lookup("id"))
}

func deleteOrg(ctx context.Context) error {
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

	oID := viper.GetString("org.delete.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	o, err := c.DeleteOrganization(ctx, oID, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
