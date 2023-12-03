package config

type (
	// ConfigProvider serves as a common interface to read echo server configuration
	ConfigProvider interface {
		GetConfig() (*Config, error)
	}
)
