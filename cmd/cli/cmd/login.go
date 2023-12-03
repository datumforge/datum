package datum

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/cmd/cli/vars"
	"github.com/datumforge/datum/internal/auth"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	scopes = []string{"email", "profile"}
)

// loginCmd represents the command to login to datum
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Datum account",
	Long:  "Log in to a Datum account via clerk authentication oauth2.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return login(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func login(ctx context.Context) error {
	ep := oauth2.Endpoint{
		AuthURL:  vars.AuthURL,
		TokenURL: vars.TokenURL,
	}

	oauthConfig := &oauth2.Config{
		ClientID:    vars.OauthClientID,
		Scopes:      scopes,
		Endpoint:    ep,
		RedirectURL: vars.RedirectURL,
	}

	token, err := auth.AuthPKCE(oauthConfig, vars.Audience)
	if err != nil {
		return err
	}

	fmt.Printf("auth token successfully retrieved, %v.\n", token)

	return nil
}
