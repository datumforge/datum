package config

import (
	"crypto/tls"
	"net/http"
	"time"

	echo "github.com/datumforge/echox"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"

	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/tokens"
)

type (
	// Config contains the configuration for the datum server
	Config struct {
		// RefreshInterval holds often to reload the config
		RefreshInterval time.Duration `yaml:"refreshInterval"`

		// Server contains the echo server settings
		Server Server `yaml:"server"`

		// Auth contains the authentication provider(s)
		Auth Auth `yaml:"auth"`

		// Authz contains the authorization settings
		Authz fga.Config `yaml:"authz"`

		// DB contains the database configuration
		DB DB `yaml:"auth"`

		// Logger contains the logger used by echo functions
		Logger *zap.SugaredLogger `yaml:"logger"`
	}

	// Server settings
	Server struct {
		// Debug enables echo's Debug option.
		Debug bool `yaml:"debug"`
		// Dev enables echo's dev mode options.
		Dev bool `yaml:"dev"`
		// Listen sets the listen address to serve the echo server on.
		Listen string
		// ShutdownGracePeriod sets the grace period for in flight requests before shutting down.
		ShutdownGracePeriod time.Duration `yaml:"shutdownGracePeriod"`
		// ReadTimeout sets the maximum duration for reading the entire request including the body.
		ReadTimeout time.Duration `yaml:"readTimeout"`
		// WriteTimeout sets the maximum duration before timing out writes of the response.
		WriteTimeout time.Duration `yaml:"writeTimeout"`
		// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled.
		IdleTimeout time.Duration `yaml:"idleTimeout"`
		// ReadHeaderTimeout sets the amount of time allowed to read request headers.
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
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
	}

	// DB Settings
	DB struct {
		// Debug to print debug database logs
		Debug bool
		// SQL Driver name from dialect.Driver
		DriverName string
		// MultiWrite enabled writing to two databases
		MultiWrite bool
		// Primary write database source (required)
		PrimaryDBSource string
		// Secondary write databsae source (optional)
		SecondaryDBSource string
	}

	// Auth settings including providers and the ability to enable/disable auth all together
	Auth struct {
		// Enabled - checks this first before reading your provider config
		Enabled bool `yaml:"enabled"`
		// JWTSigningKey contains a 32 byte array to sign with the HmacSha256 algorithms
		JWTSigningKey []byte `yaml:"jwtSigningKey"`
		// A list of auth providers. Currently enables only the first provider in the list.
		Providers []AuthProvider `yaml:"providers"`
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
		Enabled bool
		// CertFile location for the TLS server
		CertFile string
		// CertKey file location for the TLS server
		CertKey string
		// AutoCert generates the cert with letsencrypt, this does not work on localhost
		AutoCert bool
	}

	AuthProvider struct {
		// Label for the provider (optional)
		Label string `yaml:"label"`
		// Type of the auth provider, currently only OIDC is supported
		Type string `yaml:"type"`
		// OIDC .well-known/openid-configuration URL, ex. https://accounts.google.com/
		ProviderURL string `yaml:"providerUrl"`
		// IssuerURL is only needed when it differs from the ProviderURL (optional)
		IssuerURL string `yaml:"issuerUrl"`
		// ClientID of the oauth2 provider
		ClientID string `yaml:"clientId"`
		// ClientSecret is the private key that authenticates your integration when requesting an OAuth token (optional when using PKCE)
		ClientSecret string `yaml:"clientSecret"`
		// Scopes for authentication, typically [openid, profile, email]
		Scopes []string `yaml:"scopes"`
		// CallbackURL after a successful auth, e.g. https://localhost:8080/oauth/callback
		CallbackURL string `yaml:"callbackUrl"`
		// Options added as URL query params when redirecting to auth provider. Can be used to configure custom auth flows such as Auth0 invitation flow.
		Options map[string]interface{} `yaml:"options"`
	}
)

// Ensure that *Config implements ConfigProvider interface.
var _ ConfigProvider = &Config{}

// GetConfig implements ConfigProvider.
func (c *Config) GetConfig() (*Config, error) {
	return c, nil
}

// NewConfig creates a new empty config
func NewConfig() *Config {
	c := Config{}

	return &c
}

// SetDefaults sets default values if not already defined.
func (c *Config) SetDefaults() *Config {
	if c.Server.TLS.Enabled {
		// use 443 for secure servers as the default port
		c.Server.Listen = ":443"
		c.Server.TLS.Config = DefaultTLSConfig
	} else if c.Server.Listen == "" {
		// set default port if none is provided
		c.Server.Listen = ":8080"
	}

	if c.Server.ShutdownGracePeriod <= 0 {
		c.Server.ShutdownGracePeriod = DefaultShutdownGracePeriod
	}

	if c.Server.ReadTimeout <= 0 {
		c.Server.ReadTimeout = DefaultReadTimeout
	}

	if c.Server.WriteTimeout <= 0 {
		c.Server.WriteTimeout = DefaultWriteTimeout
	}

	if c.Server.IdleTimeout <= 0 {
		c.Server.IdleTimeout = DefaultIdleTimeout
	}

	if c.Server.ReadHeaderTimeout <= 0 {
		c.Server.ReadHeaderTimeout = DefaultReadHeaderTimeout
	}

	return c
}

// WithDebug enables echo's Debug option.
func (c *Config) WithDebug(debug bool) *Config {
	c.Server.Debug = debug

	return c
}

// WithDev enables echo's dev mode options.
func (c *Config) WithDev(dev bool) *Config {
	c.Server.Dev = dev

	return c
}

// WithListen sets the listen address to serve the echo server on.
func (c *Config) WithListen(listen string) *Config {
	c.Server.Listen = listen

	return c
}

// WithHTTPS enables https server options
func (c *Config) WithHTTPS(https bool) *Config {
	c.Server.TLS.Enabled = https

	return c
}

// WithTLSDefaults sets tls default settings assuming a default cert and key file location.
func (c Config) WithTLSDefaults() Config {
	c.WithDefaultTLSConfig()
	c.Server.TLS.CertFile = DefaultCertFile
	c.Server.TLS.CertKey = DefaultKeyFile

	return c
}

// WithShutdownGracePeriod sets the grace period for in flight requests before shutting down.
func (c *Config) WithShutdownGracePeriod(period time.Duration) *Config {
	c.Server.ShutdownGracePeriod = period

	return c
}

// WithDefaultReadTimeout sets the maximum duration for reading the entire request including the body.
func (c *Config) WithDefaultReadTimeout(period time.Duration) *Config {
	c.Server.ReadTimeout = period

	return c
}

// WithWriteTimeout sets the maximum duration before timing out writes of the response.
func (c *Config) WithWriteTimeout(period time.Duration) *Config {
	c.Server.WriteTimeout = period

	return c
}

// WithIdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled.
func (c *Config) WithIdleTimeout(period time.Duration) *Config {
	c.Server.IdleTimeout = period

	return c
}

// WithReadHeaderTimeout sets the amount of time allowed to read request headers.
func (c *Config) WithReadHeaderTimeout(period time.Duration) *Config {
	c.Server.ReadHeaderTimeout = period

	return c
}

// WithMiddleware includes the provided middleware when echo is initialized.
func (c Config) WithMiddleware(mdw ...echo.MiddlewareFunc) Config {
	c.Server.Middleware = append(c.Server.Middleware, mdw...)

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
