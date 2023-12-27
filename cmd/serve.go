package cmd

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/config"
	authmw "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/httpserve/serveropts"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/marionette"
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
	if err := tokens.RegisterAuthFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	// OpenFGA configuration settings
	if err := fga.RegisterFGAFlags(viper.GetViper(), serveCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		entdbClient *ent.Client
		fgaClient   *fga.Client
		err         error
		mw          []echo.MiddlewareFunc
	)

	// create ent dependency injection
	entOpts := []ent.Option{ent.Logger(*logger)}

	// get settings for the server
	settings := viper.AllSettings()

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithServer(settings),
		serveropts.WithLogger(logger),
		serveropts.WithHTTPS(settings),
		serveropts.WithSQLiteDB(settings),
		serveropts.WithAuth(settings),
		serveropts.WithFGAAuthz(settings),
	)

	// Create keys for development
	if dev := viper.GetBool("server.dev"); dev {
		serverOpts = append(serverOpts, serveropts.WithGeneratedKeys(settings))
	}

	so := serveropts.NewServerOptions(serverOpts)

	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	if so.Config.Authz.Enabled {
		az := so.Config.Authz
		config := fga.NewAuthzConfig(az, logger)

		fgaClient, err = fga.CreateFGAClientWithStore(ctx, config)
		if err != nil {
			return err
		}

		// add client as ent dependency
		entOpts = append(entOpts, ent.Authz(*fgaClient))

		// add auth middleware
		authMiddleware := authmw.Authenticate()

		mw = append(mw, authMiddleware)
	}

	// Setup DB connection
	dbConfig := entdb.NewDBConfig(so.Config.DB, logger)

	entdbClient, err = dbConfig.NewMultiDriverDBClient(ctx, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// add ready checks
	so.AddServerOptions(serveropts.WithReadyChecks(dbConfig, fgaClient))

	// Start task manager
	tmConfig := marionette.Config{
		Logger: logger,
	}

	marionette.New(tmConfig).Start()

	// Add Driver to the Handlers Config
	so.Config.Server.Handler.DBClient = entdbClient

	srv := server.NewServer(so.Config.Server, so.Config.Logger)

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient, settings, mw))

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
