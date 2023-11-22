package echox

import (
	"context"
	"crypto/tls"
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
	"go.uber.org/zap"
)

// CheckFunc is a function that can be used to check the status of a service
type CheckFunc func(ctx context.Context) error

// Server implements the HTTPS Server
type Server struct {
	debug           bool
	dev             bool
	listen          string
	https           bool
	httpsConfig     HTTPSConfig
	logger          *zap.Logger
	middleware      []echo.MiddlewareFunc
	readinessChecks map[string]CheckFunc
	handlers        []handler
	timeouts        serverTimeouts
}

type serverTimeouts struct {
	shutdownTimeout   time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
	readHeaderTimeout time.Duration
}

// HTTPSConfig contains HTTPS server settings
type HTTPSConfig struct {
	tlsConfig *tls.Config
	certFile  string
	certKey   string
}

// NewServer will return an opinionated echo server for processing API requests.
func NewServer(logger *zap.Logger, cfg Config) (*Server, error) {
	// setup echo server
	cfg = cfg.WithDefaults()

	t := serverTimeouts{
		shutdownTimeout:   cfg.ShutdownGracePeriod,
		readTimeout:       cfg.ReadTimeout,
		writeTimeout:      cfg.WriteTimeout,
		idleTimeout:       cfg.IdleTimeout,
		readHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	s := &Server{
		debug:           cfg.Debug,
		dev:             cfg.Dev,
		https:           cfg.HTTPS,
		listen:          cfg.Listen,
		logger:          logger.Named("echox"),
		middleware:      cfg.Middleware,
		readinessChecks: map[string]CheckFunc{},
		timeouts:        t,
	}

	if s.https {
		s.httpsConfig = HTTPSConfig{
			tlsConfig: cfg.TLSConfig.TLSConfig,
		}

		// add the cert files if not using autocert
		if !cfg.TLSConfig.AutoCert {
			s.httpsConfig.certFile = cfg.TLSConfig.CertFile
			s.httpsConfig.certKey = cfg.TLSConfig.CertKey
		}
	}

	return s, nil
}

type handler interface {
	Routes(*echo.Group)
}

// AddHandler provides the ability to add additional HTTP handlers that process
// requests. The handler that is provided should have a Routes(*echo.Group)
// function, which allows the routes to be added to the server.
func (s *Server) AddHandler(h handler) *Server {
	s.handlers = append(s.handlers, h)

	return s
}

// AddReadinessCheck will accept a function to be ran during calls to /readyz
// These functions should accept a context and only return an error. When adding
// a readiness check a name is also provided, this name will be used when returning
// the state of all the checks
func (s *Server) AddReadinessCheck(name string, f CheckFunc) *Server {
	s.readinessChecks[name] = f

	return s
}

// Handler returns a new http.Handler for serving requests.
func (s *Server) Handler() http.Handler {
	srv := echo.New()

	// add middleware
	srv.Use(middleware.RequestID())
	srv.Use(middleware.Recover())

	// hides echo's startup banner
	srv.HideBanner = true
	srv.HidePort = true

	// set CORS in dev mode
	if s.dev {
		srv.Use(middleware.CORS())
	}

	zapLogger, _ := zap.NewProduction()
	srv.Use(echozap.ZapLogger(zapLogger))

	// Add echo context to middleware
	srv.Use(EchoContextToContextMiddleware())

	srv.Debug = s.debug

	srv.Use(s.middleware...)

	// Health endpoints
	srv.GET("/livez", s.livenessCheckHandler)
	srv.GET("/readyz", s.readinessCheckHandler)

	for _, handler := range s.handlers {
		handler.Routes(srv.Group(""))
	}

	return srv
}

func (s *Server) defaultServer() *http.Server {
	return &http.Server{
		ReadTimeout:       s.timeouts.readTimeout,
		WriteTimeout:      s.timeouts.writeTimeout,
		IdleTimeout:       s.timeouts.idleTimeout,
		ReadHeaderTimeout: s.timeouts.readHeaderTimeout,
	}
}

// ServeHTTPWithContext serves an http server on the provided listener.
// Serve blocks until SIGINT or SIGTERM are signalled,
// or if the http serve fails.
// A graceful shutdown will be attempted
func (s *Server) ServeHTTPWithContext(ctx context.Context, listener net.Listener) error {
	logger := s.logger.With(zap.String("address", listener.Addr().String()))

	logger.Info("starting server")

	srv := s.defaultServer()
	srv.Handler = s.Handler()

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

	ctx, cancel := context.WithTimeout(ctx, s.timeouts.shutdownTimeout)
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
	srv := s.defaultServer()
	srv.Handler = s.Handler()
	srv.TLSConfig = s.httpsConfig.tlsConfig
	srv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

	var (
		exit = make(chan error, 1)
		quit = make(chan os.Signal, 2) //nolint:gomnd
	)

	// Serve in a go routine.
	// If serve returns an error, capture the error to return later.
	go func() {
		if err := srv.ServeTLS(listener, s.httpsConfig.certFile, s.httpsConfig.certKey); err != nil {
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

	ctx, cancel := context.WithTimeout(ctx, s.timeouts.shutdownTimeout)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		logger.Error("server shutdown timed out", zap.Error(err))

		return err
	}

	return nil
}

// RunWithContext listens and serves the echo server on the configured address.
// See ServeWithContext for more details.
func (s *Server) RunWithContext(ctx context.Context) error {
	listener, err := net.Listen("tcp", s.listen)
	if err != nil {
		return err
	}

	defer listener.Close() //nolint:errcheck // No need to check error.

	if s.https {
		return s.ServeHTTPSWithContext(ctx, listener)
	}

	return s.ServeHTTPWithContext(ctx, listener)
}
