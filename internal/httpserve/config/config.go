package config

import (
	"crypto/tls"
	"net/http"
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/fgax"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	"github.com/datumforge/datum/internal/cache"
	"github.com/datumforge/datum/internal/entdb"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/otelx"
	"github.com/datumforge/datum/internal/sessions"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/internal/utils/emails"
)

var (
	// DefaultConfigRefresh sets the default interval to refresh the config.
	DefaultConfigRefresh = 10 * time.Minute
	// DefaultTLSConfig is the default TLS config used when HTTPS is enabled
	DefaultTLSConfig = &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
)

type (
	// Config contains the configuration for the datum server
	Config struct {
		// RefreshInterval holds often to reload the config
		RefreshInterval time.Duration `yaml:"refreshInterval" split_words:"true" default:"10m"`

		// Server contains the echo server settings
		Server Server `yaml:"server"`

		// Auth contains the authentication provider(s)
		Auth Auth `yaml:"auth"`

		// Authz contains the authorization settings
		Authz fgax.Config `yaml:"authz"`

		// DB contains the database configuration
		DB entdb.Config `yaml:"db"`

		// RedisConfig contains the redis configuration
		RedisConfig cache.Config `yaml:"redisConfig"`

		// Logger contains the logger used by echo functions
		Logger *zap.SugaredLogger `yaml:"logger"`

		// Tracer contains the tracing config used by datum
		Tracer otelx.Config `yaml:"tracer"`

		// Email contains email sending configuration
		Email emails.Config `yaml:"email"`
	}

	// Server settings
	Server struct {
		// Debug enables echo's Debug option.
		Debug bool `yaml:"debug" split_words:"true" default:"false"`
		// Dev enables echo's dev mode options.
		Dev bool `yaml:"dev" split_words:"true" default:"false"`
		// Listen sets the listen address to serve the echo server on.
		Listen string `yaml:"listen" split_words:"true" default:":17608"`
		// ShutdownGracePeriod sets the grace period for in flight requests before shutting down.
		ShutdownGracePeriod time.Duration `yaml:"shutdownGracePeriod" split_words:"true" default:"10s"`
		// ReadTimeout sets the maximum duration for reading the entire request including the body.
		ReadTimeout time.Duration `yaml:"readTimeout" split_words:"true" default:"15s"`
		// WriteTimeout sets the maximum duration before timing out writes of the response.
		WriteTimeout time.Duration `yaml:"writeTimeout" split_words:"true" default:"15s"`
		// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled.
		IdleTimeout time.Duration `yaml:"idleTimeout" split_words:"true" default:"30s"`
		// ReadHeaderTimeout sets the amount of time allowed to read request headers.
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout" split_words:"true" default:"2s"`
		// TLS contains the tls configuration settings
		TLS TLS `yaml:"tls"`
		// CORS contains settings to allow cross origin settings and insecure cookies
		CORS CORS `yaml:"cors"`
		// Routes contains the handler functions
		Routes []http.Handler `yaml:"routes"`
		// Middleware to enable on the echo server
		Middleware []echo.MiddlewareFunc `yaml:"middleware"`
		// Handler contains the required settings for REST handlers including ready checks and JWT keys
		Handler handlers.Handler `yaml:"checks"`
		// Token contains the token config settings
		Token tokens.Config `yaml:"token"`
		// SM manages sessions for users
		SM sessions.Store[map[string]string]
	}

	// Auth settings including providers and the ability to enable/disable auth all together
	Auth struct {
		// Enabled - checks this first before reading your provider config
		Enabled bool `yaml:"enabled" split_words:"true" default:"true"`
	}

	// CORS settings
	CORS struct {
		// AllowOrigins is a list of allowed origin to indicate whether the response can be shared with
		// requesting code from the given origin
		AllowOrigins []string `yaml:"allowOrigins"`
		// CookieInsecure allows CSRF cookie to be sent to servers that the browser considers
		// unsecured. Useful for cases where the connection is secured via VPN rather than
		// HTTPS directly.
		CookieInsecure bool `yaml:"cookieInsecure"`
	}

	// TLS settings
	TLS struct {
		// Config contains the tls.Config settings
		Config *tls.Config `yaml:"config"`
		// Enabled turns on TLS settings for the server
		Enabled bool `yaml:"enabled" split_words:"true" default:"false"`
		// CertFile location for the TLS server
		CertFile string `yaml:"certFile" split_words:"true" default:"server.crt"`
		// CertKey file location for the TLS server
		CertKey string `yaml:"certKey" split_words:"true" default:"server.key"`
		// AutoCert generates the cert with letsencrypt, this does not work on localhost
		AutoCert bool `yaml:"autoCert" split_words:"true" default:"false"`
	}
)

// Ensure that *Config implements ConfigProvider interface.
var _ ConfigProvider = &Config{}

// GetConfig implements ConfigProvider.
func (c *Config) GetConfig() (*Config, error) {
	return c, nil
}

// NewServerConfig creates a new empty config
func NewServerConfig() *Config {
	// tls settings
	t := &TLS{}

	// load defaults and env vars
	if err := envconfig.Process("datum_tls", t); err != nil {
		panic(err)
	}

	s := &Server{}

	// load defaults and env vars
	if err := envconfig.Process("datum_server", s); err != nil {
		panic(err)
	}

	s.TLS = *t

	return &Config{
		Server: *s,
	}
}

// WithTLSDefaults sets tls default settings assuming a default cert and key file location.
func (c Config) WithTLSDefaults() Config {
	c.WithDefaultTLSConfig()

	return c
}

// WithDefaultTLSConfig sets the default TLS Configuration
func (c Config) WithDefaultTLSConfig() Config {
	c.Server.TLS.Enabled = true
	c.Server.TLS.Config = DefaultTLSConfig

	return c
}

// WithTLSCerts sets the TLS Cert and Key locations
func (c *Config) WithTLSCerts(certFile, certKey string) *Config {
	c.Server.TLS.CertFile = certFile
	c.Server.TLS.CertKey = certKey

	return c
}

// WithAutoCert generates a letsencrypt certificate, a valid host must be provided
func (c *Config) WithAutoCert(host string) *Config {
	autoTLSManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Cache certificates to avoid issues with rate limits (https://letsencrypt.org/docs/rate-limits)
		Cache:      autocert.DirCache("/var/www/.cache"),
		HostPolicy: autocert.HostWhitelist(host),
	}

	c.Server.TLS.Enabled = true
	c.Server.TLS.Config = DefaultTLSConfig

	c.Server.TLS.Config.GetCertificate = autoTLSManager.GetCertificate
	c.Server.TLS.Config.NextProtos = []string{acme.ALPNProto}

	return c
}
