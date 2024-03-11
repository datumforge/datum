package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/fgax"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/server"
	"github.com/datumforge/datum/internal/httpserve/serveropts"
	"github.com/datumforge/datum/pkg/cache"
	"github.com/datumforge/datum/pkg/otelx"
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

	serveCmd.PersistentFlags().String("config", "./config/.config.yaml", "config file location")
	viperBindFlag("config", serveCmd.PersistentFlags().Lookup("config"))
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		fgaClient *fgax.Client
		err       error
	)

	// create ent dependency injection
	entOpts := []ent.Option{ent.Logger(*logger)}

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithLogger(logger),
		serveropts.WithHTTPS(),
		serveropts.WithEmailManager(),
		serveropts.WithTaskManager(),
		serveropts.WithSentry(),
		serveropts.WithMiddleware(),
		serveropts.WithRateLimiter(),
		serveropts.WithAnalytics(),
	)

	so := serveropts.NewServerOptions(serverOpts, viper.GetString("config"))

	err = otelx.NewTracer(so.Config.Settings.Tracer, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	fgaClient, err = fgax.CreateFGAClientWithStore(ctx, so.Config.Settings.Authz, so.Config.Logger)
	if err != nil {
		return err
	}

	// Setup Redis connection
	redisClient := cache.New(so.Config.Settings.Redis)
	defer redisClient.Close()

	// add otp manager, after redis is setup
	so.AddServerOptions(
		serveropts.WithOTP(),
	)

	// add additional ent dependencies
	entOpts = append(
		entOpts,
		ent.Authz(*fgaClient),
		ent.Emails(so.Config.Handler.EmailManager),
		ent.Marionette(so.Config.Handler.TaskMan),
		ent.Analytics(so.Config.Handler.AnalyticsClient),
		ent.TOTP(so.Config.Handler.OTPManager),
	)

	// Setup DB connection
	entdbClient, dbConfig, err := entdb.NewMultiDriverDBClient(ctx, so.Config.Settings.DB, logger, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// Add Driver to the Handlers Config
	so.Config.Handler.DBClient = entdbClient

	// Add redis client to Handlers Config
	so.Config.Handler.RedisClient = redisClient

	// add ready checks
	so.AddServerOptions(
		serveropts.WithReadyChecks(dbConfig, fgaClient, redisClient),
	)

	// add auth options
	so.AddServerOptions(
		serveropts.WithAuth(),
	)

	// Create keys for development
	if so.Config.Settings.Server.Dev {
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
