package config

import (
	"crypto/tls"
	"fmt"
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
	"github.com/datumforge/datum/pkg/middleware/cachecontrol"
	"github.com/datumforge/datum/pkg/middleware/cors"
	"github.com/datumforge/datum/pkg/middleware/mime"
	"github.com/datumforge/datum/pkg/middleware/ratelimit"
	"github.com/datumforge/datum/pkg/middleware/redirect"
	"github.com/datumforge/datum/pkg/middleware/secure"
	"github.com/datumforge/datum/pkg/otelx"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
	"github.com/datumforge/datum/pkg/utils/emails"
	"github.com/datumforge/datum/pkg/utils/sentry"
	"github.com/datumforge/datum/pkg/utils/totp"
)

var (
	DefaultConfigFilePath = "./config/.config.yaml"
)

// Config contains the configuration for the datum server
type Config struct {
	// RefreshInterval determines how often to reload the config
	RefreshInterval time.Duration `json:"refreshInterval" koanf:"refreshInterval" default:"10m"`

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

	// TOTP contains the configuration for the TOTP provider
	TOTP totp.Config `json:"totp" koanf:"totp"`

	// Ratelimit contains the configuration for the rate limiter
	Ratelimit ratelimit.Config `json:"ratelimit" koanf:"ratelimit"`
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
	ShutdownGracePeriod time.Duration `json:"shutdownGracePeriod" koanf:"shutdownGracePeriod" default:"10s"`
	// ReadTimeout sets the maximum duration for reading the entire request including the body
	ReadTimeout time.Duration `json:"readTimeout" koanf:"readTimeout" default:"15s"`
	// WriteTimeout sets the maximum duration before timing out writes of the response
	WriteTimeout time.Duration `json:"writeTimeout" koanf:"writeTimeout" default:"15s"`
	// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled
	IdleTimeout time.Duration `json:"idleTimeout" koanf:"idleTimeout" default:"30s"`
	// ReadHeaderTimeout sets the amount of time allowed to read request headers
	ReadHeaderTimeout time.Duration `json:"readHeaderTimeout" koanf:"readHeaderTimeout" default:"2s"`
	// TLS contains the tls configuration settings
	TLS TLS `json:"tls" koanf:"tls"`
	// CORS contains settings to allow cross origin settings and insecure cookies
	CORS cors.Config `json:"cors" koanf:"cors"`
	// Secure contains settings for the secure middleware
	Secure secure.Config `json:"secure" koanf:"secure"`
	// Redirect contains settings for the redirect middleware
	Redirects redirect.Config `json:"redirect" koanf:"redirects"`
	// CacheControl contains settings for the cache control middleware
	CacheControl cachecontrol.Config `json:"cacheControl" koanf:"cacheControl"`
	// Mime contains settings for the mime middleware
	Mime mime.Config `json:"mime" koanf:"mime"`
}

// Auth settings including oauth2 providers and datum token configuration
type Auth struct {
	// Enabled authentication on the server, not recommended to disable
	Enabled bool `json:"enabled" koanf:"enabled" default:"true"`
	// Token contains the token config settings for Datum issued tokens
	Token tokens.Config `json:"token" koanf:"token" jsonschema:"required" alias:"tokenconfig"`
	// SupportedProviders are the supported oauth providers that have been configured
	SupportedProviders []string `json:"supportedProviders" koanf:"supportedProviders"`
	// Providers contains supported oauth2 providers configuration
	Providers handlers.OauthProviderConfig `json:"providers" koanf:"providers"`
}

// TLS settings for the server for secure connections
type TLS struct {
	// Config contains the tls.Config settings
	Config *tls.Config `json:"config" koanf:"config" jsonschema:"-"`
	// Enabled turns on TLS settings for the server
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// CertFile location for the TLS server
	CertFile string `json:"certFile" koanf:"certFile" default:"server.crt"`
	// CertKey file location for the TLS server
	CertKey string `json:"certKey" koanf:"certKey" default:"server.key"`
	// AutoCert generates the cert with letsencrypt, this does not work on localhost
	AutoCert bool `json:"autoCert" koanf:"autoCert" default:"false"`
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

	all := k.All()
	fmt.Printf("all: %+v\n", all)

	return conf, nil
}
