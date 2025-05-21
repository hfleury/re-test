package config

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	PackSize []int `yaml:"pack_sizes"`
}

func NewConfig(path string) (*Config, error) {
	return LoadFromPath(path)
}

func LoadFromPath(path string) (*Config, error) {
	var config Config
	
	data, err := os.ReadFile(path)	
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	if len(config.PackSize) == 0 {
		return nil, fmt.Errorf("pack_size is empty or not configured")
	}
	
	return &config, nil
}