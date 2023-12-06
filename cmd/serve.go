package cmd

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/auth"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/graphapi"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/server"
)

var (
	serveDevMode bool
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
	if err := entdb.RegisterDatabaseFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	// Auth configuration settings
	if err := auth.RegisterAuthFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	// OpenFGA configuration settings
	if err := fga.RegisterFGAFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	// only available as a CLI arg because these should only be used in dev environments
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables graph playground")
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		entdbClient      *ent.Client
		fgaClient        *fga.Client
		err              error
		mw               []echo.MiddlewareFunc
		enablePlayground bool
	)

	// debug setting
	debug := viper.GetBool("debug")

	// create ready checks
	readyChecks := handlers.Checks{}

	// auth setting
	authEnabled := viper.GetBool("oidc.enabled")

	// dev mode settings
	if serveDevMode {
		enablePlayground = true
	}

	// create ent dependency injection
	opts := []ent.Option{ent.Logger(*logger)}

	// add the fga client if oidc is enabled
	if authEnabled {
		config := fga.Config{
			Name:           "datum",
			Host:           viper.GetString("fga.host"),
			Scheme:         viper.GetString("fga.scheme"),
			StoreID:        viper.GetString("fga.store.id"),
			ModelID:        viper.GetString("fga.model.id"),
			CreateNewModel: viper.GetBool("fga.model.create"),
		}

		logger.Infow("setting up fga client", "host", config.Host, "scheme", config.Scheme)

		fgaClient, err = fga.CreateFGAClientWithStore(ctx, config, logger)
		if err != nil {
			return err
		}

		// add client as ent dependency
		opts = append(opts, ent.Authz(*fgaClient))

		// add ready checkz
		readyChecks.AddReadinessCheck("fga", fga.Healthcheck(*fgaClient))

		// add jwt middleware
		secretKey := []byte(viper.GetString("jwt.secretkey"))
		jwtMiddleware := auth.CreateJwtMiddleware([]byte(secretKey))

		mw = append(mw, jwtMiddleware)
	}

	// create new ent db client
	entConfig := entdb.EntClientConfig{
		Debug:           debug,
		DriverName:      dialect.SQLite,
		Logger:          *logger,
		PrimaryDBSource: viper.GetString("db.primary"),
	}

	if viper.GetBool("db.multi-write") {
		entConfig.SecondaryDBSource = viper.GetString("db.secondary")

		entdbClient, err = entConfig.NewMultiDriverDBClient(ctx, opts)
		if err != nil {
			return err
		}
	} else {
		entdbClient, err = entConfig.NewEntDBDriver(ctx, opts)
		if err != nil {
			return err
		}
	}

	defer entdbClient.Close()

	// Setup server config
	serverConfig, err := setupServerConfig(readyChecks, authEnabled)
	if err != nil {
		return err
	}

	cp, err := config.NewConfigProviderWithRefresh(serverConfig)
	if err != nil {
		return err
	}

	// Get server refresh config
	s, err := cp.GetConfig()
	if err != nil {
		return err
	}

	srv := server.NewServer(s.Server, s.Logger)

	// Setup Graph API Handlers
	r := graphapi.NewResolver(entdbClient, serverConfig.Auth.Enabled).
		WithLogger(logger.Named("resolvers"))

	handler := r.Handler(enablePlayground, mw...)

	// Add Graph Handler
	srv.AddHandler(handler)

	if err := srv.StartEchoServer(); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}

func setupServerConfig(readyChecks handlers.Checks, authEnabled bool) (*config.Config, error) {
	// Setup server config
	httpsEnabled := viper.GetBool("server.https")
	listenAddr := viper.GetString("server.listen")
	shutdownGracePeriod := viper.GetDuration("server.shutdown-grace-period")
	serverDebug := viper.GetBool("server.debug")
	autoCert := viper.GetBool("server.auto-cert")

	serverConfig := config.NewConfig().
		WithListen(listenAddr).                       // set custom port
		WithHTTPS(httpsEnabled).                      // enable https
		WithShutdownGracePeriod(shutdownGracePeriod). // override default grace period shutdown
		WithDebug(serverDebug).                       // enable debug mode
		WithDev(serveDevMode).                        // enable dev mode
		SetDefaults()                                 // set defaults if not already set

	serverConfig.Logger = logger.Desugar()
	serverConfig.RefreshInterval = viper.GetDuration("server.config-refresh")
	serverConfig.Auth.Enabled = authEnabled
	serverConfig.Server.Checks = readyChecks

	if httpsEnabled {
		serverConfig.WithTLSDefaults()

		if autoCert {
			serverConfig.WithAutoCert(viper.GetString("server.cert-host"))
		} else {
			certFile, certKey, err := server.GetCertFiles(viper.GetString("server.ssl-cert"), viper.GetString("server.ssl-key"))
			if err != nil {
				return nil, err
			}

			serverConfig.WithTLSCerts(certFile, certKey)
		}
	}

	return serverConfig, nil
}
