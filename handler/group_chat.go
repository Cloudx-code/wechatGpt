package handler

import (
	"strings"
	"time"

	"wechatGpt/common/utils"
	"wechatGpt/service/llm/weTab"

	"github.com/eatmoreapple/openwechat"
)

func GroupMessageContextHandler() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		if msg.IsText() && (strings.Contains(msg.Content, "xy的小号") || strings.Contains(msg.Content, "孙尚香菜")) {
			sender, _ := msg.Sender()
			if strings.Contains(msg.Content, "清空内容") {
				utils.Reply(msg, "已清空历史聊天记录")
				cache1.Delete("group" + sender.ID())
			}

			// 获取conversationId
			var conversationId string
			id, ok := cache1.Get("group" + sender.ID())
			if ok {
				conversationId = id.(string)
			}
			reply, conversationId := weTab.GetGpt(strings.Replace(msg.Content, "@孙尚香菜", "", -1), conversationId)
			if reply != "" {
				utils.Reply(msg, reply)
			}
			cache1.Set("group"+sender.ID(), conversationId, time.Minute*5)
		}
	}
}
