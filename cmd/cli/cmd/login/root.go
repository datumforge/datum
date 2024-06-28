package datumlogin

import (
	"context"
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"golang.org/x/term"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/models"
)

var cmd = &cobra.Command{
	Use:   "login",
	Short: "authenticate with the datum API using password credentials",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := login(cmd.Context())
		cobra.CheckErr(err)
	},
}

func init() {
	datum.RootCmd.AddCommand(cmd)

	cmd.Flags().StringP("username", "u", "", "username (email) to authenticate with password auth")
}

func login(ctx context.Context) (*oauth2.Token, error) {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	cobra.CheckErr(err)

	username := datum.Config.String("username")
	if username == "" {
		return nil, datum.NewRequiredFieldMissingError("username")
	}

	tokens, err := passwordAuth(ctx, client, username)
	cobra.CheckErr(err)

	fmt.Println("\nAuthentication Successful!")

	err = datum.StoreToken(tokens)
	cobra.CheckErr(err)

	datum.StoreSessionCookies(client)

	fmt.Println("auth tokens successfully stored in keychain")

	return tokens, nil
}

func passwordAuth(ctx context.Context, client *datumclient.DatumClient, username string) (*oauth2.Token, error) {
	// read password from terminal if not set in environment variable
	password := datum.Config.String("password")

	if password == "" {
		fmt.Print("Password: ")

		bytepw, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}

		password = string(bytepw)
	}

	login := models.LoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := client.Login(ctx, &login)
	if err != nil {
		return nil, err
	}

	return &oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		TokenType:    resp.TokenType,
	}, nil
}
