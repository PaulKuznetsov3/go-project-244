package code

import (
    "fmt"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
)

/**
* Функция парсинга файла. 
* 
* path Путь до файла.
*/
func Parser(path string)(map[string]any, error) {
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
	   	return nil, fmt.Errorf("file reading error %s: %w", path, err)
	}

	ext := strings.ToLower(filepath.Ext(path))
	
	
	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("ошибка парсинга JSON файла %s: %w", path, err)
		}
		
	// case ".yaml", ".yml":
	// 	if err := yaml.Unmarshal(data, &result); err != nil {
	// 		return nil, fmt.Errorf("ошибка парсинга YAML файла %s: %w", path, err)
	// 	}
		
	default:
		return nil, fmt.Errorf("неподдерживаемый формат файла %s. Поддерживаемые форматы: .json, .yaml, .yml", ext)
	}
	
	return result, nil
}