package server

import (
	"context"

	echo "github.com/datumforge/echox"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/httpserve/config"
	echodebug "github.com/datumforge/datum/internal/httpserve/middleware/debug"
	"github.com/datumforge/datum/internal/httpserve/route"
	"github.com/datumforge/datum/internal/tokens"
)

type Server struct {
	// config contains the base server settings
	config config.Server
	// logger contains the zap logger
	logger *zap.SugaredLogger
	// handlers contains additional handlers to register with the echo server
	handlers []handler
}

type handler interface {
	Routes(*echo.Group)
}

// AddHandler provides the ability to add additional HTTP handlers that process
// requests. The handler that is provided should have a Routes(*echo.Group)
// function, which allows the routes to be added to the server.
func (s *Server) AddHandler(r handler) {
	s.handlers = append(s.handlers, r)
}

// NewServer returns a new Server configuration
func NewServer(c config.Server, l *zap.SugaredLogger) *Server {
	return &Server{
		config: c,
		logger: l,
	}
}

// StartEchoServer creates and starts the echo server with configured middleware and handlers
func (s *Server) StartEchoServer(ctx context.Context) error {
	_ = sentry.Init(s.config.Sentry.ClientOptions())

	srv := echo.New()

	sc := echo.StartConfig{
		HideBanner:      true,
		HidePort:        true,
		Address:         s.config.Listen,
		GracefulTimeout: s.config.ShutdownGracePeriod,
		GracefulContext: ctx,
	}

	srv.Debug = s.config.Debug

	if srv.Debug {
		srv.Use(echodebug.BodyDump(s.logger))
	}

	for _, m := range s.config.DefaultMiddleware {
		srv.Use(m)
	}

	// Setup token manager
	tm, err := tokens.New(s.config.Token)
	if err != nil {
		return err
	}

	keys, err := tm.Keys()
	if err != nil {
		return err
	}

	// pass to the REST handlers
	s.config.Handler.JWTKeys = keys
	s.config.Handler.TM = tm

	// Add base routes to the server
	if err := route.RegisterRoutes(srv, &s.config.Handler); err != nil {
		return err
	}

	// Registers additional routes for the graph endpoints with middleware defined
	for _, handler := range s.handlers {
		handler.Routes(srv.Group("", s.config.GraphMiddleware...))
	}

	// Print routes on startup
	routes := srv.Router().Routes()
	for _, r := range routes {
		s.logger.Infow("registered route", "route", r.Path(), "method", r.Method())
	}

	// if TLS is enabled, start new echo server with TLS
	if s.config.TLS.Enabled {
		s.logger.Infow("starting in https mode")

		return sc.StartTLS(srv, s.config.TLS.CertFile, s.config.TLS.CertKey)
	}

	// otherwise, start without TLS
	return sc.Start(srv)
}
