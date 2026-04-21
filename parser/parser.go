package code

import (
    "fmt"
	"os"
	"path/filepath"
)

/**
* Функция парсинга файла. 
* 
* path Путь до файла.
*/
func Parser(path string)(string, error) {
	if !filepath.IsAbs(path) {
        absFilePath, err := filepath.Abs(path)
		if err != nil {
			return "", fmt.Errorf("Ошибка преобразования пути")
		}
        path = absFilePath
    }

	data, err := os.ReadFile(path)
	if err != nil {
	   	return "", fmt.Errorf("Ошибка чтения файла")
	}

	return string(data), nil
}