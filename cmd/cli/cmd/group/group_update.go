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

var groupUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing datum group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateGroup(cmd.Context())
	},
}

func init() {
	groupCmd.AddCommand(groupUpdateCmd)

	groupUpdateCmd.Flags().StringP("id", "i", "", "group id to update")
	datum.ViperBindFlag("group.update.id", groupUpdateCmd.Flags().Lookup("id"))

	groupUpdateCmd.Flags().StringP("name", "n", "", "name of the group")
	datum.ViperBindFlag("group.update.name", groupUpdateCmd.Flags().Lookup("name"))

	groupUpdateCmd.Flags().StringP("short-name", "s", "", "display name of the group")
	datum.ViperBindFlag("group.update.short-name", groupUpdateCmd.Flags().Lookup("short-name"))

	groupUpdateCmd.Flags().StringP("description", "d", "", "description of the group")
	datum.ViperBindFlag("group.update.description", groupUpdateCmd.Flags().Lookup("description"))
}

func updateGroup(ctx context.Context) error {
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

	oID := viper.GetString("group.update.id")
	if oID == "" {
		return datum.NewRequiredFieldMissingError("group id")
	}

	name := viper.GetString("group.update.name")
	displayName := viper.GetString("group.update.short-name")
	description := viper.GetString("group.update.description")

	input := datumclient.UpdateGroupInput{}

	if name != "" {
		input.Name = &name
	}

	if displayName != "" {
		input.DisplayName = &displayName
	}

	if description != "" {
		input.Description = &description
	}

	o, err := c.UpdateGroup(ctx, oID, input, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
