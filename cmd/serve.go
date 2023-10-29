package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/brpaz/echozap"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/api"
	"github.com/datumforge/datum/internal/echox"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/entdb"
)

const (
	defaultListenAddr            = ":17608"
	defaultShutdownGracePeriod   = 5 * time.Second
	defaultDBPrimaryURI          = "datum.db?mode=memory&_fk=1"
	defaultDBSecondaryURI        = "backup.db?mode=memory&_fk=1"
	defaultFGAScheme             = "https"
	defaultFGAHost               = ""
	defaultOIDCJWKSRemoteTimeout = 5 * time.Second
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

	serveCmd.Flags().Duration("shutdown-grace-period", defaultShutdownGracePeriod, "server shutdown grace period")
	viperBindFlag("server.shutdown-grace-period", serveCmd.Flags().Lookup("shutdown-grace-period"))

	// Database flags
	serveCmd.Flags().Bool("db-mutli-write", true, "write to a primary and secondary database")
	viperBindFlag("server.db.multi-write", serveCmd.Flags().Lookup("db-mutli-write"))

	serveCmd.Flags().String("db-primary", defaultDBPrimaryURI, "db primary uri")
	viperBindFlag("server.db-primary", serveCmd.Flags().Lookup("db-primary"))

	serveCmd.Flags().String("db-secondary", defaultDBSecondaryURI, "db secondary uri")
	viperBindFlag("server.db-secondary", serveCmd.Flags().Lookup("db-secondary"))

	// echo-jwt flags
	serveCmd.Flags().String("jwt-secretkey", "", "secret key for echojwt config")
	viperBindFlag("jtw.secretkey", serveCmd.Flags().Lookup("jwt-secretkey"))

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
	serveCmd.Flags().String("fgaHost", defaultFGAHost, "fga host without the scheme (e.g. api.fga.example instead of https://api.fga.example)")
	viperBindFlag("fga.host", serveCmd.Flags().Lookup("fgaHost"))

	serveCmd.Flags().String("fgaScheme", defaultFGAScheme, "fga scheme")
	viperBindFlag("fga.scheme", serveCmd.Flags().Lookup("fgaScheme"))

	serveCmd.Flags().String("fgaStoreID", "", "fga store ID")
	viperBindFlag("fga.storeID", serveCmd.Flags().Lookup("fgaStoreID"))

	// only available as a CLI arg because these should only be used in dev environments
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		client *ent.Client
		err    error
	)

	if viper.GetBool("server.db.multi-write") {
		client, err = newMultiDriverDBClient(ctx)
		if err != nil {
			return err
		}
	} else {
		client, err = newEntDBDriver(viper.GetString("server.db-primary"), dialect.SQLite)
		if err != nil {
			return err
		}
	}
	defer client.Close()

	var mw []echo.MiddlewareFunc

	srv := echo.New()
	srv.Use(middleware.RequestID())
	srv.Use(middleware.Recover())

	// dev mode settings
	if serveDevMode {
		enablePlayground = true

		srv.Use(middleware.CORS())
	}

	// add logging
	zapLogger, _ := zap.NewProduction()
	srv.Use(echozap.ZapLogger(zapLogger))

	srv.Debug = viper.GetBool("server.debug")

	// add jwt middleware
	if viper.GetBool("oidc.enabled") {
		secretKey := viper.GetString("jtw.secretkey")
		jwtConfig := createJwtMiddleware([]byte(secretKey))

		mw = append(mw, jwtConfig)
	}

	// Add echo context to middleware
	srv.Use(echox.EchoContextToContextMiddleware())
	mw = append(mw, echox.EchoContextToContextMiddleware())

	r := api.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground, mw...)

	handler.Routes(srv.Group(""))

	listener, err := net.Listen("tcp", viper.GetString("server.listen"))
	if err != nil {
		return err
	}

	defer listener.Close() //nolint:errcheck // No need to check error.

	logger.Info("starting server")

	s := &http.Server{
		Handler: srv.Server.Handler,
	}

	var (
		exit = make(chan error, 1)
		quit = make(chan os.Signal, 2) //nolint:gomnd
	)

	// Serve in a go routine.
	// If serve returns an error, capture the error to return later.
	go func() {
		if err := s.Serve(listener); err != nil {
			exit <- err

			return
		}

		exit <- nil
	}()

	// close server to kill active connections.
	defer s.Close() //nolint:errcheck // server is being closed, we'll ignore this.

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-exit:
		return err
	case sig := <-quit:
		logger.Warn(fmt.Sprintf("%s received, server shutting down", sig.String()))
	case <-ctx.Done():
		logger.Warn("context done, server shutting down")

		// Since the context has already been canceled, the server would immediately shutdown.
		// We'll reset the context to allow for the proper grace period to be given.
		ctx = context.Background()
	}

	ctx, cancel := context.WithTimeout(ctx, viper.GetDuration("server.shutdown-grace-period"))
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown timed out", zap.Error(err))

		return err
	}

	return nil
}

// newEntDB creates returns new sql db connection
func newEntDB(dataSource, dbDriverName string) (*entsql.Driver, error) {
	// setup db connection
	db, err := sql.Open(dbDriverName, dataSource)
	if err != nil {
		return nil, fmt.Errorf("failed connecting to database: %w", err)
	}

	// verify db connection using ping
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed verifying database connection: %w", err)
	}

	return entsql.OpenDB(dialect.SQLite, db), nil
}

func newEntDBDriver(dataSource, driver string) (*ent.Client, error) {
	db, err := newEntDB(dataSource, driver)
	if err != nil {
		return nil, err
	}

	cOpts := []ent.Option{ent.Driver(db)}

	if viper.GetBool(("debug")) {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)

	return client, nil
}

func newMultiDriverDBClient(ctx context.Context) (*ent.Client, error) {
	debug := viper.GetBool(("debug"))

	primaryDB, err := newEntDB(viper.GetString("server.db-primary"), dialect.SQLite)
	if err != nil {
		return nil, err
	}

	if err := createSchema(ctx, primaryDB, debug); err != nil {
		return nil, err
	}

	secondaryDB, err := newEntDB(viper.GetString("server.db-secondary"), dialect.SQLite)
	if err != nil {
		return nil, err
	}

	if err := createSchema(ctx, secondaryDB, debug); err != nil {
		return nil, err
	}

	// Create Multiwrite driver
	cOpts := []ent.Option{ent.Driver(&entdb.MultiWriteDriver{Wp: primaryDB, Ws: secondaryDB})}
	if debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)

	return client, nil
}

func createEntDBClient(db *entsql.Driver, debug bool) *ent.Client {
	cOpts := []ent.Option{ent.Driver(db)}

	if debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	return ent.NewClient(cOpts...)
}

func createSchema(ctx context.Context, db *entsql.Driver, debug bool) error {
	client := createEntDBClient(db, debug)

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Errorf("failed creating schema resources", zap.Error(err))

		return err
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
