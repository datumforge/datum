package cmd

import (
	"context"

	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/analytics"
	"github.com/datumforge/datum/internal/analytics/posthog"
	"github.com/datumforge/datum/internal/cache"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/httpserve/serveropts"
	"github.com/datumforge/datum/internal/otelx"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Datum Graph API",
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
		fgaClient   *fgax.Client
		err         error
	)

	// create ent dependency injection
	entOpts := []ent.Option{ent.Logger(*logger)}

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithLogger(logger),
		serveropts.WithHTTPS(),
<<<<<<< HEAD
		serveropts.WithAuth(),
=======
		serveropts.WithSQLiteDB(),
		serveropts.WithRedisCache(),
		serveropts.WithFGAAuthz(),
		serveropts.WithTracer(),
>>>>>>> 0616791e (update auth and session middleware to support PAT auth)
		serveropts.WithEmailManager(),
		serveropts.WithTaskManager(),
		serveropts.WithSentry(),
		serveropts.WithMiddleware(),
	)

	so := serveropts.NewServerOptions(serverOpts)

	err = otelx.NewTracer(so.Config.Settings.Tracer, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

<<<<<<< HEAD
	// Create keys for development
	if so.Config.Settings.Server.Dev {
		so.AddServerOptions(serveropts.WithGeneratedKeys())
	}

=======
>>>>>>> 0616791e (update auth and session middleware to support PAT auth)
	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	fgaClient, err = fgax.CreateFGAClientWithStore(ctx, so.Config.Settings.Authz, so.Config.Logger)
	if err != nil {
		return err
	}

	phclient := posthog.Init()
	analytics := analytics.EventManager{}
	analytics.Handler = phclient

	// add additional ent dependencies
	entOpts = append(
		entOpts,
		ent.Authz(*fgaClient),
		ent.Emails(so.Config.Handler.EmailManager),
		ent.Marionette(so.Config.Handler.TaskMan),
		ent.Analytics(&analytics),
	)

	// Setup DB connection
	dbConfig := entdb.NewDBConfig(so.Config.Settings.DB, logger)

	entdbClient, err = dbConfig.NewMultiDriverDBClient(ctx, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// Setup Redis connection
	redisClient := cache.New(so.Config.Settings.Redis)
	defer redisClient.Close()

	// Add Driver to the Handlers Config
	so.Config.Handler.DBClient = entdbClient

	// Add redis client to Handlers Config
	so.Config.Handler.RedisClient = redisClient

	// add ready checks
	so.AddServerOptions(
		serveropts.WithReadyChecks(dbConfig, fgaClient, redisClient),
	)

	// add authz options
	so.AddServerOptions(
		serveropts.WithAuth(),
	)

	// Create keys for development
	if so.Config.Server.Dev {
		so.AddServerOptions(serveropts.WithGeneratedKeys())
	}

	// add session manager
	so.AddServerOptions(
		serveropts.WithSessionManager(redisClient),
	)

	srv := server.NewServer(so.Config, so.Config.Logger)

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient))

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
