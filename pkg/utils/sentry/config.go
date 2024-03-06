package sentry

import (
	"errors"
	"fmt"

	"github.com/getsentry/sentry-go"
)

// Config settings for the Sentry client
type Config struct {
	// Enabled indicates whether the Sentry client is enabled
	Enabled bool `json:"enabled" koanf:"enabled" default:"false"`
	// DSN is the Data Source Name for the Sentry client
	DSN string `json:"dsn" koanf:"dsn"`
	// Environment is the environment in which the Sentry client is running
	Environment string `json:"environment" koanf:"environment" default:"development"`
	// EnableTracing indicates whether tracing is enabled for the Sentry client
	EnableTracing bool `json:"enableTracing" koanf:"enableTracing" default:"false"`
	// TracesSampler is the sampling rate for tracing in the Sentry client
	TracesSampler float64 `json:"traceSampler" koanf:"traceSampler" default:"1.0"`
	// AttachStacktrace indicates whether to attach stack traces to events in the Sentry client
	AttachStacktrace bool `json:"attachStacktrace" koanf:"attachStacktrace" default:"true"`
	// SampleRate is the sampling rate for events in the Sentry client
	SampleRate float64 `json:"sampleRate" koanf:"sampleRate" default:"0.2"`
	// TracesSampleRate is the sampling rate for tracing events in the Sentry client
	TracesSampleRate float64 `json:"traceSampleRate" koanf:"traceSampleRate" default:"0.2"`
	// ProfilesSampleRate is the sampling rate for profiling events in the Sentry client
	ProfilesSampleRate float64 `json:"profileSampleRate" koanf:"profileSampleRate" default:"0.2"`
	// Repanic indicates whether to repanic after capturing an event in the Sentry client
	Repanic bool `json:"repanic" koanf:"repanic" ignored:"true"`
	// Debug indicates whether debug mode is enabled for the Sentry client
	Debug bool `json:"debug" koanf:"debug" default:"false"`
	// ServerName is the name of the server running the Sentry client
	ServerName string `json:"serverName" koanf:"serverName"`
}

// UseSentry true if Sentry is enabled (e.g. a DSN is configured)
func (c Config) UseSentry() bool {
	return c.DSN != "" && c.Enabled
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
