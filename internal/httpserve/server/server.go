package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/brpaz/echozap"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	echoprometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	// config contains the base server settings
	config config.Server
	// tls contains the tls settings
	tls config.TLS
	// logger contains the zap logger
	logger *zap.Logger
	// handlers contain echo http handlers
	handlers handlers.Handler
}

// HTTPSConfig contains HTTPS server settings
type HTTPSConfig struct {
	tlsConfig *tls.Config
	certFile  string
	certKey   string
}

// DefaultHTTPServer creates a default http Server with default timeouts
func (s *Server) DefaultHTTPServer() *http.Server {
	return &http.Server{
		ReadTimeout:       s.config.ReadTimeout,
		WriteTimeout:      s.config.WriteTimeout,
		IdleTimeout:       s.config.IdleTimeout,
		ReadHeaderTimeout: s.config.ReadHeaderTimeout,
	}
}

// DefaultEchoServer creates the default echo server with standard middleware
func (s *Server) DefaultEchoServer() *echo.Echo {
	srv := echo.New()

	// hides echo's startup banner
	srv.HideBanner = true
	srv.HidePort = true
	srv.Debug = s.config.Debug

	// add middleware
	// TODO (sfunk): go back and use the middleware packages
	srv.Use(middleware.RequestID())
	srv.Use(middleware.Recover())
	srv.Use(echoprometheus.MetricsMiddleware())
	srv.Use(echozap.ZapLogger(s.logger))

	return srv

}

// RunWithContext listens and serves the echo server on the configured address.
// See ServeWithContext for more details.
func (s *Server) RunWithContext(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.config.Listen)
	if err != nil {
		return err
	}

	defer listener.Close() //nolint:errcheck // No need to check error.

	if s.tls.Enabled {
		return s.ServeHTTPSWithContext(ctx, listener)
	}

	return s.ServeHTTPWithContext(ctx, listener)
}

// ServeHTTPWithContext serves an http server on the provided listener.
// Serve blocks until SIGINT or SIGTERM are signalled,
// or if the http serve fails.
// A graceful shutdown will be attempted
func (s *Server) ServeHTTPWithContext(ctx context.Context, listener net.Listener) error {
	logger := s.logger.With(zap.String("address", listener.Addr().String()))

	logger.Info("starting server")

	srv := s.DefaultHTTPServer()
	srv.Handler = s.DefaultEchoServer()

	var (
		exit = make(chan error, 1)
		quit = make(chan os.Signal, 2) //nolint:gomnd
	)

	// Serve in a go routine.
	// If serve returns an error, capture the error to return later.
	go func() {
		if err := srv.Serve(listener); err != nil {
			exit <- err

			return
		}

		exit <- nil
	}()

	// close server to kill active connections.
	defer srv.Close() //nolint:errcheck // server is being closed, we'll ignore this.

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var err error

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

	ctx, cancel := context.WithTimeout(ctx, s.config.ShutdownGracePeriod)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown timed out", zap.Error(err))

		return err
	}

	return nil
}

// ServeHTTPSWithContext serves an https server on the provided listener.
// Serve blocks until SIGINT or SIGTERM are signalled,
// or if the http serve fails.
// A graceful shutdown will be attempted
func (s *Server) ServeHTTPSWithContext(ctx context.Context, listener net.Listener) error {
	logger := s.logger.With(zap.String("address", listener.Addr().String()))

	logger.Info("starting https server")

	// TODO: Add ability to do HTTPS Redirect with middleware.HTTPSRedirect()
	srv := s.DefaultHTTPServer()
	srv.Handler = s.DefaultEchoServer()

	srv.TLSConfig = s.tls.Config
	srv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

	var (
		exit = make(chan error, 1)
		quit = make(chan os.Signal, 2) //nolint:gomnd
	)

	// Serve in a go routine.
	// If serve returns an error, capture the error to return later.
	go func() {
		if err := srv.ServeTLS(listener, s.tls.CertFile, s.tls.CertKey); err != nil {
			exit <- err

			return
		}

		exit <- nil
	}()

	// close server to kill active connections.
	defer srv.Close() //nolint:errcheck // server is being closed, we'll ignore this.

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var err error

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

	ctx, cancel := context.WithTimeout(ctx, s.config.ShutdownGracePeriod)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown timed out", zap.Error(err))

		return err
	}

	return nil
}
