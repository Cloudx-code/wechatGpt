package sqlite

import (
	"fmt"
	"testing"

	"wechatGpt/common/utils"
)

func TestName(t *testing.T) {
	Init("test.db")
	//Register(DBProxy, "test", "11111", "", time.Hour)
	var userInfo []User

	if err := DBProxy.Where("user_id=?", "11111").Find(&userInfo).Error; err != nil {
		fmt.Println("fail to find:", err)
		return
	}
	fmt.Println(utils.Encode(userInfo))
}
