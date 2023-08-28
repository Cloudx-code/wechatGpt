package main

import (
	"wechatGpt/bootstrap"
	"wechatGpt/common/logs"
	"wechatGpt/dao/local_cache"
)

func main() {
	logs.Init(true)
	local_cache.InitCache()
	bootstrap.Run()
}
