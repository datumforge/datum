package serveropts

import (
	"github.com/datumforge/datum/internal/server/config"
	"github.com/datumforge/datum/internal/server/middleware"
)

type (
	ServerOptions struct {
		ConfigProvider config.ConfigProvider
		APIMiddleware  []middleware.MiddlewareFunc
	}
)

func NewServerOptions(opts []ServerOption) *ServerOptions {
	so := &ServerOptions{
		// Set defaults here.
	}

	for _, opt := range opts {
		opt.apply(so)
	}

	return so
}
