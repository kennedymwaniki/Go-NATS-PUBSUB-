package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	NATSUrl string
	Subject string
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	subject := os.Getenv("NATS_SUBJECT")
	if subject == "" {
		subject = "notifications"
	}

	return &Config{
		NATSUrl: natsURL,
		Subject: subject,
	}
}
