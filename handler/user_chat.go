package handler

import (
	"time"

	"wechatGpt/common/utils"
	"wechatGpt/service/llm/webTab"

	"github.com/eatmoreapple/openwechat"
)

type UserChat struct {
	UserId int
}

func HandleUserMessage() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		sender, _ := msg.Sender()
		if msg.Content == "帮助" {

		}
		if msg.Content == "清空内容" {
			utils.Reply(msg, "已清空历史聊天记录")
			cache1.Delete(sender.AvatarID())
			return
		}
		// conversationId
		// 获取conversationId
		var conversationId string

		id, ok := cache1.Get(sender.AvatarID())
		if ok {
			conversationId = id.(string)
		}
		reply, conversationId := webTab.GetGpt(msg.Content, conversationId)
		if reply != "" {
			utils.Reply(msg, reply)
		}
		// 缓存conversationId
		cache1.Set(sender.AvatarID(), conversationId, time.Minute*5)
	}
}
