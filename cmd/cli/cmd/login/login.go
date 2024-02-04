package datumlogin

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"syscall"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/term"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/analytics"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
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
}

func login(ctx context.Context) (*oauth2.Token, error) {
	// currently only username/password authentication is supported
	// so username must be provided
	username := viper.GetString("login.username")
	if username == "" {
		return nil, datum.NewRequiredFieldMissingError("username")
	}

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

	login := handlers.LoginRequest{
		Username: username,
		Password: password,
	}

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
	c := datumclient.NewClient(h, datum.DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*datumclient.Client)

	defer datum.StoreSessionCookies(dc)

	tokens, err := datumclient.Login(dc, ctx, login)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nAuthentication Successful!")

	if err := datum.StoreToken(tokens); err != nil {
		return nil, err
	}

	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	orgID, err := auth.GetOrganizationIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	analytics.AssociateUser(userID, orgID)

	analytics.Event("login", map[string]string{
		"kind": "cli-login",
	})

	fmt.Println("auth tokens successfully stored in keychain")

	return tokens, nil
}
