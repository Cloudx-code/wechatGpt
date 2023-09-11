package bootstrap

import (
	"fmt"
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/handler"

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

	// 获取所有的好友
	friends, err := self.Friends()
	fmt.Println(friends, err)

	// 获取所有的群组
	groups, err := self.Groups()
	fmt.Println(groups, err)
	go sendMsg2Active(friends)
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}

func sendMsg2Active(friends openwechat.Friends) {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		//friends.First().SendText("hello")
		err := friends.SearchByNickName(1, "指尖轻挑").SendText("我没有！！！\n我不齐！！！\n晓得哒！！！\n我睡滴！！！\n不管我！！！")
		err = friends.SearchByNickName(1, "sub(n,n,17)").SendText("测试")
		if err != nil {
			logs.Error("fail to send friend,err:%v", err)
		} else {
			logs.Info("success to send friend")
		}

	}
}
