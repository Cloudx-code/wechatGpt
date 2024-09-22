package factory

import (
	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/llm"
	"wechatGpt/service/llm/openai"
	"wechatGpt/service/llm/weTab"
	"wechatGpt/service/llm/wenxin"
)

// GetLLMService 获取具体大语言模型服务
func GetLLMService(sendId string) llm.LLM {
	modelName := local_cache.GetCurrentModel(sendId)

	switch modelName {
	case consts.ModelNameWenXinYiYan:
		return wenxin.NewWenXinYiYanService(sendId)
	case consts.ModelNameWeTab:
		return weTab.NewWeTabService(sendId)
	case consts.ModelNameGPT4:
		return openai.NewGPT4VService(sendId)
	case consts.ModelNameGPT4O:
		return openai.NewGPT4OService(sendId)
	default:
		// 默认文心一言
		return wenxin.NewWenXinYiYanService(sendId)
	}
}
