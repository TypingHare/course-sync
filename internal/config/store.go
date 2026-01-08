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

func (store *Store) seedDefaults(def Config) error {
	// SetDefault wants key/value pairs, so either:
	// A) call SetDefault per field (most explicit), or
	// B) convert struct -> map and SetDefault recursively (more generic)
	//
	// I’ll show option B using JSON round-trip to map:
	m, err := configToMap(def)
	if err != nil {
		return err
	}
	setDefaultsFromMap(store.v, "", m)
	return nil
}

func configToMap(cfg Config) (map[string]any, error) {
	b, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func setDefaultsFromMap(v *viper.Viper, prefix string, m map[string]any) {
	for k, val := range m {
		fullKey := k
		if prefix != "" {
			fullKey = prefix + "." + k
		}

		// Recurse into nested objects
		if child, ok := val.(map[string]any); ok {
			setDefaultsFromMap(v, fullKey, child)
			continue
		}

		v.SetDefault(fullKey, val)
	}
}

func NewStore() (*Store, error) {
	v := viper.New()

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	v.SetConfigFile(configFilePath)
	v.SetConfigType("json")

	store := &Store{v: v}

	// Seed defaults into memory always
	if err := store.seedDefaults(GetDefault()); err != nil {
		return nil, fmt.Errorf("seed defaults: %w", err)
	}

	// Read file if present; if missing, keep defaults
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	return store, nil
}

// Load reads the config from file.
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

func (store *Store) InitFile(force bool) error {
	path := store.v.ConfigFileUsed()
	if path == "" {
		return fmt.Errorf("config file path is not set")
	}

	// If file exists and not force -> stop
	if _, err := os.Stat(path); err == nil && !force {
		return fmt.Errorf("config file already exists (use --force to overwrite)")
	} else if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("stat config: %w", err)
	}

	// Ensure parent dir exists
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	// Get the current in-memory config (file merged over defaults)
	cfg, err := store.Load()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	// Write (overwrite if force, otherwise create new)
	flags := os.O_WRONLY | os.O_CREATE
	if force {
		flags |= os.O_TRUNC
	} else {
		flags |= os.O_EXCL
	}

	f, err := os.OpenFile(path, flags, 0o644)
	if err != nil {
		return fmt.Errorf("write config file: %w", err)
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("write config bytes: %w", err)
	}

	return nil
}
