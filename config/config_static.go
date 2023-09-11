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
