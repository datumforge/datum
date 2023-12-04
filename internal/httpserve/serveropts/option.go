package serveropts

import (
	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/datum/internal/httpserve/middleware"
)

type (
	ServerOption interface {
		apply(*ServerOptions)
	}
)

// WithConfigProvider supplies the config for the UI server
func WithConfigProvider(cfgProvider config.ConfigProvider) ServerOption {
	return newApplyFuncContainer(func(s *ServerOptions) {
		s.ConfigProvider = cfgProvider
	})
}

// WithAPIMiddleware supplies API middleware
func WithAPIMiddleware(middleware []middleware.MiddlewareFunc) ServerOption {
	return newApplyFuncContainer(func(s *ServerOptions) {
		s.APIMiddleware = append(s.APIMiddleware, middleware...)
	})
}

type (
	applyFuncContainer struct {
		applyInternal func(*ServerOptions)
	}
)

func (fso *applyFuncContainer) apply(s *ServerOptions) {
	fso.applyInternal(s)
}

func newApplyFuncContainer(apply func(option *ServerOptions)) *applyFuncContainer {
	return &applyFuncContainer{
		applyInternal: apply,
	}
}
