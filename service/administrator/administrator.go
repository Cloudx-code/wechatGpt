package administrator

import (
	"errors"
	"strings"

	"wechatGpt/common/consts"
	"wechatGpt/common/logs"
	"wechatGpt/common/utils"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
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
		r.Command = strings.Replace(r.Command, aimStr, "", -1)
		r.Command = strings.Replace(r.Command, "+", "", -1)
		r.Command = strings.TrimSpace(r.Command)
		err := r.ChangeModel()
		if err != nil {
			return ""
		}
		return "模型切换成功，当前模型为：" + r.Command
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
