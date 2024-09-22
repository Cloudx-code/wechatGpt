package wenxin

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
)

func TestWenXin(t *testing.T) {
	logs.Init(false)
	local_cache.InitCache()
	config.InitStaticConfig("../../../config/static_config.json")
	s := NewWenXinYiYanService("test")
	fmt.Println(utils.Encode(s))
	s.getAccessToken()
	ans, err := s.Query("你是谁")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(utils.Encode(ans))
}
