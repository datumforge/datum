package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/api"
	ent "github.com/datumforge/datum/internal/ent/generated"
)

const (
	defaultListenAddr          = ":17608"
	defaultShutdownGracePeriod = 5 * time.Second
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

	serveCmd.Flags().Bool("debug", false, "enable server debug")
	viperBindFlag("server.debug", serveCmd.Flags().Lookup("debug"))

	serveCmd.Flags().String("listen", defaultListenAddr, "address to listen on")
	viperBindFlag("server.listen", serveCmd.Flags().Lookup("listen"))

	serveCmd.Flags().Duration("shutdown-grace-period", defaultShutdownGracePeriod, "server shutdown grace period")
	viperBindFlag("server.shutdown-grace-period", serveCmd.Flags().Lookup("shutdown-grace-period"))

	// only available as a CLI arg because these should only be used in dev environments
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) error {
	if serveDevMode {
		enablePlayground = true
	}

	// TODO add db connection

	cOpts := []ent.Option{}

	if viper.GetBool(("debug")) {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)
	defer client.Close()

	// TODO uncomment after db setup
	// // Run the automatic migration tool to create all schema resources.
	// if err := client.Schema.Create(ctx); err != nil {
	// 	logger.Errorf("failed creating schema resources", zap.Error(err))
	// 	return err
	// }

	// TODO jwt auth middleware

	var mw []echo.MiddlewareFunc

	r := api.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground, mw...)

	srv := echo.New()

	srv.Use(middleware.RequestID())
	srv.Use(middleware.Recover())

	// add logging
	zapLogger, _ := zap.NewProduction()
	srv.Use(echozap.ZapLogger(zapLogger))

	srv.Debug = viper.GetBool("server.debug")

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
