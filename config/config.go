package config

import (
	"crypto/tls"
	"strings"
	"time"

	"github.com/datumforge/entx"
	"github.com/datumforge/fgax"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/mcuadros/go-defaults"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/analytics/posthog"
	"github.com/datumforge/datum/pkg/cache"
	"github.com/datumforge/datum/pkg/otelx"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/sentry"
)

var (
	DefaultConfigFilePath = "./config/.config.yaml"
)

// Config contains the configuration for the datum server
type Config struct {
	// RefreshInterval determines how often to reload the config
	RefreshInterval time.Duration `json:"refresh_interval" koanf:"refresh_interval" default:"10m"`

	// Server contains the echo server settings
	Server Server `json:"server" koanf:"server"`

	// Auth contains the authentication token settings and provider(s)
	Auth Auth `json:"auth" koanf:"auth"`

	// Authz contains the authorization settings for fine grained access control
	Authz fgax.Config `json:"authz" koanf:"authz"`

	// DB contains the database configuration for the ent client
	DB entx.Config `json:"db" koanf:"db"`

	// Redis contains the redis configuration for the key-value store
	Redis cache.Config `json:"redis" koanf:"redis"`

	// Tracer contains the tracing config for opentelemetry
	Tracer otelx.Config `json:"tracer" koanf:"tracer"`

	// Email contains email sending configuration for the server
	Email emails.Config `json:"email" koanf:"email"`

	// Sessions config for user sessions and cookies
	Sessions sessions.Config `json:"sessions" koanf:"sessions"`

	// Sentry contains the sentry configuration for error tracking
	Sentry sentry.Config `json:"sentry" koanf:"sentry"`

	// PostHog contains the configuration for the PostHog analytics
	PostHog posthog.Config `json:"posthog" koanf:"posthog"`
}

// Server settings for the echo server
type Server struct {
	// Debug enables debug mode for the server
	Debug bool `json:"debug" koanf:"debug" default:"false"`
	// Dev enables echo's dev mode options
	Dev bool `json:"dev" koanf:"dev" default:"false"`
	// Listen sets the listen address to serve the echo server on
	Listen string `json:"listen" koanf:"listen" jsonschema:"required" default:":17608"`
	// ShutdownGracePeriod sets the grace period for in flight requests before shutting down
	ShutdownGracePeriod time.Duration `json:"shutdown_grace_period" koanf:"shutdown_grace_period" default:"10s"`
	// ReadTimeout sets the maximum duration for reading the entire request including the body
	ReadTimeout time.Duration `json:"read_timeout" koanf:"read_timeout" default:"15s"`
	// WriteTimeout sets the maximum duration before timing out writes of the response
	WriteTimeout time.Duration `json:"write_timeout" koanf:"write_timeout" default:"15s"`
	// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled
	IdleTimeout time.Duration `json:"idle_timeout" koanf:"idle_timeout" default:"30s"`
	// ReadHeaderTimeout sets the amount of time allowed to read request headers
	ReadHeaderTimeout time.Duration `json:"read_header_timeout" koanf:"read_header_timeout" default:"2s"`
	// TLS contains the tls configuration settings
	TLS TLS `json:"tls" koanf:"tls"`
	// CORS contains settings to allow cross origin settings and insecure cookies
	CORS CORS `json:"cors" koanf:"cors"`
}

// Auth settings including oauth2 providers and datum token configuration
type Auth struct {
	// Enabled authentication on the server, not recommended to disable
	Enabled bool `json:"enabled" koanf:"enabled" default:"true"`
	// Token contains the token config settings for Datum issued tokens
	Token tokens.Config `json:"token" koanf:"token" jsonschema:"required" alias:"tokenconfig"`
	// SupportedProviders are the supported oauth providers that have been configured
	SupportedProviders []string `json:"supported_providers" koanf:"supported_providers"`
	// Providers contains supported oauth2 providers configuration
	Providers handlers.OauthProviderConfig `json:"providers" koanf:"providers"`
}

// CORS settings for the server to allow cross origin requests
type CORS struct {
	// AllowOrigins is a list of allowed origin to indicate whether the response can be shared with
	// requesting code from the given origin
	AllowOrigins []string `json:"allow_origins" koanf:"allow_origins"`
	// CookieInsecure allows CSRF cookie to be sent to servers that the browser considers
	// unsecured. Useful for cases where the connection is secured via VPN rather than
	// HTTPS directly.
	CookieInsecure bool `json:"cookie_insecure" koanf:"cookie_insecure"`
}

// TLS settings for the server for secure connections
type TLS struct {
	// Config contains the tls.Config settings
	Config *tls.Config `json:"config" koanf:"config" jsonschema:"-"`
	// Enabled turns on TLS settings for the server
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// CertFile location for the TLS server
	CertFile string `json:"cert_file" koanf:"cert_file" default:"server.crt"`
	// CertKey file location for the TLS server
	CertKey string `json:"cert_key" koanf:"cert_key" default:"server.key"`
	// AutoCert generates the cert with letsencrypt, this does not work on localhost
	AutoCert bool `json:"auto_cert" koanf:"auto_cert" default:"false"`
}

// Load is responsible for loading the configuration from a YAML file and environment variables.
// If the `cfgFile` is empty or nil, it sets the default configuration file path.
// Config settings are taken from default values, then from the config file, and finally from environment
// the later overwriting the former.
func Load(cfgFile *string) (*Config, error) {
	k := koanf.New(".")

	if cfgFile == nil || *cfgFile == "" {
		*cfgFile = DefaultConfigFilePath
	}

	// load defaults
	conf := &Config{}
	defaults.SetDefaults(conf)

	// parse yaml config
	if err := k.Load(file.Provider(*cfgFile), yaml.Parser()); err != nil {
		panic(err)
	}

	// unmarshal the config
	if err := k.Unmarshal("", &conf); err != nil {
		panic(err)
	}

	// load env vars
	if err := k.Load(env.Provider("DATUM_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, "DATUM_")), "_", ".")
	}), nil); err != nil {
		panic(err)
	}

	// unmarshal the env vars
	if err := k.Unmarshal("", &conf); err != nil {
		panic(err)
	}

	return conf, nil
}
