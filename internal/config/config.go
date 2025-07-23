package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents the application configuration structure
type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const (
	configFileName = ".gatorconfig.json"
	fileMode       = 0644
)

// Read loads and returns the configuration from the config file
func Read() (Config, error) {
	config := Config{}
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get config file path: %w", err)
	}

	configBytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(configBytes, &config); err != nil {
		return Config{}, fmt.Errorf("failed to parse config file: %w", err)
	}

	return config, nil
}

// getConfigFilePath returns the full path to the configuration file
func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}

// SetUser updates the current user in the configuration and saves it
func (c *Config) SetUser(userName string) error {
	if userName == "" {
		return fmt.Errorf("username cannot be empty")
	}

	c.CurrentUserName = userName
	if err := write(*c); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}

// write persists the current configuration to disk
func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}

	configBytes, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to serialize configuration: %w", err)
	}

	if err := os.WriteFile(filePath, configBytes, fileMode); err != nil {
		return fmt.Errorf("failed to write configuration file: %w", err)
	}

	return nil
}
