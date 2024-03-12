package secure

import (
	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// Config contains the types used in the mw middleware
type Config struct {
	Enabled               bool               `json:"enabled" koanf:"enabled" default:"true"`
	Skipper               middleware.Skipper `json:"-" koanf:"-"`
	XSSProtection         string             `json:"xssprotection" koanf:"xssprotection" default:"1; mode=block"`
	ContentTypeNosniff    string             `json:"contenttypenosniff" koanf:"contenttypenosniff" default:"nosniff"`
	XFrameOptions         string             `json:"xframeoptions" koanf:"xframeoptions" default:"SAMEORIGIN"`
	HSTSPreloadEnabled    bool               `json:"hstspreloadenabled" koanf:"hstspreloadenabled" default:"false"`
	HSTSMaxAge            int                `json:"hstsmaxage" koanf:"hstsmaxage" default:"31536000"`
	ContentSecurityPolicy string             `json:"contentsecuritypolicy" koanf:"contentsecuritypolicy" default:"default-src 'self'"`
	ReferrerPolicy        string             `json:"referrerpolicy" koanf:"referrerpolicy" default:"same-origin"`
	CSPReportOnly         bool               `json:"cspreportonly" koanf:"cspreportonly" default:"false"`
}

// DefaultConfig struct is a populated config struct that can be referenced if the default konaf configurations are not available
var DefaultConfig = Config{
	Enabled:               true,
	Skipper:               middleware.DefaultSkipper,
	XSSProtection:         "1; mode=block",
	ContentTypeNosniff:    "nosniff",
	XFrameOptions:         "SAMEORIGIN",
	HSTSPreloadEnabled:    false,
	HSTSMaxAge:            31536000, //nolint: gomnd
	ContentSecurityPolicy: "default-src 'self'",
	ReferrerPolicy:        "same-origin",
	CSPReportOnly:         false,
}

// Secure returns a secure middleware with default unless overridden via the config
func Secure(conf *Config) echo.MiddlewareFunc {
	if conf.Enabled {
		secureConfig := middleware.SecureConfig{
			XSSProtection:         conf.XSSProtection,
			ContentTypeNosniff:    conf.ContentTypeNosniff,
			XFrameOptions:         conf.XFrameOptions,
			HSTSPreloadEnabled:    conf.HSTSPreloadEnabled,
			HSTSMaxAge:            conf.HSTSMaxAge,
			ReferrerPolicy:        conf.ReferrerPolicy,
			CSPReportOnly:         conf.CSPReportOnly,
			ContentSecurityPolicy: conf.ContentSecurityPolicy,
			Skipper:               conf.Skipper,
		}

		return middleware.SecureWithConfig(secureConfig)
	}

	return nil
}
