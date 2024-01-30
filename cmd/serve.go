package cmd

import (
	"context"

	_ "github.com/lib/pq"           // postgres driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/cache"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/config"
	authmw "github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/httpserve/serveropts"
	"github.com/datumforge/datum/internal/otelx"
	"github.com/datumforge/datum/internal/sessions"
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
		serveropts.WithRedisCache(),
		serveropts.WithAuth(),
		serveropts.WithFGAAuthz(),
		serveropts.WithTracer(),
		serveropts.WithEmailManager(),
		serveropts.WithTaskManager(),
		serveropts.WithSessionManager(),
	)

	so := serveropts.NewServerOptions(serverOpts)

	err = otelx.NewTracer(so.Config.Tracer, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

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
		conf := authmw.NewAuthOptions(
			authmw.WithAudience(so.Config.Server.Token.Audience),
			authmw.WithIssuer(so.Config.Server.Token.Issuer),
			authmw.WithJWKSEndpoint(so.Config.Server.Token.JWKSEndpoint),
		)

		authMiddleware := authmw.Authenticate(conf)

		mw = append(mw, authMiddleware)
	}

	// add additional ent dependencies
	entOpts = append(
		entOpts,
		ent.Emails(so.Config.Server.Handler.EmailManager),
		ent.Marionette(so.Config.Server.Handler.TaskMan),
	)

	// Setup DB connection
	dbConfig := entdb.NewDBConfig(so.Config.DB, logger)

	entdbClient, err = dbConfig.NewMultiDriverDBClient(ctx, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// Setup Redis connection
	redisClient := cache.New(so.Config.RedisConfig)
	defer redisClient.Close()

	// add ready checks
	so.AddServerOptions(serveropts.WithReadyChecks(dbConfig, fgaClient, redisClient))

	// Add Driver to the Handlers Config
	so.Config.Server.Handler.DBClient = entdbClient

	// Add redis client to Handlers Config
	so.Config.Server.Handler.RedisClient = redisClient

	srv := server.NewServer(so.Config.Server, so.Config.Logger)

	// add session middleware, this has to be added after the authMiddleware so we have the user id
	// when we get to the session. this is also added here so its only added to the graph routes
	// REST routes are expected to add the session middleware, as required
	mw = append(mw, sessions.LoadAndSave(
		so.Config.Server.SM,
		redisClient,
		so.Config.Logger,
	))

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient, mw))

	if err := srv.StartEchoServer(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
