package model

// 大语言模型通用返回结构
type LLMResponse struct {
	Content    string // 大模型返回结果
	CostToken  int64  // 消耗token数
	CostDetail string // 开销详情
	Desc       string // 模型描述
}
