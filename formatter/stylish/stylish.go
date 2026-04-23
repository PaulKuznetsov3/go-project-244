package stylish

import (
	"code/compareFiles"
	"strings"
	"fmt"
)

var defaultIndent int = 4

// MakeIndent создает отступ для текущего уровня
func MakeIndent(depth, spaceCount int) string {
	return strings.Repeat(" ", depth*spaceCount-2)
}

// MakeBackIndent создает закрывающий отступ
func MakeBackIndent(depth, spaceCount int) string {
	return strings.Repeat(" ", depth*spaceCount)
}

// Stringify преобразует значение в строку
func Stringify(value any, depth int) string {
	if !isMap(value) {
		return fmt.Sprintf("%v", value)
	}

	valueMap, ok := value.(map[string]any)
	if !ok {
		return fmt.Sprintf("%v", value)
	}

	// Сортируем ключи для стабильного вывода
	keys := getSortedKeysFromMap(valueMap)
	
	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		val := valueMap[key]
		indent := MakeIndent(depth, defaultIndent)
		lines = append(lines, fmt.Sprintf("%s  %s: %s", indent, key, Stringify(val, depth+1)))
	}
	
	result := strings.Join(lines, "\n")
	backIndent := MakeBackIndent(depth-1, defaultIndent)
	
	return fmt.Sprintf("{\n%s\n%s}", result, backIndent)
}

// isMap проверяет, является ли значение map
func isMap(value any) bool {
	_, ok := value.(map[string]any)
	return ok
}

// getSortedKeysFromMap возвращает отсортированные ключи из map
func getSortedKeysFromMap(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}
	
	return keys
}

// Stylish форматирует дерево различий в стиле stylish
func Stylish(tree []comparefiles.Node) string {
	iter := func(nodes []comparefiles.Node, depth int) string {
		lines := make([]string, 0, len(nodes))
		
		for _, node := range nodes {
			switch node.Type {
			case "added":
				indent := MakeIndent(depth, defaultIndent)
				lines = append(lines, fmt.Sprintf("%s+ %s: %s", 
					indent, node.Key, Stringify(node.NewValue, depth+1)))
				
			case "deleted":
				indent := MakeIndent(depth, defaultIndent)
				lines = append(lines, fmt.Sprintf("%s- %s: %s", 
					indent, node.Key, Stringify(node.OldValue, depth+1)))
				
			case "nested":
				indent := MakeIndent(depth, defaultIndent)
				nestedContent := Stylish(node.Children)
				nestedContent = nestedContent[1 : len(nestedContent)-1]
				lines = append(lines, fmt.Sprintf("%s  %s: {\n%s\n%s", 
					indent, node.Key, nestedContent, MakeBackIndent(depth, 4)))
				
			case "changed":
				indent := MakeIndent(depth, defaultIndent)
				lines = append(lines, fmt.Sprintf("%s- %s: %s", 
					indent, node.Key, Stringify(node.OldValue, depth+1)))
				lines = append(lines, fmt.Sprintf("%s+ %s: %s", 
					indent, node.Key, Stringify(node.NewValue, depth+1)))
				
			case "unchanged":
				indent := MakeIndent(depth, defaultIndent)
				lines = append(lines, fmt.Sprintf("%s  %s: %s", 
					indent, node.Key, Stringify(node.OldValue, depth+1)))
				
			default:
				panic(fmt.Sprintf("unknown format: '%s'!", node.Type))
			}
		}
		
		return strings.Join(lines, "\n")
	}
	
	return fmt.Sprintf("{\n%s\n}", iter(tree, 1))
}