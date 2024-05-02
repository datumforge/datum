package kafkaconfig

// Config is the configuration for the Kafka event source
type Config struct {
	Enabled   bool     `json:"enabled" koanf:"enabled" default:"false"`
	AppName   string   `json:"appName" koanf:"appName" default:"datum"`
	Address   string   `json:"address" koanf:"address" default:"localhost:10000"`
	Addresses []string `json:"addresses" koanf:"addresses" default:"[localhost:10000]"`
	Debug     bool     `json:"debug" koanf:"debug" default:"false"`
}
