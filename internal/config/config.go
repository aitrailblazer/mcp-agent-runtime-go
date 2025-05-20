// Package config provides configuration management for the MCP Agent Runtime
package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port string
}

// Load loads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		Port: getenv("MCP_PORT", "8080"),
	}
}

// getenv retrieves an environment variable or returns a fallback value
func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
