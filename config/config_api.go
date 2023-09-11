package config

var (
	StaticConf  BotStaticConfig
	DynamicConf BotDynamicConfig
)

func GetLLMConfig() *LLMConf {
	return &StaticConf.LLMConf
}
