package config

type BotStaticConfig struct {
	LLMConf LLMConf `json:"LLMConf"`
}

type LLMConf struct {
	WeTabEmail string `json:"WeTabEmail,omitempty"`
	WeTabPwd   string `json:"WeTabPwd,omitempty"`
	WenXinAk   string `json:"WenXinAk,omitempty"`
	WenXinSk   string `json:"WenXinSk,omitempty"`
	GPTAk      string `json:"GPTAk,omitempty"`
}

// BotDynamicConfig 动态配置中心
type BotDynamicConfig struct {
	// 是否更改配置
	IsChange bool `json:"IsChange"`
	// 管理员配置
	AuthorityList []string `json:"AuthorityList"`
}
