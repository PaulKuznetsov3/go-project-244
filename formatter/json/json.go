package json

import (
	"code/compareFiles"
    "encoding/json"
	"fmt"
)

// Json форматирует дерево различий в стиле json
func Json(nodes []comparefiles.Node) (string, error) {
	// Оборачиваем массив в структуру с ключом "diff"
	wrapper := map[string]any{
		"diff": nodes,
	}
	
	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(data), nil
}