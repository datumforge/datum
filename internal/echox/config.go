package echox

import (
	"crypto/tls"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	// DefaultShutdownGracePeriod sets the default for how long we give the sever
	// to shutdown before forcefully stopping the server.
	DefaultShutdownGracePeriod = 5 * time.Second
	// DefaultReadTimeout sets the default maximum duration for reading the entire request including the body.
	DefaultReadTimeout = 15 * time.Second
	// DefaultWriteTimeout sets the default maximum duration before timing out writes of the response.
	DefaultWriteTimeout = 15 * time.Second
	// DefaultIdleTimeout sets the default maximum amount of time to wait for the next request when keep-alives are enabled.
	DefaultIdleTimeout = 30 * time.Second
	// DefaultReadHeaderTimeout sets the default amount of time allowed to read request headers.
	DefaultReadHeaderTimeout = 2 * time.Second
	// DefaultCertFile is the default cert file location
	DefaultCertFile = "server.crt"
	// DefaultKeyFile is the default key file location
	DefaultKeyFile = "server.key"
)

// Config is used to configure a new echo server
type Config struct {
	// Debug enables echo's Debug option.
	Debug bool

	// Dev enables echo's dev mode options.
	Dev bool

	// Listen sets the listen address to serve the echo server on.
	Listen string

	// HTTPS configures an https server
	HTTPS bool

	// ShutdownGracePeriod sets the grace period for in flight requests before shutting down.
	ShutdownGracePeriod time.Duration

	// ReadTimeout sets the maximum duration for reading the entire request including the body.
	ReadTimeout time.Duration

	// WriteTimeout sets the maximum duration before timing out writes of the response.
	WriteTimeout time.Duration

	// IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled.
	IdleTimeout time.Duration

	// ReadHeaderTimeout sets the amount of time allowed to read request headers.
	ReadHeaderTimeout time.Duration

	// TrustedProxies defines the allowed ip / network ranges to trust a proxy from.
	TrustedProxies []string

	// Middleware includes the provided middleware when echo is initialized.
	Middleware []echo.MiddlewareFunc

	TLSConfig TLSConfig
}

// TLSConfig contains config options for the https server
type TLSConfig struct {
	// TLSConfig
	TLSConfig *tls.Config

	CertFile string
	CertKey  string
}

// WithDefaults creates a new config with defaults set if not already defined.
func (c Config) WithDefaults() Config {
	if c.Listen == "" {
		if c.HTTPS {
			// use 443 for secure servers as the default port
			c.Listen = ":443"
		} else {
			c.Listen = ":8080"
		}
	}

	if c.ShutdownGracePeriod <= 0 {
		c.ShutdownGracePeriod = DefaultShutdownGracePeriod
	}

	if c.ReadTimeout <= 0 {
		c.ReadTimeout = DefaultReadTimeout
	}

	if c.WriteTimeout <= 0 {
		c.WriteTimeout = DefaultWriteTimeout
	}

	if c.IdleTimeout <= 0 {
		c.IdleTimeout = DefaultIdleTimeout
	}

	if c.ReadHeaderTimeout <= 0 {
		c.ReadHeaderTimeout = DefaultReadHeaderTimeout
	}

	return c
}

// WithTLSDefaults sets tls default settings
func (c Config) WithTLSDefaults() Config {
	c.WithTLSConfig()
	c.TLSConfig.CertFile = DefaultCertFile
	c.TLSConfig.CertKey = DefaultKeyFile

	return c
}

// WithDebug enables echo's Debug option.
func (c Config) WithDebug(debug bool) Config {
	c.Debug = debug

	return c
}

// WithDev enables echo's dev mode options.
func (c Config) WithDev(dev bool) Config {
	c.Dev = dev

	return c
}

// WithListen sets the listen address to serve the echo server on.
func (c Config) WithListen(listen string) Config {
	c.Listen = listen

	return c
}

// WithHTTPS enables https server options
func (c Config) WithHTTPS(https bool) Config {
	c.HTTPS = https

	return c
}

// WithShutdownGracePeriod sets the grace period for in flight requests before shutting down.
func (c Config) WithShutdownGracePeriod(period time.Duration) Config {
	c.ShutdownGracePeriod = period

	return c
}

// WithDefaultReadTimeout sets the maximum duration for reading the entire request including the body.
func (c Config) WithDefaultReadTimeout(period time.Duration) Config {
	c.ReadTimeout = period

	return c
}

// WithWriteTimeout sets the maximum duration before timing out writes of the response.
func (c Config) WithWriteTimeout(period time.Duration) Config {
	c.WriteTimeout = period

	return c
}

// WithIdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled.
func (c Config) WithIdleTimeout(period time.Duration) Config {
	c.IdleTimeout = period

	return c
}

// WithReadHeaderTimeout sets the amount of time allowed to read request headers.
func (c Config) WithReadHeaderTimeout(period time.Duration) Config {
	c.ReadHeaderTimeout = period

	return c
}

// WithMiddleware includes the provided middleware when echo is initialized.
func (c Config) WithMiddleware(mdw ...echo.MiddlewareFunc) Config {
	c.Middleware = append(c.Middleware, mdw...)

	return c
}

// WithTLSConfig sets the TLS Configuration
func (c Config) WithTLSConfig() Config {
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}

	c.TLSConfig.TLSConfig = cfg

	return c
}

// WithTLSCerts sets the TLS Cert and Key locations
func (c Config) WithTLSCerts(certFile, certKey string) Config {
	c.TLSConfig.CertFile = certFile
	c.TLSConfig.CertKey = certKey

	return c
}
