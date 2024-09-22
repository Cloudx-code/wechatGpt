package llm

type LLM interface {
	// PreQuery 查询大语言前置操作
	PreQuery()
	// Query 查询大语言模型获取回复, temperature:温度，控制模型随机性，0-2，设定为0时输出结果稳定
	Query(content string) (string, error)
	// PostQuery 查询大语言模型后置操作
	PostQuery()
	// GetName 获取名称
	GetName() string
}
