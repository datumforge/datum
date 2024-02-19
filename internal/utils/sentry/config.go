package sentry

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
)

// Config for the Sentry client
type Config struct {
	Enabled            bool    `split_words:"true" default:"false"`
	DSN                string  `split_words:"true"`
	ServerName         string  `split_words:"true"`
	Environment        string  `split_words:"true"`
	EnableTracing      bool    `split_words:"true" default:"false"`
	TracesSampler      float64 `split_words:"true" default:"1.0"`
	AttachStacktrace   bool    `split_words:"true" default:"true"`
	SampleRate         float64 `split_words:"true" default:"0.2"`
	TracesSampleRate   float64 `split_words:"true" default:"0.2"`
	ProfilesSampleRate float64 `split_words:"true" default:"0.2"`
	Repanic            bool    `ignored:"true"`
	Debug              bool    `default:"false"`
}

// UseSentry true if Sentry is enabled (e.g. a DSN is configured)
func (c Config) UseSentry() bool {
	return c.DSN != ""
}

// UsePerformanceTracking tracking is enabled if Sentry is enabled and track performance is explicitly set
func (c Config) UsePerformanceTracking() bool {
	return c.UseSentry() && c.EnableTracing
}

// ErrInvalidConfiguration is returned when the configuration is invalid
var ErrInvalidConfiguration = errors.New("invalid configuration")

// Validate the configuration
func (c Config) Validate() error {
	if c.UseSentry() && c.Environment == "" {
		return fmt.Errorf("%w: environment must be configured when Sentry is enabled", ErrInvalidConfiguration)
	}

	return nil
}

// ClientOptions returns the sentry.ClientOptions for the configuration
func (c Config) ClientOptions() sentry.ClientOptions {
	opts := sentry.ClientOptions{
		Dsn:              c.DSN,
		ServerName:       c.ServerName,
		Environment:      c.Environment,
		TracesSampleRate: c.TracesSampler,
		EnableTracing:    c.EnableTracing,
		AttachStacktrace: c.AttachStacktrace,
		SampleRate:       c.SampleRate,
		Debug:            c.Debug,
	}

	return opts
}
