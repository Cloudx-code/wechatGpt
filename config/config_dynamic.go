package config

// BotDynamicConfig 动态配置中心
type BotDynamicConfig struct {
	// 群配置
	GroupInfoMap map[string]GroupDetail `json:"GroupInfoMap"`
}

type GroupDetail struct {
}
