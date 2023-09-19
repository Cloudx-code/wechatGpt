package handler

import (
	"wechatGpt/common/logs"

	"github.com/eatmoreapple/openwechat"
)

func RegisterHandler() (msgFunc func(msg *openwechat.Message), err error) {
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 处理私聊文本信息
	dispatcher.RegisterHandler(func(msg *openwechat.Message) bool {
		return msg.IsSendByFriend() && msg.IsText()
	}, HandleUserMessage())

	// 处理群聊文本信息
	dispatcher.RegisterHandler(func(msg *openwechat.Message) bool {
		return msg.IsSendByGroup() && msg.IsText()
	}, HandleGroupMessage())

	return dispatcher.AsMessageHandler(), nil
}

func HandleUserMessage() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msgHandler, err := NewUserMsgHandler(ctx)
		if err != nil {
			logs.Error("fail to NewGroupMsgHandler(ctx),err:%v", err)
			// todo 给特定人员发消息，比如我
			return
		}
		msgHandler.HandleMsg()
	}
}

func HandleGroupMessage() func(ctx *openwechat.MessageContext) {
	return func(ctx *openwechat.MessageContext) {
		msgHandler, err := NewGroupMsgHandler(ctx)
		if err != nil {
			logs.Error("fail to NewGroupMsgHandler(ctx),err:%v", err)
			// todo 给特定人员发消息，比如我
			return
		}
		msgHandler.HandleMsg()
	}
}
