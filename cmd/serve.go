package cmd

import (
	"context"
	"log"

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

	so := serveropts.NewServerOptions(serverOpts)

	// setup Authz connection
	// this must come before the database setup because the FGA Client
	// is used as an ent dependency
	if so.Config.Authz.Enabled {
		config := fga.NewAuthzConfig(so.Config.Authz, logger)

		fgaClient, err = fga.CreateFGAClientWithStore(ctx, *config)
		if err != nil {
			return err
		}

		// add client as ent dependency
		entOpts = append(entOpts, ent.Authz(*fgaClient))

		// add jwt middleware
		secretKey := []byte(viper.GetString("jwt.secretkey"))
		jwtMiddleware := auth.CreateJwtMiddleware([]byte(secretKey))

		mw = append(mw, jwtMiddleware)
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

	srv := server.NewServer(so.Config.Server, so.Config.Logger.Desugar())

	// Setup Graph API Handlers
	r := graphapi.NewResolver(entdbClient, so.Config.Auth.Enabled).
		WithLogger(logger.Named("resolvers"))

	handler := r.Handler(viper.GetBool("server.dev"), mw...)

	// Add Graph Handler
	srv.AddHandler(handler)

	if err := srv.StartEchoServer(); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return nil
}
