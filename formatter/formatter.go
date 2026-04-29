package formatter

import (
	"fmt"
	"code/compareFiles"
	"code/formatter/stylish"
    "code/formatter/plain"
    "code/formatter/json"
)

func GetFormatter (compare []comparefiles.Node, format string) (string, error) {
	switch format {
    case "stylish":
        return stylish.Stylish(compare), nil
    case "plain":
        return plain.Plain(compare), nil
   	case "json":
        data, err := json.Json(compare)
        if err != nil {
            return "", err
        }
        return data, nil
    default:
        	return "", fmt.Errorf("unknown format: %s", format)
    }
};