package wenxin

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
)

func TestWenXin(t *testing.T) {
	logs.Init(false)
	config.InitStaticConfig("../../../config/static_config.json")
	s := NewWenXinYiYanService()
	fmt.Println(utils.Encode(s))
	//s.getAccessToken()
	//s.Query("你是谁")
}
