package utils

import (
	"strings"
)

func InStrArray(str string, strArray []string) bool {
	for _, val := range strArray {
		if str == val {
			return true
		}
	}
	//
	return false
}

// ContainStrArray 判断strArray中的元素是否包含str
func ContainStrArray(str string, strArray []string) (string, bool) {
	for _, val := range strArray {
		if strings.Contains(str, val) {
			return val, true
		}
	}
	return "", false
}
