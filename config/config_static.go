package config

type BotConfig struct {
	LLMConf LLMConf `yaml:"LLMConf"`
}

type LLMConf struct {
	WeTabEmail string `yaml:"weTabEmail"`
	WeTabPwd   string `yaml:"weTabPwd"`
	WenXinAk   string `yaml:"wenXinAk"`
	WenXinSk   string `yaml:"wenXinSk"`
	GPT35Ak    string `yaml:"GPT35Ak"`
}
