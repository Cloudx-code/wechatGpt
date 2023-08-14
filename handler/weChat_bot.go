package handler

import (
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/patrickmn/go-cache"
)

var cache1 = cache.New(time.Minute*5, time.Minute*5)

func RegisterHandler() (msgFunc func(msg *openwechat.Message), err error) {
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 私聊
	dispatcher.RegisterHandler(func(msg *openwechat.Message) bool {
		return msg.IsSendByFriend()
	}, UserMessageContextHandler())

	// 群聊
	dispatcher.RegisterHandler(func(msg *openwechat.Message) bool {
		return msg.IsSendByGroup()
	}, GroupMessageContextHandler())
	return dispatcher.AsMessageHandler(), nil

}
