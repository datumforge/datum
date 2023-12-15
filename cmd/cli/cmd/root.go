// Package datum is our cobra/viper cli implementation
package datum

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/99designs/keyring"
	"github.com/TylerBrock/colorjson"
	"github.com/Yamashou/gqlgenc/clientv2"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/datumforge/datum/internal/datumclient"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/tokens"
)

const (
	appName         = "datum"
	defaultRootHost = "http://localhost:17608/"
	graphEndpoint   = "query"
)

var (
	cfgFile string
	Logger  *zap.SugaredLogger
)

var (
	// DatumHost contains the root url for the Datum API
	DatumHost string
	// GraphAPIHost contains the url for the Datum graph api
	GraphAPIHost string
)

var (
	userKeyring       keyring.Keyring
	userKeyringLoaded = false
)

type CLI struct {
	Client      datumclient.DatumClient
	Interceptor clientv2.RequestInterceptor
	AccessToken string
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   appName,
	Short: "the datum cli",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+appName+".yaml)")
	ViperBindFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	RootCmd.PersistentFlags().StringVar(&DatumHost, "host", defaultRootHost, "api host url")
	ViperBindFlag("datum.host", RootCmd.PersistentFlags().Lookup("host"))

	RootCmd.PersistentFlags().Bool("no-auth", false, "disable attempts to look up access token, any calls that require auth will fail")
	ViperBindFlag("no-auth", RootCmd.PersistentFlags().Lookup("no-auth"))

	// Logging flags
	RootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	ViperBindFlag("logging.debug", RootCmd.PersistentFlags().Lookup("debug"))

	RootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	ViperBindFlag("logging.pretty", RootCmd.PersistentFlags().Lookup("pretty"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".datum" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".datum")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix("datum")
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	GraphAPIHost = fmt.Sprintf("%s%s", DatumHost, graphEndpoint)

	setupLogging()

	if err == nil {
		Logger.Infow("using config file", "file", viper.ConfigFileUsed())
	}
}

func setupLogging() {
	cfg := zap.NewProductionConfig()
	if viper.GetBool("logging.pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	if viper.GetBool("logging.debug") {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Logger = l.Sugar().With("app", appName)
	defer Logger.Sync() //nolint:errcheck
}

// ViperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func ViperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

func JSONPrint(s []byte) error {
	var obj map[string]interface{}

	err := json.Unmarshal(s, &obj)
	if err != nil {
		return err
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	o, err := f.Marshal(obj)
	if err != nil {
		return err
	}

	fmt.Println(string(o))

	return nil
}

func GetClient(ctx context.Context) (CLI, error) {
	cli := CLI{}
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// if auth is disabled in the client, proceed without
	// obtaining the token
	if viper.GetBool("no-auth") {
		i := datumclient.WithEmptyInterceptor()

		cli.Client = datumclient.NewClient(h, GraphAPIHost, opt, i)
		cli.Interceptor = i
		cli.AccessToken = ""

		return cli, nil
	}

	// setup interceptors
	token, err := GetTokenFromKeyring(ctx)
	if err != nil {
		return cli, err
	}

	expired, err := tokens.IsExpired(token.AccessToken)
	if err != nil {
		return cli, err
	}

	// refresh and store the new token pair if the existing access token
	// is expired
	if expired {
		// refresh the token pair using the refresh token
		token, err = refreshToken(ctx, token.RefreshToken)
		if err != nil {
			return cli, err
		}

		// store the new token
		if err := StoreToken(token); err != nil {
			return cli, err
		}
	}

	accessToken := token.AccessToken

	i := datumclient.WithAccessToken(accessToken)

	cli.Client = datumclient.NewClient(h, GraphAPIHost, opt, i)
	cli.Interceptor = i
	cli.AccessToken = accessToken

	// new client with params
	return cli, nil
}

// GetTokenFromKeyring will return the oauth token from the keyring
// if the token is expired, but the refresh token is still valid, the
// token will be refreshed
func GetTokenFromKeyring(ctx context.Context) (*oauth2.Token, error) {
	ring, err := GetKeyring()
	if err != nil {
		return nil, fmt.Errorf("error opening keyring: %w", err)
	}

	access, err := ring.Get("datum_token")
	if err != nil {
		return nil, fmt.Errorf("error fetching auth token: %w", err)
	}

	refresh, err := ring.Get("datum_refresh_token")
	if err != nil {
		return nil, fmt.Errorf("error fetching refresh token: %w", err)
	}

	return &oauth2.Token{
		AccessToken:  string(access.Data),
		RefreshToken: string(refresh.Data),
	}, nil
}

func refreshToken(ctx context.Context, refresh string) (*oauth2.Token, error) {
	// setup datum http client
	h := &http.Client{}

	// set options
	opt := &clientv2.Options{}

	// new client with params
	c := datumclient.NewClient(h, DatumHost, opt, nil)

	// this allows the use of the graph client to be used for the REST endpoints
	dc := c.(*datumclient.Client)

	req := handlers.RefreshRequest{
		RefreshToken: refresh,
	}

	token, err := datumclient.Refresh(dc, ctx, req)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// GetKeyring will return the already loaded keyring so that we don't prompt users for passwords multiple times
func GetKeyring() (keyring.Keyring, error) {
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

// StoreToken in local keyring
func StoreToken(token *oauth2.Token) error {
	ring, err := GetKeyring()
	if err != nil {
		return fmt.Errorf("error opening keyring: %w", err)
	}

	err = ring.Set(keyring.Item{
		Key:  "datum_token",
		Data: []byte(token.AccessToken),
	})
	if err != nil {
		return fmt.Errorf("failed saving access token: %w", err)
	}

	err = ring.Set(keyring.Item{
		Key:  "datum_refresh_token",
		Data: []byte(token.RefreshToken),
	})
	if err != nil {
		return fmt.Errorf("failed saving refresh token: %w", err)
	}

	return nil
}
