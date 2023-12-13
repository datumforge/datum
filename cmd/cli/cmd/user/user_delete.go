package datumuser

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

var userDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing datum user",
	RunE: func(cmd *cobra.Command, args []string) error {
		return deleteUser(cmd.Context())
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)

	userDeleteCmd.Flags().StringP("id", "i", "", "user id to delete")
	datum.ViperBindFlag("user.delete.id", userDeleteCmd.Flags().Lookup("id"))
}

func deleteUser(ctx context.Context) error {
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

	userID := viper.GetString("user.delete.id")
	if userID == "" {
		return datum.NewRequiredFieldMissingError("user id")
	}

	o, err := c.DeleteUser(ctx, userID, i)
	if err != nil {
		return err
	}

	s, err = json.Marshal(o)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
