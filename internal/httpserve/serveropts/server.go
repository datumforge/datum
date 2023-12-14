package serveropts

import (
	"github.com/datumforge/datum/internal/httpserve/config"
)

type ServerOptions struct {
	ConfigProvider config.ConfigProvider
	Config         config.Config
}

func NewServerOptions(opts []ServerOption) *ServerOptions {
	so := &ServerOptions{
		Config: config.Config{
			RefreshInterval: config.DefaultConfigRefresh,
		},
	}

	for _, opt := range opts {
		opt.apply(so)
	}

	return so
}

// AddServerOptions applies a server option after the initial setup
// this should be used when information is not available on NewServerOptions
func (so *ServerOptions) AddServerOptions(opt ServerOption) {
	opt.apply(so)
}
