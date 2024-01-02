package cmd

import (
	"context"

	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/config"
	authmw "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/httpserve/serveropts"
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

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithServer(),
		serveropts.WithLogger(logger),
		serveropts.WithHTTPS(),
		serveropts.WithSQLiteDB(),
		serveropts.WithAuth(),
		serveropts.WithFGAAuthz(),
		serveropts.WithEmailManager(),
		serveropts.WithTaskManager(),
		serveropts.WithSessionManager(),
	)

	so := serveropts.NewServerOptions(serverOpts)

	// Create keys for development
	if so.Config.Server.Dev {
		so.AddServerOptions(serveropts.WithGeneratedKeys())
	}

	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	if so.Config.Authz.Enabled {
		fgaClient, err = fga.CreateFGAClientWithStore(ctx, so.Config.Authz)
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

	// Add Driver to the Handlers Config
	so.Config.Server.Handler.DBClient = entdbClient

	srv := server.NewServer(so.Config.Server, so.Config.Logger)

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient, mw))

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
