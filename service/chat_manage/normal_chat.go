package chat_manage

import (
	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
	factory "wechatGpt/service/llm/llm_factory"
)

/*
	@Author: xy
	@Date: 2023/8/28
	@Desc: 聊天功能
*/

type NormalChatService struct {
	senderId string
}

func NewNormalChatService(senderId string) *NormalChatService {
	return &NormalChatService{senderId: senderId}
}

func (n *NormalChatService) Chat(content string) (string, error) {
	if content == "模型选择" {
		local_cache.SetChatStatus(n.senderId, consts.ModelChoose)
		return consts.ModelIntroduce, nil
	}

	llmModel := factory.GetLLMService(n.senderId)
	llmModel.PreQuery()
	llmReply, err := llmModel.Query(content)
	if err != nil {
		return "", err
	}
	llmModel.PostQuery()
	return llmReply, nil

}
