package weTab

type LoginInfo struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResp struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Data      struct {
		Token        string `json:"token,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	} `json:"data"`
}

type WebTabReqBody struct {
	Prompt         string `json:"prompt"`
	AssistantID    string `json:"assistantId"`
	ConversationId string `json:"conversationId"`
}

// ErrStruct {"code":4002,"message":"token无效，过期","timestamp":1693472741522,"data":null}
type ErrStruct struct {
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}
