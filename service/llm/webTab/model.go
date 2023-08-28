package webTab

type WebTabReqBody struct {
	Prompt         string `json:"prompt"`
	AssistantID    string `json:"assistantId"`
	ConversationId string `json:"conversationId"`
}
