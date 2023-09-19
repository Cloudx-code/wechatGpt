package config

var (
	StaticConf  BotStaticConfig
	DynamicConf BotDynamicConfig
)

func GetLLMConfig() *LLMConf {
	return &StaticConf.LLMConf
}

func GetAuthorityList() []string {
	return []string{"675102103", "675103903", "675108417"}
	//return DynamicConf.AuthorityList
}
