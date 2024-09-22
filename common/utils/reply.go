package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/eatmoreapple/openwechat"
)

func Reply(msg *openwechat.Message, text string) {
	msg.ReplyText(text)
	fmt.Println("回复成功：")
	fmt.Println(text)
}

func ReplyImage(msg *openwechat.Message, path string) {
	img, err := os.Open(path)
	if err != nil {
		Reply(msg, "图片路径有误")
	}
	defer img.Close()
	msg.ReplyImage(img)
}

func ClearContent(content string, needMoveStr string) string {
	content = strings.Replace(content, needMoveStr, "", -1)
	content = strings.Replace(content, "+", "", -1)
	content = strings.TrimSpace(content)
	return content
}
