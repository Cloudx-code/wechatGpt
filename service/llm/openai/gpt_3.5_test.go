package openai

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
)

func TestGpt35(t *testing.T) {
	logs.Init(false)
	fmt.Println(Completions("你好"))
}
