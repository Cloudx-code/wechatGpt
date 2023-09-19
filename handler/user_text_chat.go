package handler

import (
	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service"
	"wechatGpt/service/administrator"
	"wechatGpt/service/authority"

	"github.com/eatmoreapple/openwechat"
)

type UserMsgService struct {
	// 接收到消息
	msg *openwechat.Message
	// 发送消息的用户
	sender *openwechat.User
}

func NewUserMsgHandler(ctx *openwechat.MessageContext) (*UserMsgService, error) {
	msg := ctx.Message
	sender, err := msg.Sender()
	if err != nil {
		logs.Error("fail to msg.Sender(),[NewUserMsgHandler],err:%v", err)
		return nil, err
	}
	return &UserMsgService{
		msg:    msg,
		sender: sender,
	}, nil
}

func (u *UserMsgService) HandleMsg() {
	// 前置校验权限
	if reply, err := authority.NewManageAuthorityService(u.sender.AvatarID(), u.sender.NickName).CheckAuthority(); err != nil && len(reply) == 0 {
		utils.Reply(u.msg, reply)
		return
	}

	var reply string
	chatStatus := local_cache.GetChatStatus(u.sender.AvatarID())

	switch chatStatus {
	case consts.Administrator:
		reply = administrator.NewAdministratorService(u.sender.AvatarID(), u.sender.NickName, u.msg.Content).HandlerMsg()
	default:
		reply = service.NewUserChatService(u.sender.AvatarID(), u.sender.NickName, u.msg.Content, chatStatus).HandleMsg()
	}
	utils.Reply(u.msg, reply)
}
