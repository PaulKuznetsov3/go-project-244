package comparefiles

import (
	"sort"
	"fmt"
)

type Node struct {
	Key      string
	OldValue   any
	NewValue   any
	Type     string
	Children []Node
}


func CompareFiles(dataFile1, dataFile2 map[string]any )[]Node {
	sortKeys := getSortedKeys(dataFile1, dataFile2)
	result := make([]Node, 0, len(sortKeys))

	for _, key := range sortKeys {
		val1, ok1 := dataFile1[key]
		val2, ok2 := dataFile2[key]

		if !ok1 {
			result = append(result, Node{
				Key:      key,
				NewValue: val2,
				Type:     "added",
			})
			continue
		}

		if !ok2 {
			result = append(result, Node{
				Key:      key,
				OldValue: val1,
				Type:     "deleted",
			})
			continue
		}

		if isMap(val1) && isMap(val2) {
			result = append(result, Node{
				Key:      key,
				Children: CompareFiles(val1.(map[string]any), val2.(map[string]any)),
				Type:     "nested",
			})
			continue
		}

		if !isEqual(val1, val2) {
			result = append(result, Node{
				Key:      key,
				OldValue: val1,
				NewValue: val2,
				Type:     "changed",
			})
			continue
		}

		result = append(result, Node{
			Key:      key,
			OldValue: val1,
			Type:     "unchanged",
		})
	}

	return result
}

func getSortedKeys(file1, file2 map[string]any) []string {
	keysMap := make(map[string]bool)
	
	for key := range file1 {
		keysMap[key] = true
	}
	
	for key := range file2 {
		keysMap[key] = true
	}
	
	keys := make([]string, 0, len(keysMap))
	for key := range keysMap {
		keys = append(keys, key)
	}
	
	sort.Strings(keys)
	
	return keys
}

func isEqual(a, b any) bool {
	return fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b)
}

// isMap Проверяет является ли значение map
func isMap(value any) bool {
    _, ok := value.(map[string]any)
    return ok
}