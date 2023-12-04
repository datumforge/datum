package config

type (
	// ConfigProvider serves as a common interface to read echo server configuration
	ConfigProvider interface {
		// GetConfig returns the server configuration
		GetConfig() (*Config, error)
	}
)
