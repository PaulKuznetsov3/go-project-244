package formatter

import (
	"fmt"
	"code/compareFiles"
	"code/formatter/stylish"
)

func GetFormatter (compare []comparefiles.Node, format string) (string, error) {
	switch format {
    case "stylish":
        return stylish.Stylish(compare), nil
    // case "plain":
    //     return fmt.Sprintf("%.1f%s", currentSize/KB, "KB")
   	// case "json":
    //     return fmt.Sprintf("%.1f%s", currentSize/PB, "PB")
    default:
        	return "", fmt.Errorf("unknown format: %s", format)
    }
};