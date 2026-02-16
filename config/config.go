package config

import "os"

// Config holds application configuration
type Config struct {
	NatsURL string
}

// Load loads configuration from environment variables
func Load() *Config {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	return &Config{
		NatsURL: natsURL,
	}
}
