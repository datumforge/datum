package config

import (
	"time"
)

type (
	// Config contains the configuration for the datum server
	Config struct {
		// How often to reload the config
		RefreshInterval time.Duration `yaml:"refreshInterval"`
	}
	//TODO: blow this out to hold all our other configs
)

// Ensure that *Config implements ConfigProvider interface.
var _ ConfigProvider = &Config{}

// GetConfig implements ConfigProvider.
func (c *Config) GetConfig() (*Config, error) {
	return c, nil
}
