package config

import (
	"os"
	"time"
)

type Config struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	RequestTimeout time.Duration
}

func Load() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		ReadTimeout:    getDurationEnv("READ_TIMEOUT", 10*time.Second),
		WriteTimeout:   getDurationEnv("WRITE_TIMEOUT", 10*time.Second),
		RequestTimeout: getDurationEnv("REQUEST_TIMEOUT", 30*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}
