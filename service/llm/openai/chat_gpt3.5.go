package openai

// ChatGPTResponseBody 响应
type ChatGPTResponseBody struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int                    `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChoiceItem           `json:"choices"`
	Usage   map[string]interface{} `json:"usage"`
}

type ChoiceItem struct {
	Message      message `json:"message"`
	Index        int     `json:"index"`
	Logprobs     int     `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}
type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatGPTRequestBody 请求体
type ChatGPTRequestBody struct {
	Model            string     `json:"model"`
	Messages         []Messages `json:"messages"`
	MaxTokens        uint       `json:"max_tokens"`
	Temperature      float64    `json:"temperature"`
	TopP             int        `json:"top_p"`
	FrequencyPenalty int        `json:"frequency_penalty"`
	PresencePenalty  int        `json:"presence_penalty"`
}
type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
