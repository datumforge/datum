package server

import (
	echoprometheus "github.com/datumforge/echo-prometheus/v5"
	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/datumforge/echozap"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/route"
)

type Server struct {
	// config contains the base server settings
	config config.Server
	// logger contains the zap logger
	logger *zap.Logger

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
func NewServer(c config.Server, l *zap.Logger) *Server {
	return &Server{
		config: c,
		logger: l,
	}
}

// StartEchoServer creates and starts the echo server with configured middleware and handlers
func (s *Server) StartEchoServer() error {
	srv := echo.New()

	sc := echo.StartConfig{
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

	// add echo context to parent context
	srv.Use(echox.EchoContextToContextMiddleware())

	// add all configured middleware
	for _, m := range s.config.Middleware {
		srv.Use(m)
	}

	// Add base routes to the server
	if err := route.RegisterBaseRoutes(srv); err != nil {
		return err
	}

	// Registers additional routes
	for _, handler := range s.handlers {
		handler.Routes(srv.Group(""))
	}

	// Print routes on startup
	routes := srv.Router().Routes()
	for _, r := range routes {
		s.logger.Sugar().Infow("registered route", "route", r.Path(), "method", r.Method())
	}

	// if TLS is enabled, start new echo server with TLS
	if s.config.TLS.Enabled {
		return sc.StartTLS(srv, s.config.TLS.CertFile, s.config.TLS.CertKey)
	}

	// otherwise, start without TLS
	return sc.Start(srv)
}
