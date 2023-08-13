package utils

import (
	"encoding/json"
)

func Encode(srcObj interface{}) string {
	if resBytes, err := json.Marshal(srcObj); err != nil {
		return ""
	} else {
		return string(resBytes)
	}
}
