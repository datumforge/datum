package sentry

import (
	"errors"

	"github.com/getsentry/sentry-go"
)

// Config configuration for use in application-configuration
type Config struct {
	DSN              string  `split_words:"true"`
	ServerName       string  `split_words:"true"`
	Environment      string  `split_words:"true"`
	Release          string  `split_words:"true"`
	TrackPerformance bool    `split_words:"true" default:"false"`
	SampleRate       float64 `split_words:"true" default:"0.2"`
	UseStatusSampler bool    `split_words:"true" default:"true"`
	ReportErrors     bool    `split_words:"true" default:"true"`
	Repanic          bool    `ignored:"true"`
	Debug            bool    `default:"false"`
}

// UseSentry returns true if Sentry is enabled (e.g. a DSN is configured)
func (c Config) UseSentry() bool {
	return c.DSN != ""
}

// UsePerformanceTracking is enabled if Sentry is enabled and track performance is explicitly set
func (c Config) UsePerformanceTracking() bool {
	return c.UseSentry() && c.TrackPerformance
}

func (c Config) Validate() error {
	if c.UseSentry() && c.Environment == "" {
		return errors.New("invalid configuration: environment must be configured when Sentry is enabled")
	}
	return nil
}

func (c Config) ClientOptions() sentry.ClientOptions {
	opts := sentry.ClientOptions{
		Dsn:              c.DSN,
		Environment:      c.Environment,
		AttachStacktrace: true,
		Debug:            c.Debug,
		ServerName:       c.ServerName,
		EnableTracing:    c.TrackPerformance,
	}

	if c.TrackPerformance {
		if c.UseStatusSampler {
			opts.TracesSampler = NewStatusSampler(c.SampleRate)
		} else {
			opts.TracesSampleRate = c.SampleRate
		}
	}

	return opts
}
