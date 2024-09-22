package openai

type GPT4VRequestBody struct {
	Model     string         `json:"model"`
	Messages  []GPT4VMessage `json:"messages"`
	MaxTokens int            `json:"max_tokens"`
}

type GPT4ORequestBody struct {
	Model     string         `json:"model"`
	Messages  []GPT4OMessage `json:"messages"`
	MaxTokens int            `json:"max_tokens"`
}

type GPT4OMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPT4VMessage struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string         `json:"id,omitempty"`
	Object  string         `json:"object,omitempty"`
	Created int            `json:"created,omitempty"`
	Model   string         `json:"model,omitempty"`
	Choices []ChoiceItem   `json:"choices,omitempty"`
	Usage   *Usage         `json:"usage,omitempty"`
	Error   *ChatGPTErrMsg `json:"error,omitempty"`
}
type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens,omitempty"`
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	TotalTokens      int64 `json:"total_tokens,omitempty"`
}

type ChoiceItem struct {
	Message      *ChatGPTHistoryMsg `json:"message,omitempty"`
	Index        int                `json:"index,omitempty"`
	FinishReason string             `json:"finish_reason,omitempty"`
}

// ChatGPT请求/响应中的Message,对应先前的上下文
type ChatGPTHistoryMsg struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

// ChatGPT响应错误结构
type ChatGPTErrMsg struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
