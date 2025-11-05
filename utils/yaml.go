package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadYAML[T any](path string, out *T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", path, err)
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return fmt.Errorf("failed to parse YAML %s: %w", path, err)
	}
	return nil
}
