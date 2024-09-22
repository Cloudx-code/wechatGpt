package chat_manage

import (
	factory "wechatGpt/service/llm/llm_factory"
)

/*
	@Author: xy
	@Date: 2023/8/28
	@Desc: 聊天功能
*/

type NormalChatService struct {
	senderId string
	content  string
}

func NewNormalChatService(senderId string, content string) *NormalChatService {
	return &NormalChatService{senderId: senderId, content: content}
}

func (n *NormalChatService) Chat() (string, error) {
	llmModel := factory.GetLLMService(n.senderId)
	llmModel.PreQuery()
	llmReply, err := llmModel.Query(n.content)
	if err != nil {
		return "", err
	}
	llmModel.PostQuery()
	llmReply += "\n本轮对话所有模型为：" + llmModel.GetName()
	return llmReply, nil
}
