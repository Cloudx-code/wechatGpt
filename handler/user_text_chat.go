package handler

import (
	"fmt"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service"
	"wechatGpt/service/chat_manage"
	"wechatGpt/service/image"

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
	logs.Info("ID：%v", u.sender.AvatarID())
	// 选择模式（聊天？切换模式？）
	chatStatus := local_cache.GetChatStatus(u.sender.AvatarID())
	// 校验特定用语
	reply := service.CheckSpecialText2(u.sender.AvatarID(), u.msg.Content, chatStatus)
	if len(reply) > 0 {
		utils.Reply(u.msg, reply)
		return
	}
	switch chatStatus {
	case consts.NormalChat:
		chatService := chat_manage.NewNormalChatService(u.sender.AvatarID(), u.msg.Content)
		reply, err := chatService.Chat()
		if err != nil {
			utils.Reply(u.msg, "聊天模式出错，错误为："+err.Error())
			return
		}
		utils.Reply(u.msg, reply)
	case consts.ModelChoose:
		err := chat_manage.NewModelChooseService(u.sender.AvatarID()).ChangeModel(u.msg.Content)
		if err != nil {
			utils.Reply(u.msg, "模式切换出错，错误为："+err.Error())
			return
		}
		utils.Reply(u.msg, fmt.Sprintf("模型切换成功，当前模型为：%v", u.msg.Content))
	case consts.CreateImage:
		urlInfo, err := image.CreateImage(u.msg.Content)
		if err != nil {
			utils.Reply(u.msg, "模式切换出错，错误为："+err.Error())
			return
		}
		utils.Reply(u.msg, "图片url为："+urlInfo)
	}
}
