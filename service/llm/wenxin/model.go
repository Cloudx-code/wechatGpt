package wenxin

/*
	详情可参考：https://cloud.baidu.com/doc/WENXINWORKSHOP/s/jlil56u11
*/

type wenXinYiYanResp struct {
	Id          string `json:"id,omitempty"`
	Object      string `json:"object,omitempty"`
	Created     int    `json:"created,omitempty"`
	SentenceId  int    `json:"sentence_id,omitempty"`
	IsEnd       bool   `json:"is_end,omitempty"`
	IsTruncated bool   `json:"is_truncated,omitempty"`
	Result      string `json:"result,omitempty"`
	// 表示用户输入是否存在安全，是否关闭当前会话，清理历史会话信息。true：是，表示用户输入存在安全风险，建议关闭当前会话，清理历史会话信息 false：否，表示用户输入无安全风险
	NeedClearHistory bool `json:"need_clear_history,omitempty"`
	// 当need_clear_history为true时，此字段会告知第几轮对话有敏感信息，如果是当前问题，ban_round=-1
	BanRound int   `json:"ban_round,omitempty"`
	Usage    usage `json:"usage"`
}

type usage struct {
	// 问题tokens数
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// 回答tokens数
	CompletionTokens int `json:"completion_tokens,omitempty"`
	//tokens总数
	TotalTokens int `json:"total_tokens,omitempty"`
}

// 请求文心一言Request中的body内容
type wenXinYiYanReqBody struct {
	// 唯一的必填字段
	Messages    []message `json:"messages,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	TopP        float64   `json:"top_p,omitempty"`
	// 通过对已生成的token增加惩罚，减少重复生成的现象。说明： （1）值越大表示惩罚越大 （2）默认1.0，取值范围：[1.0, 2.0]
	PenaltyScore float64 `json:"penalty_score,omitempty"`
	Stream       bool    `json:"stream,omitempty"`
	// 表示最终用户的唯一标识符，可以监视和检测滥用行为，防止接口恶意调用
	UserId string `json:"user_id,omitempty"`
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AccessTokenResp 百度access_token返回内容
type AccessTokenResp struct {
	AccessToken      string  `json:"access_token,omitempty"`
	ExpiresIn        int     `json:"expires_in,omitempty"`
	Error            *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}
