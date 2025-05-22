package services

import (
	"errors"
	"fmt"

	"github.com/hfleury/re-test/config"
)

type ConfigService struct {
	config   *config.Config
	filePath string
}

func NewConfigService(cfg *config.Config, filePath string) *ConfigService {
	return &ConfigService{config: cfg, filePath: filePath}
}

func (s *ConfigService) UpdatePackSizes(sizes []int) error {
	if len(sizes) == 0 {
		return errors.New("pack sizes cannot be empty")
	}

	s.config.PackSize = sizes

	// Persist to the configuration.yaml file
	err := s.config.SaveToFile(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func (s *ConfigService) GetPackSizes() []int {
	return s.config.PackSize
}
