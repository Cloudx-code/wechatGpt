package main

import (
	"wechatGpt/bootstrap"
	"wechatGpt/common/logs"
	"wechatGpt/config"
	"wechatGpt/dao/local_cache"
)

/*
1) 可配置化
2) 群、好友权限功能
3) 管理员板块设置
4) 数据存储板块设计
*/
func main() {
	logs.Init(true)
	local_cache.InitCache()
	//sqlite.Init("fish_moon.db")
	config.InitStaticConfig("config/static_config.json")
	//config.InitDynamicConfig("config.json")
	//fmt.Println(utils.Encode(config.StaticConf))
	//fmt.Println(utils.Encode(config.DynamicConf))
	//for {
	//
	//}
	bootstrap.Run()
}
