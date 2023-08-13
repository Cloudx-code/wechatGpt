package utils

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
)

func Reply(msg *openwechat.Message, text string) {
	msg.ReplyText(text)
	fmt.Println("回复成功：")
	fmt.Println(text)

}
