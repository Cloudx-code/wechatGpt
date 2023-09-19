package bootstrap

import (
	"fmt"
	"math/rand"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/handler"
	"wechatGpt/service/self_bot"

	"github.com/eatmoreapple/openwechat"
)

/*
	@Author: xy
	@Date: 2023/8/4
	@Desc: 微信启动程序
*/

func Run() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册消息处理函数
	handler, err := handler.RegisterHandler()
	if err != nil {
		logs.Error("fail to handler.RegisterHandler,err:%v", err)
		return
	}
	bot.MessageHandler = handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 初始化机器人本身的功能
	self_bot.InitBotSelf(self)
	logs.Info("初始化机器人成功！")
	// 定时发消息
	go sendMsg2Active()
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

func sendMsg2Active() {
	var reply string
	var err error
	botSelf := self_bot.NewBotOperateService()
	msgList := []string{"我没有！！！", "我不齐！！！", "晓得哒！！！", "我睡滴！！！", "不管我！！！"}
	ticker := time.NewTicker(30 * time.Minute)
	for range ticker.C {
		rand.Seed(time.Now().UnixNano())
		reply, err = botSelf.SendMsg2Friends("指尖轻挑", msgList[rand.Intn(len(msgList))])
		if err != nil {
			logs.Error("fail to send friend old run,err:%v", err)
			botSelf.SendMsg2Friends("sub(n,n,17)", reply)
			continue
		}
		reply, err = botSelf.SendMsg2Friends("sub(n,n,17)", msgList[rand.Intn(len(msgList))])
		if err != nil {
			logs.Error("fail to send friend xy,err:%v", err)
			botSelf.SendMsg2Friends("指尖轻挑)", reply)
			continue
		}
		logs.Info("success to send friend")
	}
}
