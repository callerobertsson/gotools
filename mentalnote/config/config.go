// Package config defines the configuration for mentalnote
package config

import (
	"encoding/json"
	"os"
)

// Config type holds the configuration settings
type Config struct {
	APIToken  string `json:"api-token"`
	ChannelID string `json:"channel-id"`
	Username  string `json:"username"`
	IconURL   string `json:"icon-url"`
}

// New creates a new Config
func New(filePath string) (c Config, err error) {

	file, _ := os.Open(filePath)
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	return c, err
}
