package weTab

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
	"wechatGpt/dao/local_cache"
)

func TestName(t *testing.T) {
	local_cache.InitCache()
	logs.Init(false)
	fmt.Println(NewWeTabService("test").Query("你好"))
}
