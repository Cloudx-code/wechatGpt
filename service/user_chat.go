package service

import (
	"errors"
	"fmt"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/chat_manage"
	"wechatGpt/service/image"
)

/*
	@Author: xy
	@Date: 2023/9/15
	@Desc: 私聊处理service，这一层完全与微信解耦合，聚焦于内容本身
*/

type UserChatService struct {
	SenderId   string            // 发送人ID
	SenderName string            // 发送人昵称
	Content    string            // 聊天消息
	ChatStatus consts.ChatStatus // 当前聊天状态（不同状态，机器人的操作不一样）
}

func NewUserChatService(senderId, senderName, content string, chatStatus consts.ChatStatus) *UserChatService {
	return &UserChatService{
		Content:    content,
		SenderId:   senderId,
		SenderName: senderName,
		ChatStatus: chatStatus,
	}
}

func (u *UserChatService) HandleMsg() string {
	// 选择模式（聊天？切换模式？）
	logs.Info("用户文本：%v，status:%v", u.Content, u.ChatStatus)
	// 校验特定用语
	reply := u.CheckSpecialText()
	if len(reply) > 0 {
		return reply
	}
	switch u.ChatStatus {
	case consts.NormalChat:
		chatService := chat_manage.NewNormalChatService(u.SenderId, u.Content)
		reply, err := chatService.Chat()
		if err != nil {
			return fmt.Sprintf("聊天模式出错，错误为：%v", err.Error())
		}
		return reply
	case consts.ModelChoose:
		err := u.ChangeModel()
		if err != nil {
			return fmt.Sprintf("模式切换出错，错误为：%v", err.Error())
		}
		return fmt.Sprintf("模型切换成功，当前模型为：%v", u.Content)
	case consts.CreateImage:
		urlInfo, err := image.CreateImage(u.Content)
		if err != nil {
			return fmt.Sprintf("生成图片有误，错误为：%v", err.Error())
		}
		return fmt.Sprintf("图片url为：%v", urlInfo)
	}

	return ""
}

// CheckSpecialText 校验特定话术
func (u *UserChatService) CheckSpecialText() string {
	if u.ChatStatus == consts.ModelChoose {
		return ""
	}
	if aimStr, ok := utils.ContainStrArray(u.Content, []string{"模型选择", "模型切换"}); ok {
		u.Content = utils.ClearContent(u.Content, aimStr)
		err := u.ChangeModel()
		if err != nil {
			return ""
		}
		return "模型切换成功，当前模型为：" + u.Content
	}
	switch u.Content {
	case "模型选择":
		local_cache.SetChatStatus(u.SenderId, consts.ModelChoose)
		return consts.ModelIntroduce
	case "帮助", "help", "Help", "HELP":
		return consts.ModelHelpText
	case "给俞越打招呼":
		return "你好，鱼跃、happy、fish moon"
	}
	return ""
}

// ChangeModel 模型选择
func (u *UserChatService) ChangeModel() error {
	logs.Info("start GroupChatService ChangeModel,info:%v", utils.Encode(u))
	status := consts.NormalChat
	modelName := consts.ModelName(u.Content)
	if _, ok := consts.ModelInfoMap[modelName]; !ok {
		return errors.New("模型不存在，请重新输入")
	}
	if modelName == consts.ModelNameAdministrator {
		if !utils.InStrArray(u.SenderId, config.GetAuthorityList()) {
			return errors.New("模型不存在，请重新输入")
		}
		status = consts.Administrator
	}
	local_cache.SetCurrentModel(u.SenderId, modelName)
	local_cache.SetChatStatus(u.SenderId, status)
	return nil
}
