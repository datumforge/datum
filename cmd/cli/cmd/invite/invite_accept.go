package datuminvite

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

var inviteAcceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "accept an invite",
	RunE: func(cmd *cobra.Command, args []string) error {
		return inviteAccept(cmd.Context())
	},
}

func init() {
	inviteCmd.AddCommand(inviteAcceptCmd)

	inviteAcceptCmd.Flags().StringP("token", "t", "", "invite token")
	datum.ViperBindFlag("invite.accept.token", inviteAcceptCmd.Flags().Lookup("token"))

	inviteAcceptCmd.Flags().StringP("password", "p", "", "new password of the user")
	datum.ViperBindFlag("invite.accept.password", inviteAcceptCmd.Flags().Lookup("password"))

	inviteAcceptCmd.Flags().StringP("first-name", "f", "", "first name of the user")
	datum.ViperBindFlag("invite.accept.first-name", inviteAcceptCmd.Flags().Lookup("first-name"))

	inviteAcceptCmd.Flags().StringP("last-name", "l", "", "last name of the user")
	datum.ViperBindFlag("invite.accept.last-name", inviteAcceptCmd.Flags().Lookup("last-name"))
}

func inviteAccept(ctx context.Context) error {
	var s []byte

	password := viper.GetString("invite.accept.password")
	if password == "" {
		return datum.NewRequiredFieldMissingError("password")
	}

	token := viper.GetString("invite.accept.token")
	if token == "" {
		return datum.NewRequiredFieldMissingError("token")
	}

	firstName := viper.GetString("invite.accept.first-name")
	if firstName == "" {
		return datum.NewRequiredFieldMissingError("first name")
	}

	lastName := viper.GetString("invite.accept.last-name")
	if lastName == "" {
		return datum.NewRequiredFieldMissingError("last name")
	}

	invite := handlers.Invite{
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Token:     token,
	}

	// setup datum http client with cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	h := &http.Client{
		Jar: jar,
	}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := datumclient.NewClient(h, datum.DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*datumclient.Client)

	defer datum.StoreSessionCookies(dc)

	registration, tokens, err := datumclient.OrgInvite(dc, ctx, invite)
	if err != nil {
		return err
	}

	if err := datum.StoreToken(tokens); err != nil {
		return err
	}

	s, err = json.Marshal(registration)
	if err != nil {
		return err
	}

	return datum.JSONPrint(s)
}
