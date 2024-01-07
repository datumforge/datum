package server

import (
	"context"

	echoprometheus "github.com/datumforge/echo-prometheus/v5"
	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/datumforge/echozap"
	"go.uber.org/zap"

	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/middleware/cachecontrol"
	"github.com/datumforge/datum/internal/httpserve/middleware/cors"
	echodebug "github.com/datumforge/datum/internal/httpserve/middleware/debug"
	"github.com/datumforge/datum/internal/httpserve/middleware/echocontext"
	"github.com/datumforge/datum/internal/httpserve/middleware/mime"
	"github.com/datumforge/datum/internal/httpserve/middleware/ratelimit"
	"github.com/datumforge/datum/internal/httpserve/middleware/session"
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
	srv := echo.New()

	sc := echo.StartConfig{
		HideBanner:      true,
		HidePort:        true,
		Address:         s.config.Listen,
		GracefulTimeout: s.config.ShutdownGracePeriod,
		GracefulContext: ctx,
	}

	srv.Debug = s.config.Debug

	// default middleware
	defaultMW := []echo.MiddlewareFunc{}
	defaultMW = append(defaultMW,
		middleware.RequestID(), // add request id
		middleware.Recover(),   // recover server from any panic/fatal error gracefully
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, session=${header:Session}, auth=${header:Authorization}\n",
		}),
		echoprometheus.MetricsMiddleware(),           // add prometheus metrics
		echozap.ZapLogger(s.logger.Desugar()),        // add zap logger, middleware requires the "regular" zap logger
		echocontext.EchoContextToContextMiddleware(), // adds echo context to parent
		cors.New(),                               // add cors middleware
		mime.New(),                               // add mime middleware
		cachecontrol.New(),                       // add cache control middleware
		ratelimit.DefaultRateLimiter(),           // add ratelimit middleware
		middleware.Secure(),                      // add XSS middleware
		session.LoadAndSave(s.config.Handler.SM), // add session middleware
	)

	if srv.Debug {
		defaultMW = append(defaultMW, echodebug.BodyDump(s.logger))
	}

	for _, m := range defaultMW {
		srv.Use(m)
	}
	// add all configured middleware
	for _, m := range s.config.Middleware {
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
	s.config.Handler.CookieDomain = s.config.Token.CookieDomain

	// Add base routes to the server
	if err := route.RegisterRoutes(srv, &s.config.Handler); err != nil {
		return err
	}

	// Registers additional routes for the graph endpoints
	// to pass middleware only to graph routes, append here
	graphMw := []echo.MiddlewareFunc{}
	for _, handler := range s.handlers {
		handler.Routes(srv.Group("", graphMw...))
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
