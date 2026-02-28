package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type HTTPConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	HTTP HTTPConfig `yaml:"http"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("Read config.yaml: %w", err)
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("Unmarshal yaml: %w", err)
	}

	return cfg, nil
}
