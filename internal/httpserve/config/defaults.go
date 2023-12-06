package config

import (
	"crypto/tls"
	"time"
)

var (
	// DefaultListenAddr sets the default listen for the server
	DefaultListenAddr = ":17608"
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
	// DefaultConfigRefresh sets the default interval to refresh the config.
	DefaultConfigRefresh = 10 * time.Minute
	// DefaultCertFile is the default cert file location
	DefaultCertFile = "server.crt"
	// DefaultKeyFile is the default key file location
	DefaultKeyFile = "server.key"
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
