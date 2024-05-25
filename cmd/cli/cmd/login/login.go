package datumlogin

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"syscall"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/term"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/httpserve/route"
	api "github.com/datumforge/datum/pkg/datumclient"
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
	// setup datum http client with cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	h := &http.Client{
		Jar: jar,
	}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := api.NewClient(h, datum.DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*api.Client)

	// setup tokens
	var tokens *oauth2.Token

	// check provider
	provider := viper.GetString("login.provider")

	if provider == "" {
		// store session cookies on function exit
		defer datum.StoreSessionCookies(dc)

		username := viper.GetString("login.username")
		if username == "" {
			return nil, datum.NewRequiredFieldMissingError("username")
		}

		tokens, err = passwordAuth(ctx, dc, username)
		if err != nil {
			return nil, err
		}
	} else {
		var session string
		tokens, session, err = providerAuth(ctx, dc, provider)

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

func passwordAuth(ctx context.Context, client *api.Client, username string) (*oauth2.Token, error) {
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

	return api.Login(client, ctx, login)
}

// validateProvider validate the provider specified is configured
func providerAuth(ctx context.Context, client *api.Client, provider string) (*oauth2.Token, string, error) {
	isDev := strings.Contains(client.Client.BaseURL, "localhost")

	switch strings.ToUpper(provider) {
	case google.ProviderName:
		endpoint := "google/login"
		u := fmt.Sprintf("%s%s/%s", client.Client.BaseURL, route.V1Version, endpoint)

		return api.OauthLogin(u, isDev)
	case github.ProviderName:
		endpoint := "github/login"
		u := fmt.Sprintf("%s%s/%s", client.Client.BaseURL, route.V1Version, endpoint)

		return api.OauthLogin(u, isDev)
	default:
		return nil, "", datum.ErrUnsupportedProvider
	}
}
