package handler

import (
	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service"

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
	//// 前置校验权限
	//if reply, err := authority.NewManageAuthorityService(u.sender.AvatarID(), u.sender.NickName).CheckAuthority(); err != nil && len(reply) == 0 {
	//	utils.Reply(u.msg, reply)
	//	return
	//}

	var reply, urlPath string
	chatStatus := local_cache.GetChatStatus(u.sender.AvatarID())

	switch chatStatus {
	//case consts.Administrator:
	//	reply = administrator.NewAdministratorService(u.sender.AvatarID(), u.sender.NickName, u.msg.Content).HandlerMsg()
	case consts.CreateImage:
		// 获取图片url链接
		reply, urlPath = service.NewUserChatService(u.sender.AvatarID(), u.sender.NickName, u.msg.Content, chatStatus).HandleCreateImg()
		utils.Reply(u.msg, reply)
		if urlPath != "" {
			utils.ReplyImage(u.msg, urlPath)
		}
	default:
		reply = service.NewUserChatService(u.sender.AvatarID(), u.sender.NickName, u.msg.Content, chatStatus).HandleNormalChat()
		utils.Reply(u.msg, reply)
	}
}
