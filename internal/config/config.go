package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// StudentInfo represents a stutent's information.
type StudentInfo struct {
	Name           string `mapstructure:"name"`
	Email          string `mapstructure:"email"`
	RepositoryPath string `mapstructure:"repositoryPath"`
}

// Config represents the application configuration structure.
type Config struct {
	// Student related configuration.
	Student struct{} `mapstructure:"student"`

	// Master (teachers, professors, TAs, graders, etc.) related configuration.
	Master struct {
		Roster []StudentInfo `mapstructure:"roster"`
	} `mapstructure:"master"`
}

// configCache holds the cached configuration.
var configCache *Config

// GetDefault returns the default configuration.
func GetDefault() *Config {
	var config Config
	return &config
}

// Load reads the config from file.
func Load(configPath string) (*Config, error) {
	if strings.TrimSpace(configPath) == "" {
		return nil, errors.New("config path is empty")
	}

	content, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return GetDefault(), nil
		}
		return nil, fmt.Errorf("read config: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()

	config := GetDefault()
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("decode config json: %w", err)
	}

	// Ensure there isn't trailing junk after the JSON object/array
	if decoder.More() {
		return nil, errors.New("decode config json: trailing data")
	}

	if err := Validate(config); err != nil {
		return nil, err
	}

	return nil, nil
}

// Validate checks if the config struct has valid values.
func Validate(config *Config) error {
	// TODO: Implement validation logic for the config struct
	return nil
}

// Save writes cfg as JSON to path, creating parent directories if needed. Uses an atomic write
// (write temp file then rename).
func Save(configPath string, config *Config) error {
	if strings.TrimSpace(configPath) == "" {
		return errors.New("config path is empty")
	}
	if config == nil {
		return errors.New("config is nil")
	}

	if err := Validate(config); err != nil {
		return err
	}

	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := json.MarshalIndent(&config, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	data = append(data, '\n')

	tmp, err := os.CreateTemp(dir, ".config-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}

	// Make sure we clean up temp file on failure
	tmpName := tmp.Name()
	defer func() {
		_ = os.Remove(tmpName)
	}()

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("write temp config: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("sync temp config: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temp config: %w", err)
	}

	// os.Rename is atomic on the same filesystem.
	if err := os.Rename(tmpName, configPath); err != nil {
		return fmt.Errorf("rename temp config: %w", err)
	}

	return nil
}

// Get returns the cached configuration, loading it from file if not already cached.
func Get() (*Config, error) {
	if configCache != nil {
		return configCache, nil
	}

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("get config file path: %w", err)
	}

	config, err := Load(configFilePath)
	if err != nil {
		return nil, err
	}

	configCache = config
	return configCache, nil
}
