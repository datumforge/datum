package datumlogin

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/term"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/httpserve/route"
	"github.com/datumforge/datum/pkg/datumclient"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/providers/github"
	"github.com/datumforge/datum/pkg/providers/google"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "authenticate with the datum API",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := login(cmd.Context())

		return err
	},
}

func init() {
	datum.RootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "username (email) to authenticate with password auth")
	datum.ViperBindFlag("login.username", loginCmd.Flags().Lookup("username"))

	loginCmd.Flags().StringP("oauth-provider", "o", "", "oauth provider used to authenticate, default empty for password authentication")
	datum.ViperBindFlag("login.provider", loginCmd.Flags().Lookup("oauth-provider"))
}

func login(ctx context.Context) (*oauth2.Token, error) {
	// setup datum http client
	client, err := datum.SetupClient(ctx)
	if err != nil {
		return nil, err
	}

	// setup tokens
	var tokens *oauth2.Token

	// check provider
	provider := viper.GetString("login.provider")

	if provider == "" {
		// store session cookies on function exit
		defer datum.StoreSessionCookies(client)

		username := viper.GetString("login.username")
		if username == "" {
			return nil, datum.NewRequiredFieldMissingError("username")
		}

		tokens, err = passwordAuth(ctx, client, username)
		if err != nil {
			return nil, err
		}
	} else {
		var session string
		tokens, session, err = providerAuth(ctx, client, provider)

		if err != nil {
			return nil, err
		}

		if session == "" {
			return nil, datum.ErrSessionNotFound
		}

		// because of the callback, the session is not stored in the cookie jar
		// so we need to store it ourselves instead of using the defer
		if err := datum.StoreSession(session); err != nil {
			fmt.Println("unable to store session in keychain")

			return nil, err
		}
	}

	fmt.Println("\nAuthentication Successful!")

	if err := datum.StoreToken(tokens); err != nil {
		return nil, err
	}

	fmt.Println("auth tokens successfully stored in keychain")

	return tokens, nil
}

func passwordAuth(ctx context.Context, client *datumclient.DatumClient, username string) (*oauth2.Token, error) {
	// read password from terminal if not set in environment variable
	password := os.Getenv("DATUM_PASSWORD")

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

// validateProvider validate the provider specified is configured
func providerAuth(ctx context.Context, client *datumclient.DatumClient, provider string) (*oauth2.Token, string, error) {
	isDev := strings.Contains(client.Config().BaseURL.String(), "localhost")

	switch strings.ToUpper(provider) {
	case google.ProviderName:
		endpoint := "google/login"
		u := fmt.Sprintf("%s%s/%s", client.Config().BaseURL.String(), route.V1Version, endpoint)

		return datumclient.OauthLogin(u, isDev)
	case github.ProviderName:
		endpoint := "github/login"
		u := fmt.Sprintf("%s%s/%s", client.Config().BaseURL.String(), route.V1Version, endpoint)

		return datumclient.OauthLogin(u, isDev)
	default:
		return nil, "", datum.ErrUnsupportedProvider
	}
}
