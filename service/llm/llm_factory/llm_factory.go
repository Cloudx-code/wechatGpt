package factory

import (
	"wechatGpt/common/consts"
	"wechatGpt/dao/local_cache"
	"wechatGpt/service/llm"
	"wechatGpt/service/llm/webTab"
	"wechatGpt/service/llm/wenxin"
)

// GetLLMService 获取具体大语言模型服务
func GetLLMService(sendId string) llm.LLM {
	modelName := local_cache.GetCurrentModel(sendId)

	switch modelName {
	case consts.ModelNameWenXinYiYan:
		return wenxin.NewWenXinYiYanService()
	case consts.ModelNameWebTab:
		return webTab.NewWebTabService(sendId)
	default:
		// 默认文心一言
		return wenxin.NewWenXinYiYanService()
	}
}
