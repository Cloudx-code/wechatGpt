package handler

import (
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/service/chat_manage"

	"github.com/eatmoreapple/openwechat"
)

type UserChat struct {
	UserId int
}

func HandleUserMessage() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		sender, err := msg.Sender()
		if err != nil {
			logs.Error("fail to sender,err:", err)
			utils.Reply(msg, "获取信息出错，请咨询相关人员")
			return
		}
		chatService := chat_manage.NewNormalChatService(sender.AvatarID(), msg.Content)
		reply, err := chatService.Chat()
		if err != nil {
			utils.Reply(msg, "聊天模式出错，错误为："+err.Error())
			return
		}
		utils.Reply(msg, reply)
	}
}
