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

func NewModelChooseService(senderId string) *ModelChooseService {
	return &ModelChooseService{senderId: senderId}
}

func (m *ModelChooseService) ChangeModel(modelNameStr string) error {
	modelName := consts.ModelName(modelNameStr)
	switch modelName {
	case consts.ModelNameWenXinYiYan:
		local_cache.SetCurrentModel(m.senderId, consts.ModelNameWenXinYiYan)
	case consts.ModelNameWebTab:
		local_cache.SetCurrentModel(m.senderId, consts.ModelNameWebTab)
	default:
		return errors.New("模型不存在，请重新输入")
	}
	return nil
}
