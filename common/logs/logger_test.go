package logs

import (
	"testing"
)

func TestLog(t *testing.T) {
	Init(false)
	Info("hi:%v", 1)
	Logger.Printf("你好啊:%v", 1)
}
