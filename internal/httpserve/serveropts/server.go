package serveropts

import (
	"github.com/datumforge/datum/internal/httpserve/config"
)

type ServerOptions struct {
	ConfigProvider config.ConfigProvider
}

func NewServerOptions(opts []ServerOption) *ServerOptions {
	so := &ServerOptions{
		// Set defaults here.
	}

	for _, opt := range opts {
		opt.apply(so)
	}

	return so
}
