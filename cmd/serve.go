package cmd

import (
	"context"
	"log"
	"time"

	"entgo.io/ent/dialect"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/handlers"
)

const (
	defaultListenAddr            = ":17608"
	defaultDBPrimaryURI          = "datum.db?mode=memory&_fk=1"
	defaultDBSecondaryURI        = "backup.db?mode=memory&_fk=1"
	defaultOIDCJWKSRemoteTimeout = 5 * time.Second
	defaultFGAScheme             = "https"
	defaultFGAHost               = ""
)

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the example Graph API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Server flags
	if err := config.RegisterServerFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	// Database flags
	serveCmd.Flags().Bool("db-mutli-write", false, "write to a primary and secondary database")
	viperBindFlag("server.db.multi-write", serveCmd.Flags().Lookup("db-mutli-write"))

	serveCmd.Flags().String("db-primary", defaultDBPrimaryURI, "db primary uri")
	viperBindFlag("server.db-primary", serveCmd.Flags().Lookup("db-primary"))

	serveCmd.Flags().String("db-secondary", defaultDBSecondaryURI, "db secondary uri")
	viperBindFlag("server.db-secondary", serveCmd.Flags().Lookup("db-secondary"))

	// echo-jwt flags
	serveCmd.Flags().String("jwt-secretkey", "", "secret key for echojwt config")
	viperBindFlag("jwt.secretkey", serveCmd.Flags().Lookup("jwt-secretkey"))

	// OIDC Flags
	serveCmd.Flags().Bool("oidc", true, "use oidc auth")
	viperBindFlag("oidc.enabled", serveCmd.Flags().Lookup("oidc"))

	serveCmd.Flags().String("oidc-aud", "", "expected audience on OIDC JWT")
	viperBindFlag("oidc.audience", serveCmd.Flags().Lookup("oidc-aud"))

	serveCmd.Flags().String("oidc-issuer", "", "expected issuer of OIDC JWT")
	viperBindFlag("oidc.issuer", serveCmd.Flags().Lookup("oidc-issuer"))

	serveCmd.Flags().Duration("oidc-jwks-remote-timeout", defaultOIDCJWKSRemoteTimeout, "timeout for remote JWKS fetching")
	viperBindFlag("oidc.jwks.remote-timeout", serveCmd.Flags().Lookup("oidc-jwks-remote-timeout"))

	// OpenFGA configuration settings
	serveCmd.Flags().String("fga-host", defaultFGAHost, "fga host without the scheme (e.g. api.fga.example instead of https://api.fga.example)")
	viperBindFlag("fga.host", serveCmd.Flags().Lookup("fga-host"))

	serveCmd.Flags().String("fga-scheme", defaultFGAScheme, "fga scheme (http vs. https)")
	viperBindFlag("fga.scheme", serveCmd.Flags().Lookup("fga-scheme"))

	serveCmd.Flags().String("fga-store-id", "", "fga store ID")
	viperBindFlag("fga.store.id", serveCmd.Flags().Lookup("fga-store-id"))

	serveCmd.Flags().String("fga-model-id", "", "fga authorization model ID")
	viperBindFlag("fga.model.id", serveCmd.Flags().Lookup("fga-model-id"))

	serveCmd.Flags().Bool("fga-model-create", false, "force create a fga authorization model, this should be used when a model exists, but transitioning to a new model")
	viperBindFlag("fga.model.create", serveCmd.Flags().Lookup("fga-model-create"))

	// only available as a CLI arg because these should only be used in dev environments
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		client      *ent.Client
		err         error
		oidcEnabled = viper.GetBool("oidc.enabled")
	)

	entConfig := entdb.EntClientConfig{
		Debug:           viper.GetBool("debug"),
		DriverName:      dialect.SQLite,
		Logger:          *logger,
		PrimaryDBSource: viper.GetString("server.db-primary"),
	}

	// create ent dependency injection
	opts := []ent.Option{ent.Logger(*logger)}

	// add the fga client if oidc is enabled
	var fgaClient *fga.Client

	if oidcEnabled {
		config := fga.Config{
			Name:           "datum",
			Host:           viper.GetString("fga.host"),
			Scheme:         viper.GetString("fga.scheme"),
			StoreID:        viper.GetString("fga.store.id"),
			ModelID:        viper.GetString("fga.model.id"),
			CreateNewModel: viper.GetBool("fga.model.create"),
		}

		logger.Infow(
			"setting up fga client",
			"host",
			config.Host,
			"scheme",
			config.Scheme,
		)

		fgaClient, err = fga.CreateFGAClientWithStore(ctx, config, logger)
		if err != nil {
			return err
		}

		opts = append(opts, ent.Authz(*fgaClient))
	}

	// create new ent db client
	if viper.GetBool("server.db.multi-write") {
		entConfig.SecondaryDBSource = viper.GetString("server.db-secondary")

		client, err = entConfig.NewMultiDriverDBClient(ctx, opts)
		if err != nil {
			return err
		}
	} else {
		client, err = entConfig.NewEntDBDriver(ctx, opts)
		if err != nil {
			return err
		}
	}
	defer client.Close()

	var mw []echo.MiddlewareFunc

	// dev mode settings
	if serveDevMode {
		enablePlayground = true
	}

	// add jwt middleware
	if oidcEnabled {
		secretKey := viper.GetString("jwt.secretkey")
		jwtConfig := createJwtMiddleware([]byte(secretKey))

		mw = append(mw, jwtConfig)
	}

	// TODO (sfunk): move flags over to package
	httpsEnabled := viper.GetBool("server.https")
	listenAddr := viper.GetString("server.listen")
	shutdownGracePeriod := viper.GetDuration("server.shutdown-grace-period")
	debug := viper.GetBool("server.debug")
	autoCert := viper.GetBool("server.auto-cert")

	serverConfig := config.NewConfig().
		SetDefaults().                                // set defaults
		WithListen(listenAddr).                       // set custom port
		WithShutdownGracePeriod(shutdownGracePeriod). // override default grace period shutdown
		WithDebug(debug).                             // enable debug mode
		WithDev(serveDevMode).                        // enable dev mode
		WithHTTPS(httpsEnabled)                       // enable https

	if httpsEnabled {
		serverConfig.WithTLSDefaults()

		if autoCert {
			serverConfig.WithAutoCert(viper.GetString("server.cert-host"))
		} else {
			certFile, certKey, err := getCertFiles()
			if err != nil {
				return err
			}

			serverConfig.WithTLSCerts(certFile, certKey)
		}
	}

	// srv, err := echox.NewServer(logger.Desugar(), serverConfig)
	// if err != nil {
	// 	logger.Error("failed to create server", zap.Error(err))
	// }

	// Setup Graph API Handlers
	r := graphapi.NewResolver(client).
		WithLogger(logger.Named("resolvers"))

	if !oidcEnabled {
		serverConfig.Auth.Enabled = false
		r = r.WithAuthDisabled(true)
	}

	h := handlers.NewHandlers()

	handler := r.Handler(enablePlayground, mw...)

	h.AddHandler(handler)

	if err := srv.RunWithContext(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}

// createJwtMiddleware, TODO expand the config settings
func createJwtMiddleware(secret []byte) echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: secret,
	}

	return echojwt.WithConfig(config)
}

// getCertFiles for https enabled echo server
// TODO (sfunk): move to httpserve
func getCertFiles() (string, string, error) {
	certFile := viper.GetString("server.ssl-cert")
	keyFile := viper.GetString("server.ssl-key")

	if certFile == "" {
		return "", "", ErrCertFileMissing
	}

	if keyFile == "" {
		return "", "", ErrKeyFileMissing
	}

	return certFile, keyFile, nil
}
