package config

import (
	"fmt"
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Webauthn WebauthnSettings
}

type WebauthnSettings struct {
	RelyingParty     RelyingParty `yaml:"relying_party" json:"relying_party,omitempty" koanf:"relying_party" split_words:"true"`
	Timeout          int          `yaml:"timeout" json:"timeout,omitempty" koanf:"timeout" jsonschema:"default=60000"`
	UserVerification string       `yaml:"user_verification" json:"user_verification,omitempty" koanf:"user_verification" split_words:"true" jsonschema:"default=preferred,enum=required,enum=preferred,enum=discouraged"`
}

type RelyingParty struct {
	ID          string   `yaml:"id" json:"id,omitempty" koanf:"id" jsonschema:"default=localhost"`
	DisplayName string   `yaml:"display_name" json:"display_name,omitempty" koanf:"display_name" split_words:"true" jsonschema:"default=Datum Authentication Service"`
	Icon        string   `yaml:"icon" json:"icon,omitempty" koanf:"icon"`
	Origins     []string `yaml:"origins" json:"origins,omitempty" koanf:"origins" jsonschema:"minItems=1,default=http://localhost:17608"`
}

var (
	DefaultConfigFilePath = "./config/config.yaml"
)

// Load is responsible for loading the configuration from a YAML file and
// environment variables. It takes a pointer to a string `cfgFile` as a parameter, which represents the
// path to the configuration file. If the `cfgFile` is empty or nil, it sets the default configuration
// file path.
func Load(cfgFile *string) (*Config, error) {
	k := koanf.New(".")

	var err error

	if cfgFile == nil || *cfgFile == "" {
		*cfgFile = DefaultConfigFilePath
	}

	if err = k.Load(file.Provider(*cfgFile), yaml.Parser()); err != nil {
		if *cfgFile != DefaultConfigFilePath {
			return nil, fmt.Errorf("failed to load config from: %s: %w", *cfgFile, err)
		}

		log.Println("failed to load config, skipping...")
	} else {
		log.Println("Using config file:", *cfgFile)
	}

	c := DefaultConfig()
	err = k.Unmarshal("", c)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := envconfig.Process("", c); err != nil {
		return nil, fmt.Errorf("failed to load config from env vars: %w", err)
	}

	return c, nil
}

// DefaultConfig returns a pointer to a `Config` struct with default values set
func DefaultConfig() *Config {
	return &Config{
		Webauthn: WebauthnSettings{
			RelyingParty: RelyingParty{
				ID:          "localhost", // nolint: gomnd
				DisplayName: "Datum Authentication Service",
				Origins:     []string{"http://localhost:16708"},
			},
			UserVerification: "preferred",
			Timeout:          60000, // nolint: gomnd
		},
	}
}
