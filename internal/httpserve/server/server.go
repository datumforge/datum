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

	echoprometheus "github.com/datumforge/echo-prometheus/v5"
	"github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/datumforge/echozap"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/route"
)

type Server struct {
	// config contains the base server settings
	config config.Server
	// logger contains the zap logger
	logger *zap.Logger
}

// HTTPSConfig contains HTTPS server settings
type HTTPSConfig struct {
	tlsConfig *tls.Config //nolint:unused
	certFile  string      //nolint:unused
	certKey   string      //nolint:unused
}

// NewServer returns a new Server configuration
func NewServer(c config.Server, l *zap.Logger) *Server {
	return &Server{
		config: c,
		logger: l,
	}
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
func (s *Server) DefaultEchoServer() (*echox.Echo, echox.StartConfig) {
	srv := echox.New()

	sc := echox.StartConfig{
		HideBanner: true,
		HidePort:   true,
		Address:    s.config.Listen,
	}

	// hides echo's startup banner
	srv.Debug = s.config.Debug

	// add middleware
	srv.Use(middleware.RequestID())
	srv.Use(middleware.Recover())
	srv.Use(echoprometheus.MetricsMiddleware())
	srv.Use(echozap.ZapLogger(s.logger))

	// add all configured middleware
	for _, m := range s.config.Middleware {
		srv.Use(m)
	}

	return srv, sc
}

// RunWithContext listens and serves the echo server on the configured address.
// See ServeWithContext for more details.
func (s *Server) RunWithContext(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.config.Listen)
	if err != nil {
		return err
	}

	defer listener.Close() //nolint:errcheck // No need to check error.

	if s.config.TLS.Enabled {
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
	echoServer, _ := s.DefaultEchoServer()
	srv.Handler = echoServer

	// Add routes to the server
	// TODO (sfunk): this seems weird, the server should be in the config maybe?
	if err := route.RegisterHandlers(echoServer); err != nil {
		return err
	}

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
	srv.Handler, _ = s.DefaultEchoServer()

	srv.TLSConfig = s.config.TLS.Config
	srv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

	var (
		exit = make(chan error, 1)
		quit = make(chan os.Signal, 2) //nolint:gomnd
	)

	// Serve in a go routine.
	// If serve returns an error, capture the error to return later.
	go func() {
		if err := srv.ServeTLS(listener, s.config.TLS.CertFile, s.config.TLS.CertKey); err != nil {
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
