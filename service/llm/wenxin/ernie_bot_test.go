package wenxin

import (
	"testing"

	"wechatGpt/common/logs"
)

func TestWenXin(t *testing.T) {
	logs.Init(false)
	s := NewWenXinYiYanService()
	//s.getAccessToken()
	s.Query("你是谁")
}
