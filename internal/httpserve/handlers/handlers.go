package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	handlers []handler
}

// CheckFunc is a function that can be used to check the status of a service
type CheckFunc func(ctx context.Context) error

type handler interface {
	Routes(*echo.Group)
}

func NewHandlers() *Handler {
	h := []handler{}

	return &Handler{handlers: h}
}

// AddHandler provides the ability to add additional HTTP handlers that process
// requests. The handler that is provided should have a Routes(*echo.Group)
// function, which allows the routes to be added to the server.
func (s *Handler) AddHandler(h handler) *Handler {
	s.handlers = append(s.handlers, h)

	return s
}

// Handler returns a new http.Handler for serving requests.
func HealthHandler() http.Handler {
	// Health endpoints
	s.GET("/livez", s.livenessCheckHandler)
	s.GET("/readyz", s.readinessCheckHandler)
	s.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	for _, handler := range s.handlers {
		handler.Routes(srv.Group(""))
	}

	return s
}
