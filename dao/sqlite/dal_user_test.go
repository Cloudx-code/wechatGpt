package sqlite

import (
	"fmt"
	"testing"

	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
)

func TestName(t *testing.T) {
	Init("test.db")
	logs.Init(false)
	//Register(DBProxy, "test", "11111", "", time.Hour)
	s, err := NewDalUser().GetUserById("2222")
	if err != nil {
		fmt.Println("fail to err:", err)
		return
	}
	fmt.Println(utils.Encode(s))
}
