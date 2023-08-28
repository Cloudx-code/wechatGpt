package handler

import (
	"fmt"
	"strings"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/chat_manage"

	"github.com/eatmoreapple/openwechat"
)

type GroupMsgHandler struct {
	// 获取自己
	self *openwechat.Self
	// 群
	group *openwechat.Group
	// 接收到消息
	msg *openwechat.Message
	// 发送消息的用户
	sender *openwechat.User
}

func NewGroupMsgHandler(ctx *openwechat.MessageContext) (*GroupMsgHandler, error) {
	msg := ctx.Message
	sender, err := msg.Sender()
	if err != nil {
		logs.Error("fail to msg.Sender(),[NewGroupMsgHandler]")
		return nil, err
	}

	group := &openwechat.Group{User: sender}
	groupSender, err := msg.SenderInGroup()
	if err != nil {
		return nil, err
	}
	return &GroupMsgHandler{
		self:   sender.Self(),
		group:  group,
		msg:    msg,
		sender: groupSender,
	}, nil
}

func (g *GroupMsgHandler) handleMsg() {
	if !g.msg.IsAt() {
		return
	}
	// 清楚艾特内容
	g.msg.Content = strings.Replace(g.msg.Content, "@"+g.self.NickName, "", -1)
	// todo 待删
	logs.Info("验证下文本", g.msg.Content)
	// 选择模式（聊天？切换模式？）
	chatStatus := local_cache.GetChatStatus(g.group.AvatarID())
	switch chatStatus {
	case consts.NormalChat:
		reply, err := chat_manage.NewNormalChatService(g.group.AvatarID()).Chat(g.msg.Content)
		if err != nil {
			utils.Reply(g.msg, err.Error())
		}
		utils.Reply(g.msg, reply)
	case consts.ModelChoose:
		err := chat_manage.NewModelChooseService(g.group.AvatarID()).ChangeModel(g.msg.Content)
		if err != nil {
			utils.Reply(g.msg, err.Error())
		}
		utils.Reply(g.msg, fmt.Sprintf("模型切换成功，当前模型为：%v", g.msg.Content))
	}

}
