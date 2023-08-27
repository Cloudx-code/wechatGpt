package wenxin

import (
	"testing"
)

func TestWenXin(t *testing.T) {
	s := NewWenXinYiYanService()
	s.Query("你好")
}
