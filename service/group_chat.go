package service

import (
	"errors"
	"fmt"
	"strings"

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
	@Desc: 群聊处理service，这一层完全与微信解耦合，聚焦于内容本身
*/

type GroupChatService struct {
	GroupId    string            // 群ID
	SenderId   string            // 发送人ID
	SenderName string            // 发送人昵称
	Content    string            // 聊天消息
	ChatStatus consts.ChatStatus // 当前聊天状态（不同状态，机器人的操作不一样）
}

func NewGroupChatService(groupId, senderId, senderName, content string, chatStatus consts.ChatStatus) *GroupChatService {
	return &GroupChatService{
		GroupId:    groupId,
		Content:    content,
		SenderId:   senderId,
		SenderName: senderName,
		ChatStatus: chatStatus,
	}
}

func (g *GroupChatService) HandleMsg() string {
	// 选择模式（聊天？切换模式？）
	logs.Info("用户文本：%v，status:%v", g.Content, g.ChatStatus)
	// 校验特定用语
	reply := g.CheckSpecialText()
	if len(reply) > 0 {
		return reply
	}
	switch g.ChatStatus {
	case consts.NormalChat:
		chatService := chat_manage.NewNormalChatService(g.GroupId, g.Content)
		reply, err := chatService.Chat()
		if err != nil {
			return fmt.Sprintf("聊天模式出错，错误为：%v", err.Error())
		}
		return reply
	case consts.ModelChoose:
		err := g.ChangeModel()
		if err != nil {
			return fmt.Sprintf("模式切换出错，错误为：%v", err.Error())
		}
		return fmt.Sprintf("模型切换成功，当前模型为：%v", g.Content)
	case consts.CreateImage:
		urlInfo, err := image.CreateImage(g.Content)
		if err != nil {
			return fmt.Sprintf("生成图片有误，错误为：%v", err.Error())
		}
		return fmt.Sprintf("图片url为：%v", urlInfo)
	}

	return ""
}

// CheckSpecialText 校验特定话术
func (g *GroupChatService) CheckSpecialText() string {
	if g.ChatStatus == consts.ModelChoose {
		return ""
	}
	if aimStr, ok := utils.ContainStrArray(g.Content, []string{"模型选择", "模型切换"}); ok {
		g.Content = strings.Replace(g.Content, aimStr, "", -1)
		g.Content = strings.Replace(g.Content, "+", "", -1)
		g.Content = strings.TrimSpace(g.Content)
		err := g.ChangeModel()
		if err != nil {
			return ""
		}
		return "模型切换成功，当前模型为：" + g.Content
	}
	switch g.Content {
	case "模型选择":
		local_cache.SetChatStatus(g.GroupId, consts.ModelChoose)
		return consts.ModelIntroduce
	case "帮助", "help", "Help", "HELP":
		return consts.ModelHelpText
	case "给俞越打招呼":
		return "你好，鱼跃、happy、fish moon"
	}
	return ""
}

// ChangeModel 模型选择
func (g *GroupChatService) ChangeModel() error {
	logs.Info("start GroupChatService ChangeModel,info:%v", utils.Encode(g))
	status := consts.NormalChat
	modelName := consts.ModelName(g.Content)
	if _, ok := consts.ModelInfoMap[modelName]; !ok {
		return errors.New("模型不存在，请重新输入")
	}
	if modelName == consts.ModelNameAdministrator {
		if !utils.InStrArray(g.GroupId, config.GetAuthorityList()) {
			return errors.New("模型不存在，请重新输入")
		}
		status = consts.Administrator
	}

	local_cache.SetCurrentModel(g.GroupId, modelName)
	local_cache.SetChatStatus(g.GroupId, status)
	return nil
}
