package handler

import (
	"strings"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service"
	"wechatGpt/service/administrator"

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
		logs.Error("fail to msg.Sender(),[NewGroupMsgHandler],err:%v", err)
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

// HandleMsg 处理群消息
func (g *GroupMsgHandler) HandleMsg() {
	if !g.msg.IsAt() {
		return
	}
	// 清除艾特的内容
	g.msg.Content = strings.Replace(g.msg.Content, "@"+g.self.NickName, "", -1)
	g.msg.Content = strings.TrimSpace(g.msg.Content)
	if len(g.msg.Content) == 0 {
		return
	}
	var reply string
	chatStatus := local_cache.GetChatStatus(g.group.AvatarID())
	switch chatStatus {
	case consts.Administrator:
		// todo xiongyun，现在整个群都会变管理员！！！
		reply = administrator.NewAdministratorService(g.group.AvatarID(), g.sender.NickName, g.msg.Content).HandlerMsg()
	default:
		reply = service.NewGroupChatService(g.group.AvatarID(), g.sender.AvatarID(), g.sender.NickName, g.msg.Content, chatStatus).HandleMsg()
	}
	utils.Reply(g.msg, reply)
}
