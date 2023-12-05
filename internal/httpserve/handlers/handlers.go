package handlers

import (
	"context"
	"net/http"

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
func (h *Handler) AddHandler(r handler) *Handler {
	h.handlers = append(h.handlers, r)

	return h
}

// AddRoutes returns a new http.Handler for serving requests with routes added
func (h *Handler) AddRoutes(srv *echo.Echo) http.Handler {
	// Health endpoints
	srv.GET("/livez", h.livenessCheckHandler)
	// srv.GET("/readyz", h.readinessCheckHandler)

	// Metrics endpoints
	srv.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	for _, handler := range h.handlers {
		handler.Routes(srv.Group(""))
	}

	return srv
}
