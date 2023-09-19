package administrator

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
	"wechatGpt/dao/sqlite"
	"wechatGpt/service/self_bot"
)

type RootService struct {
	Id      string
	Name    string
	Command string
}

func NewAdministratorService(administratorId, administratorName, command string) *RootService {
	return &RootService{Id: administratorId, Name: administratorName, Command: command}
}

func (r *RootService) HandlerMsg() string {
	if self_bot.GetBotProxy() == nil {
		return "当前机器人未初始化"
	}
	logs.Info("check msg:%v", utils.Encode(r))
	reply := r.CheckSpecialText()
	if len(reply) > 0 {
		return reply
	}
	switch r.Command {
	case GetAllFriend:
		logs.Info("start toGetAllFriend")
		reply = self_bot.NewBotOperateService().ShowFriends()
		return reply
	case SendMsg:
		return "功能暂未开发..."
	default:
		return `未知的管理员指令，请输入"帮助"查看详情`
	}
	return ""
}

func (r *RootService) CheckSpecialText() string {
	if aimStr, ok := utils.ContainStrArray(r.Command, []string{"模型选择", "模型切换"}); ok {
		// 清除符号
		r.Command = utils.ClearContent(r.Command, aimStr)
		err := r.ChangeModel()
		if err != nil {
			return ""
		}
		return "模型切换成功，当前模型为：" + r.Command
	}
	// 添加权限
	if aimStr, ok := utils.ContainStrArray(r.Command, []string{"添加权限"}); ok {
		r.Command = utils.ClearContent(r.Command, aimStr)
		return r.AddAuthority(r.Command)
	}
	switch r.Command {
	case "帮助", "help", "Help", "HELP":
		return HelpInfoText
	case "给俞越打招呼":
		return "你好，鱼跃、happy、fish moon"
	}
	return ""
}

// ChangeModel 模型选择
func (r *RootService) ChangeModel() error {
	status := consts.NormalChat
	modelName := consts.ModelName(r.Command)
	if _, ok := consts.ModelInfoMap[modelName]; !ok {
		return errors.New("模型不存在，请重新输入")
	}
	if modelName == consts.ModelNameAdministrator {
		if !utils.InStrArray(r.Id, config.GetAuthorityList()) {
			return errors.New("模型不存在，请重新输入")
		}
		status = consts.Administrator
	}
	local_cache.SetCurrentModel(r.Id, modelName)
	local_cache.SetChatStatus(r.Id, status)
	return nil
}

// AddAuthority 给用户加权限
func (r *RootService) AddAuthority(userId string) string {
	if _, err := strconv.ParseInt(userId, 10, 64); err != nil {
		return fmt.Sprintf("传的用户ID有误，不为整数，ID:%v", userId)
	}
	friendsInfo, err := self_bot.NewBotOperateService().GetAllFriendDetail()
	if err != nil {
		return fmt.Sprintf("获取朋友信息出错，请重试，err:%v", err)
	}
	for _, friend := range friendsInfo {
		if friend.AvatarID() == userId {
			err = sqlite.NewDalUser().Register(friend.NickName, userId, 5*time.Hour)
			if err == nil {
				return "注册成功！"
			}
		}
	}
	return fmt.Sprintf("加权限失败，可能为ID上传有误,ID为：%v，错误信息为：%v", userId, err)
}
