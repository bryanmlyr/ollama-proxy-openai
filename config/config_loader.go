package config

import (
	"gopkg.in/yaml.v3"
	"ollamaproxy/model"
	"os"
	"strings"
)

type ConfigLoader struct{}

func NewConfigLoader() *ConfigLoader { return &ConfigLoader{} }

func (c *ConfigLoader) LoadConfig() ([]model.OllamaProxyConfig, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var wrapper []model.OllamaProxyConfig
	if err := yaml.NewDecoder(f).Decode(&wrapper); err != nil {
		return nil, err
	}

	for i := range wrapper {
		resolveEnvVarsOllamaProxyConfig(&wrapper[i])
	}

	return wrapper, nil
}

// handle your config struct specifically and safely, not reflectively
func resolveEnvVarsOllamaProxyConfig(cfg *model.OllamaProxyConfig) {
	// Identifier
	if strings.HasPrefix(cfg.Identifier, "$") {
		cfg.Identifier = os.Getenv(strings.TrimPrefix(cfg.Identifier, "$"))
	}
	// Endpoint
	if strings.HasPrefix(cfg.Endpoint, "$") {
		cfg.Endpoint = os.Getenv(strings.TrimPrefix(cfg.Endpoint, "$"))
	}
	// Implementation (if needed, add your logic here, depending on its type)
	// Models
	for i, v := range cfg.Models {
		if strings.HasPrefix(v, "$") {
			cfg.Models[i] = os.Getenv(strings.TrimPrefix(v, "$"))
		}
	}
	// Key
	if cfg.Key != nil && strings.HasPrefix(*cfg.Key, "$") {
		envVal := os.Getenv(strings.TrimPrefix(*cfg.Key, "$"))
		cfg.Key = &envVal
	}
}
