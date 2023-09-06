package utils

import (
	"fmt"
	"os"

	"github.com/eatmoreapple/openwechat"
)

func Reply(msg *openwechat.Message, text string) {
	msg.ReplyText(text)
	fmt.Println("回复成功：")
	fmt.Println(text)

}

func ReplyImage(msg *openwechat.Message) {
	img, _ := os.Open("your file path")
	defer img.Close()
	msg.ReplyImage(img)
}
