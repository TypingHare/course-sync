package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Store struct {
	v *viper.Viper
}

func NewStore() (*Store, error) {
	v := viper.New()

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	v.SetConfigFile(configFilePath)
	v.SetConfigType("json")

	// Create config file with defaults if it doesn't exist
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(configFilePath), 0o755); err != nil {
			return nil, fmt.Errorf("create config dir: %w", err)
		}

		defaultConfig := GetDefault()

		bytes, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			return nil, fmt.Errorf("marshal default config: %w", err)
		}

		if err := os.WriteFile(configFilePath, bytes, 0o644); err != nil {
			return nil, fmt.Errorf("write default config: %w", err)
		}
	}

	// Now read the config (guaranteed to exist)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	return &Store{v: v}, nil
}

// Load reads the config from file and environment variables.
func (store *Store) Load() (Config, error) {
	var config Config
	if err := store.v.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("unmarshal config: %w", err)
	}

	return config, nil
}

// Set sets a config value and saves the config to file.
func (store *Store) Set(key string, value any) error {
	store.v.Set(key, value)
	return store.Save()
}

// Save writes the current config to file.
func (store *Store) Save() error {
	if err := store.v.WriteConfig(); err != nil {
		if err := store.v.SafeWriteConfig(); err != nil {
			if err2 := store.v.WriteConfig(); err2 != nil {
				return fmt.Errorf("write config: %w", err2)
			}
		}
	}

	return nil
}
