// Package config provides environment management for necessary sensitive information
package config

import "os"

type Config struct {
	BaseURL string // AnythingLLM Base URL
	APIKey  string // AnythingLLM API Key
}

func Load() *Config {
	baseURL := os.Getenv("AnythingLMM_URL")
	if baseURL == "" {
		baseURL = "http://localhost:3001"
	}

	return &Config{
		BaseURL: baseURL,
		APIKey:  os.Getenv("ANYTHINGLLM_API_KEY"),
	}
}
