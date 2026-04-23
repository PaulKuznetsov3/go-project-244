package code

import (
    "fmt"
	"code/parser"
	"code/formatter"
	"code/compareFiles"
)

/**
* Функция сравнивания файлов. 
* 
* filepath1, filepath1 Пути до файлов.
* fopmat Формат вывода результата сравнения файлов.
*/
func GenDiff(filepath1, filepath2, format string) (string, error){
	if filepath1 == "" || filepath2 == "" {
	    return "", fmt.Errorf("file paths cannot be empty: %q, %q", filepath1, filepath2)
	}

	data1, err1 := code.Parser(filepath1)
	data2, err2 := code.Parser(filepath2)

	if err1 != nil || err2 != nil {
    	if err1 != nil {
       		return "", err1
    	}
    	return "", err2
	}

	diff :=	comparefiles.CompareFiles(data1, data2)

	result, err := formatter.GetFormatter(diff, format)

	if err != nil {
		return "", err
	}

	return  result, nil
}
