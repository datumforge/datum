package appdata

import (
	"encoding/json"
	"fmt"
	"os"

	"dario.cat/mergo"
	"github.com/adrg/xdg"
)

const fileName = "Ampersand/config.json"

// IMPORTANT:
// Do not modify the JSON labels in this struct without ensuring backwards compatibility,
// since those strings are written to the user's config file on their computer.
type Config struct {
	Token Token `json:"token"`
}

// Token represents a JWT token.
type Token struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iat int    `json:"iat"`
	Exp int    `json:"exp"`
}

// Get returns the user's existing config, or an empty config if the file doesn't exist.
func Get() (Config, error) {
	path, err := getExistingFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("can't read config file: %w", err)
	}

	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return Config{}, fmt.Errorf("can't parse config: %w", err)
	}

	return c, nil
}

// Set writes the given config to the user's config file.
// The config can be a partial config, and it is merged with the existing config.
func Set(config Config) error {
	existing, err := Get()
	if err != nil {
		return err
	}

	merged := config
	if err := mergo.Merge(&merged, existing); err != nil {
		return fmt.Errorf("can't merge new config with existing config: %w", err)
	}

	return setEntireConfig(merged)
}

func setEntireConfig(config Config) error {
	path, err := getExistingFilePath()
	if err != nil {
		path, err = getPathForNewFile()
		if err != nil {
			return fmt.Errorf("can't get path for new config file: %w", err)
		}
	}

	js, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("can't marshal config into JSON: %w", err)
	}

	return writeFile(path, js)
}

func getPathForNewFile() (string, error) {
	return xdg.ConfigFile(fileName)
}

func getExistingFilePath() (string, error) {
	return xdg.SearchConfigFile(fileName)
}

const perm = 0o600 // Regular file with read/write permission for owner

func writeFile(path string, data []byte) error {
	return os.WriteFile(
		path,
		data,
		perm,
	)
}
