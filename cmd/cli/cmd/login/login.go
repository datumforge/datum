package datumlogin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"syscall"

	"github.com/99designs/keyring"
	"github.com/Yamashou/gqlgenc/clientv2"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/term"

	datum "github.com/datumforge/datum/cmd/cli/cmd"
	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

var (
	userKeyring       keyring.Keyring
	userKeyringLoaded = false
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

	login := handlers.User{
		Username: username,
		Password: password,
	}

	// setup datum http client
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := datumclient.NewClient(h, datum.DatumHost, opt, nil)

	c2 := c.(*datumclient.Client)

	tokens, err := datumclient.Login(c2, ctx, login)
	if err != nil {
		return nil, err
	}

	fmt.Println("\nAuthentication Successful!")

	if err := storeToken(tokens, "datum"); err != nil {
		return nil, err
	}

	fmt.Println("auth token successfully stored in keychain")
	fmt.Println(tokens.AccessToken)

	return tokens, nil
}

func storeToken(token *oauth2.Token, name string) error {
	ring, err := getKeyring()
	if err != nil {
		return fmt.Errorf("error opening keyring: %w", err)
	}

	err = ring.Set(keyring.Item{
		Key:  fmt.Sprintf("%s_token", name),
		Data: []byte(token.AccessToken),
	})
	if err != nil {
		return fmt.Errorf("failed saving access token: %w", err)
	}

	err = ring.Set(keyring.Item{
		Key:  fmt.Sprintf("%s_refresh_token", name),
		Data: []byte(token.RefreshToken),
	})
	if err != nil {
		return fmt.Errorf("failed saving refresh token: %w", err)
	}

	return nil
}

// getKeyring will return the already loaded keyring so that we don't prompt users for passwords multiple times
func getKeyring() (keyring.Keyring, error) {
	var err error

	if userKeyringLoaded {
		return userKeyring, nil
	}

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	userKeyring, err = keyring.Open(keyring.Config{
		ServiceName: "datum",

		// MacOS keychain
		KeychainTrustApplication: true,

		// KDE Wallet
		KWalletAppID:  "datum",
		KWalletFolder: "datum",

		// Windows
		WinCredPrefix: "datum",

		// Fallback encrypted file
		FileDir:          path.Join(cfgDir, "datum", "keyring"),
		FilePasswordFunc: keyring.TerminalPrompt,
	})
	if err == nil {
		userKeyringLoaded = true
	}

	return userKeyring, err
}

// GetTokenFromKeyring will return the oauth token from the keyring
func GetTokenFromKeyring(ctx context.Context) (*oauth2.Token, error) {
	ring, err := getKeyring()
	if err != nil {
		return nil, fmt.Errorf("error opening keyring: %w", err)
	}

	authToken, err := ring.Get("datum_token")
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			return login(ctx)
		}

		return nil, fmt.Errorf("error fetching auth token: %w", err)
	}

	refToken, err := ring.Get("datum_refresh_token")
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			return login(ctx)
		}

		return nil, fmt.Errorf("error fetching refresh token: %w", err)
	}

	// TODO (sfunk): add refresh logic

	return &oauth2.Token{AccessToken: string(authToken.Data), RefreshToken: string(refToken.Data)}, nil
}
