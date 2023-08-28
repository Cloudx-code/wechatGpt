package wenxin

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
)

func TestWenXin(t *testing.T) {
	logs.Init(false)
	s := NewWenXinYiYanService()
	str, _ := s.getAccessToken()
	fmt.Println(str)
}
