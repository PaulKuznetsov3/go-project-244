package plain

import (
	"fmt"
	"strings"
	"code/compareFiles"
)

// isMap проверяет, является ли значение map
func isMap(value any) bool {
	_, ok := value.(map[string]any)
	return ok
}

// Stringify преобразует значение в строку
func Stringify(value any) string {
	if value == nil {
		return "null"
	}
	if isMap(value) {
    	return "[complex value]";
  	}
  	if _, ok := value.(string); ok {
		return fmt.Sprintf("'%s'", value)
	}
  	return fmt.Sprintf("%v", value);
}

// getPath формирует путь до значения
func getPath(valuePath, key string)string {
	if valuePath == "" {
		return key
	}
	return fmt.Sprintf("%s.%s", valuePath, key)
}

// Plain форматирует дерево различий в стиле Plain
func Plain(tree []comparefiles.Node) string {
	var iter func(nodes []comparefiles.Node, path string) string
	iter = func (nodes []comparefiles.Node, basePath string) string {
		var builder strings.Builder
		for _, node := range nodes {
			path := getPath(basePath, node.Key)
			switch node.Type {
			case "nested":
				childStr := iter(node.Children, path)
				builder.WriteString(childStr)
			case "deleted":
				builder.WriteString(fmt.Sprintf("Property '%s' was removed\n", path))
			case "added":
				builder.WriteString(fmt.Sprintf("Property '%s' was added with value: %s\n",path, Stringify(node.NewValue)))
			case "changed":
				builder.WriteString(fmt.Sprintf("Property '%s' was updated. From %s to %s\n", path, Stringify(node.OldValue), Stringify(node.NewValue)))
			}
		}
		return builder.String()
	}
	
	return strings.TrimSpace(iter(tree, ""))
}
