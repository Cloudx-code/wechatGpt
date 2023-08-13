package main

import (
	"wechatGpt/bootstrap"
	"wechatGpt/common/logs"
)

func main() {
	logs.Init()
	bootstrap.Run()
}
