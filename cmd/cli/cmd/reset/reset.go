package reset

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/models"
)

var registerCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset a datum user password",
	RunE: func(cmd *cobra.Command, args []string) error {
		return registerUser(cmd.Context())
	},
}

func init() {
	datum.RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("token", "t", "", "reset token")
	datum.ViperBindFlag("reset.token", registerCmd.Flags().Lookup("token"))

	registerCmd.Flags().StringP("password", "p", "", "new password of the user")
	datum.ViperBindFlag("reset.password", registerCmd.Flags().Lookup("password"))
}

func registerUser(ctx context.Context) error {
	var s []byte

	password := viper.GetString("reset.password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	token := viper.GetString("reset.token")
	if token == "" {
		return datum.NewRequiredFieldMissingError("token")
	}

	reset := models.ResetPasswordRequest{
		Password: password,
		Token:    token,
	}

	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := datumclient.NewClient(h, datum.DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*datumclient.Client)

	reply, err := datumclient.Reset(dc, ctx, reset)
	if err != nil {
		return err
	}

	s, err = json.Marshal(reply)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
