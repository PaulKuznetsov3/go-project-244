package code

import (
    "fmt"
)

func GenDiff(filepath1, filepath2, format string) (string, error){
	if filepath1 == "" || filepath2 == "" || format == "" {
		return "", fmt.Errorf("Введите путь")
	}

	return  fmt.Sprintf(filepath1, filepath2, format), nil
}
