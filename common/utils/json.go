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

func Decode(jsonStr string, destObj interface{}) error {
	return json.Unmarshal([]byte(jsonStr), &destObj)
}
