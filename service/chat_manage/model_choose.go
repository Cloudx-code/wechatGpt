package chat_manage

import (
	"errors"

	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
)

/*
	@Author: xy
	@Date: 2023/8/25
	@Desc: 用于管理模型选择
*/

type ModelChooseService struct {
	senderId string
}

// NewModelChooseService 群聊以群为整体，而不是以个人为整体
func NewModelChooseService(senderId string) *ModelChooseService {
	return &ModelChooseService{senderId: senderId}
}

func (m *ModelChooseService) ChangeModel(modelNameStr string) error {
	modelName := consts.ModelName(modelNameStr)
	if _, ok := consts.ModelInfoMap[modelName]; !ok {
		return errors.New("模型不存在，请重新输入")
	}
	local_cache.SetCurrentModel(m.senderId, modelName)
	local_cache.SetChatStatus(m.senderId, consts.NormalChat)
	return nil
}
