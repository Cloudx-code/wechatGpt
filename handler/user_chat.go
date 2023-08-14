package handler

import (
	"time"

	"wechatGpt/common/utils"
	"wechatGpt/service/llm/weTab"

	"github.com/eatmoreapple/openwechat"
)

type UserChat struct {
	UserId int
}

func UserMessageContextHandler() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		sender, _ := msg.Sender()

		if msg.Content == "清空内容" {
			utils.Reply(msg, "已清空历史聊天记录")
			cache1.Delete(sender.ID())
			return
		}
		// conversationId
		// 获取conversationId
		var conversationId string

		id, ok := cache1.Get(sender.ID())
		if ok {
			conversationId = id.(string)
		}
		reply, conversationId := weTab.GetGpt(msg.Content, conversationId)
		if reply != "" {
			utils.Reply(msg, reply)
		}
		// 缓存conversationId
		cache1.Set(sender.ID(), conversationId, time.Minute*5)
	}
}
