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

	// Setup server config
	serverConfig, err := setupConfig(readyChecks, authEnabled)
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

	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	if authEnabled {
		config := fga.NewAuthzConfig(s.Authz, logger)

		fgaClient, err = fga.CreateFGAClientWithStore(ctx, *config)
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

	// Setup DB connection
	dbConfig := entdb.NewDBConfig(s.DB, logger)

	if viper.GetBool("db.multi-write") {
		entdbClient, err = dbConfig.NewMultiDriverDBClient(ctx, opts)
		if err != nil {
			return err
		}
	} else {
		entdbClient, err = dbConfig.NewEntDBDriver(ctx, opts)
		if err != nil {
			return err
		}
	}

	defer entdbClient.Close()

	srv := server.NewServer(s.Server, s.Logger.Desugar())

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

// TODO: move this to serveropts
func setupConfig(readyChecks handlers.Checks, authEnabled bool) (*config.Config, error) {
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

	// Logger setup
	serverConfig.Logger = logger

	// Refresh Interval Setup
	serverConfig.RefreshInterval = viper.GetDuration("server.config-refresh")

	// Auth Setup
	serverConfig.Auth.Enabled = authEnabled

	// Database Settings
	dbConfig := config.DB{
		Debug:           serverDebug,
		DriverName:      dialect.SQLite,
		PrimaryDBSource: viper.GetString("db.primary"),
	}

	if viper.GetBool("db.multi-write") {
		dbConfig.SecondaryDBSource = viper.GetString("db.secondary")
	}

	serverConfig.DB = dbConfig

	// Authz Setup
	authzConfig := config.Authz{
		Enabled:        authEnabled,
		StoreName:      "datum",
		Host:           viper.GetString("fga.host"),
		Scheme:         viper.GetString("fga.scheme"),
		StoreID:        viper.GetString("fga.store.id"),
		ModelID:        viper.GetString("fga.model.id"),
		CreateNewModel: viper.GetBool("fga.model.create"),
	}

	serverConfig.Authz = authzConfig

	return serverConfig, nil
}
