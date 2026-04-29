package code

import (
    "fmt"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
	"gopkg.in/yaml.v3"
)

// Parser функция парсинга файла. 
func Parser(path string) (map[string]any, error) {
	var result map[string]any

	if !filepath.IsAbs(path) {
		absFilePath, err := filepath.Abs(path)
		if err != nil {
			return nil, fmt.Errorf("failed to convert path to absolute: %w", err)
		}
		path = absFilePath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", path, err)
	}

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("failed to parse JSON file %s: %w", path, err)
		}

	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("failed to parse YAML file %s: %w", path, err)
		}

	default:
		return nil, fmt.Errorf("unsupported file format %s. Supported formats: .json, .yaml, .yml", ext)
	}

	return result, nil
}