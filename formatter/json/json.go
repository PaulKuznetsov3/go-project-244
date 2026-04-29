package json

import (
	"code/compareFiles"
    "encoding/json"
	"fmt"
)

// Json форматирует дерево различий в стиле json
func Json(nodes []comparefiles.Node) (string, error) {
    data, err := json.MarshalIndent(nodes, "", "  ")
    if err != nil {
        return "", fmt.Errorf("failed to marshal JSON: %w", err)
    }
    return string(data), nil
}