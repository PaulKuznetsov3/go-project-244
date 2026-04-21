package code

import (
    "fmt"
	"code/parser"
)

/**
* Функция сравнивания файлов. 
* 
* filepath1, filepath1 Пути до файлов.
* fopmat Формат вывода результата сравнения файлов.
*/
func GenDiff(filepath1, filepath2, format string) (string, error){
	if filepath1 == "" || filepath2 == "" || format == "" {
		return "", fmt.Errorf("Введите путь")
	}

	data1, err := code.Parser(filepath1)
	data2, err := code.Parser(filepath2)

	if err != nil {
		return  "", err
	}

	return  fmt.Sprint(data1, data2, format), nil
}
