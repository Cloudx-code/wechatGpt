package handler

import (
	"strings"

	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/chat_manage"
)

func CheckSpecialText(senderId, content string, chatStatus consts.ChatStatus) string {
	if chatStatus == consts.ModelChoose {
		return ""
	}
	switch content {
	case "模型选择":
		local_cache.SetChatStatus(senderId, consts.ModelChoose)
		return consts.ModelIntroduce
	case "帮助", "help":
		return consts.ModelHelpText
	case "给俞越打招呼":
		return "你好，鱼跃、happy、fish moon"
	}
	if strings.Contains(content, "模型选择") || strings.Contains(content, "模型切换") {
		content = strings.Replace(content, "模型选择", "", -1)
		content = strings.Replace(content, "+", "", -1)
		content = strings.TrimSpace(content)
		err := chat_manage.NewModelChooseService(senderId).ChangeModel(content)
		if err != nil {
			return ""
		}
		return "模型切换成功，当前模型为：" + content
	}
	return ""
}
