package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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
