package config

var conf BotConfig

func GetLLMConfig() *LLMConf {
	return &conf.LLMConf
}
