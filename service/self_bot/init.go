package self_bot

import (
	"github.com/eatmoreapple/openwechat"
)

// 机器人本身相关信息
var botSelf *openwechat.Self

func InitBotSelf(bot *openwechat.Self) {
	botSelf = bot
}

func GetBotProxy() *openwechat.Self {
	return botSelf
}
