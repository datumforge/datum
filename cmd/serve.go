package cmd

import (
	"context"
	"time"

	"entgo.io/ent/dialect"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/api"
	"github.com/datumforge/datum/internal/echox"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
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
	serveCmd.Flags().Bool("debug", false, "enable server debug")
	viperBindFlag("server.debug", serveCmd.Flags().Lookup("debug"))

	serveCmd.Flags().String("listen", defaultListenAddr, "address to listen on")
	viperBindFlag("server.listen", serveCmd.Flags().Lookup("listen"))

	serveCmd.Flags().Bool("https", false, "enable serving from https")
	viperBindFlag("server.https", serveCmd.Flags().Lookup("https"))

	serveCmd.Flags().String("ssl-cert", "", "ssl cert file location")
	viperBindFlag("server.ssl-cert", serveCmd.Flags().Lookup("ssl-cert"))

	serveCmd.Flags().String("ssl-key", "", "ssl key file location")
	viperBindFlag("server.ssl-key", serveCmd.Flags().Lookup("ssl-key"))

	serveCmd.Flags().Bool("auto-cert", false, "automatically generate tls cert")
	viperBindFlag("server.auto-cert", serveCmd.Flags().Lookup("auto-cert"))

	serveCmd.Flags().String("cert-host", "example.com", "host to use to generate tls cert")
	viperBindFlag("server.cert-host", serveCmd.Flags().Lookup("cert-host"))

	serveCmd.Flags().Duration("shutdown-grace-period", echox.DefaultShutdownGracePeriod, "server shutdown grace period")
	viperBindFlag("server.shutdown-grace-period", serveCmd.Flags().Lookup("shutdown-grace-period"))

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

	// create default server config
	httpsEnabled := viper.GetBool("server.https")
	serverConfig := echox.Config{}.WithDefaults()

	// override with flags
	serverConfig = serverConfig.WithListen(viper.GetString("server.listen")).
		WithShutdownGracePeriod(viper.GetDuration("server.shutdown-grace-period")).
		WithDebug(viper.GetBool("server.debug")).
		WithDev(serveDevMode).
		WithHTTPS(httpsEnabled)

	if httpsEnabled {
		serverConfig = serverConfig.WithTLSDefaults()

		if viper.GetBool("server.auto-cert") {
			serverConfig = serverConfig.WithAutoCert(viper.GetString("server.cert-host"))
		} else {
			certFile, certKey, err := getCertFiles()
			if err != nil {
				return err
			}

			serverConfig = serverConfig.WithTLSCerts(certFile, certKey)
		}
	}

	srv, err := echox.NewServer(logger.Desugar(), serverConfig)
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
	}

	r := api.NewResolver(client).
		WithLogger(logger.Named("resolvers"))

	if !oidcEnabled {
		r = r.WithAuthDisabled(true)
	}

	handler := r.Handler(enablePlayground, mw...)

	srv.AddHandler(handler)

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
