package service

import (
	"strings"
)

func GenReply(content string) string {
	var ans string
	if strings.Contains(content, "行头") {
		return "'我准备了一身行头：出自俞越去面试时面试官问他为这场面试准备了什么，俞越答：准备了一身行头'"
	}
	if strings.Contains(content, "俞越") {
		return "俞越，xxxxx待补充，没想好"
	}
	if strings.Contains(content, "救人") {
		return "俞越曾经在游泳时救了一个人，受到表彰"
	}

	return ans
}
