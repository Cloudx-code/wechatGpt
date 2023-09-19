package local_cache

import (
	"wechatGpt/common/consts"
)

func Get(key string) (interface{}, bool) {
	return cacheProxy.Get(key)
}

func Set(k string, v interface{}) {
	cacheProxy.Set(k, v, 0)
}

// GetChatStatus  定制化处理chatStatus
func GetChatStatus(senderId string) consts.ChatStatus {
	key := consts.RedisKeyChatStatus + senderId
	v, ok := cacheProxy.Get(key)
	if !ok {
		return consts.NormalChat
	}
	if chatStatus, ok := v.(consts.ChatStatus); ok {
		return chatStatus
	}
	return consts.NormalChat
}

func SetChatStatus(senderId string, aimStatus consts.ChatStatus) {
	key := consts.RedisKeyChatStatus + senderId
	cacheProxy.Set(key, aimStatus, 0)
}

// GetCurrentModel 获取当前聊天组模型，默认为文心一言
func GetCurrentModel(senderId string) consts.ModelName {
	key := consts.RedisKeyCurrentModel + senderId
	v, ok := cacheProxy.Get(key)
	if !ok {
		return consts.ModelNameWeTab
	}
	if currentModel, ok := v.(consts.ModelName); ok {
		return currentModel
	}
	return consts.ModelNameWeTab
}

func SetCurrentModel(senderId string, currentModel consts.ModelName) {
	key := consts.RedisKeyCurrentModel + senderId
	cacheProxy.Set(key, currentModel, 0)
}
