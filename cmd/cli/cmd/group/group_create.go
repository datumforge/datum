package datumgroup

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

var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.Flags().StringP("name", "n", "", "name of the group")
	datum.ViperBindFlag("group.create.name", groupCreateCmd.Flags().Lookup("name"))

	groupCreateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	datum.ViperBindFlag("group.create.short-name", groupCreateCmd.Flags().Lookup("short-name"))

	groupCreateCmd.Flags().StringP("description", "d", "", "description of the group")
	datum.ViperBindFlag("group.create.description", groupCreateCmd.Flags().Lookup("description"))

	groupCreateCmd.Flags().StringP("owner-id", "o", "", "owner org id")
	datum.ViperBindFlag("group.create.owner-id", groupCreateCmd.Flags().Lookup("owner-id"))
}

func createGroup(ctx context.Context) error {
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

	name := viper.GetString("group.create.name")
	if name == "" {
		return datum.NewRequiredFieldMissingError("group name")
	}

	owner := viper.GetString("group.create.owner-id")
	if owner == "" {
		return datum.NewRequiredFieldMissingError("organization id")
	}

	displayName := viper.GetString("group.create.short-name")
	description := viper.GetString("group.create.description")

	input := datumclient.CreateGroupInput{
		Name:    name,
		OwnerID: owner,
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.CreateGroup(ctx, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
