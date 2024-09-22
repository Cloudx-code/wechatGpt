package bootstrap

import (
	"fmt"
	"math/rand"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/config"
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
	msgList := []string{"心跳检测"}
	ticker := time.NewTicker(1 * time.Hour)

	for range ticker.C {
		rand.Seed(time.Now().UnixNano())
		reply, err = botSelf.SendMsg2Friends("sub(n,n,17)", msgList[rand.Intn(len(msgList))])
		if err != nil {
			logs.Error("fail to send friend xy,err:%v", err)
			botSelf.SendMsg2Friends("指尖轻挑)", reply)
			continue
		}
		if config.BanDrinkHours > 0 {
			config.BanDrinkHours--
		} else {
			if IsWorkingDay(time.Now()) && IsWorkingHours(time.Now()) {
				botSelf.SendMsg2Friends("陈彦好", config.DrinkSentence[rand.Intn(len(config.DrinkSentence))]+"\n输入：屏蔽喝水可屏蔽6小时\n输入：我要喝水可接触屏蔽")
			}
		}
		logs.Info("success to send friend")
	}
}

func IsWorkingHours(currentTime time.Time) bool {
	currentHour := currentTime.Hour()
	// 检查当前小时数是否在早上9点到晚上6点之间（不包括6点）
	return currentHour >= 9 && currentHour < 18
}

func IsWorkingDay(date time.Time) bool {
	weekday := date.Weekday()
	// 如果是周六或周日，则不是工作日
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}
	// 其他情况，默认为工作日
	return true
}
