package self_bot

import (
	"fmt"

	"wechatGpt/common/logs"

	"github.com/eatmoreapple/openwechat"
)

/*
	@Author: xy
	@Date: 2023/9/18
	@Desc: 机器人本身相关操作handler，如：发消息给某人、展示好友全部信息、给某好友加权限。
*/

type BotOperateService struct {
	commandMsg string
}

func NewBotOperateService() *BotOperateService {
	return &BotOperateService{}
}

// ShowFriends 展示全部好友
func (b *BotOperateService) ShowFriends() string {
	friends, err := b.GetAllFriendDetail()
	if err != nil {
		return fmt.Sprintf("加载全部好友失败，失败原因为:%v", err.Error())
	}
	var reply string
	for _, friend := range friends {
		logs.Info("好友ID:%v，好友名称:%v\n", friend.AvatarID(), friend.NickName)
		reply += fmt.Sprintf("好友ID:%v，好友名称:%v\n", friend.AvatarID(), friend.NickName)
	}
	return reply
}

// SendMsg2Friends 给特定好友发消息
func (b *BotOperateService) SendMsg2Friends(friendName string, content string) (string, error) {
	var friendsDetail openwechat.Friends
	var err error
	friendsDetail, err = b.GetAllFriendDetail()
	if err != nil {
		return fmt.Sprintf("加载全部好友失败，失败原因为:%v", err.Error()), err
	}
	err = friendsDetail.SearchByNickName(1, friendName).SendText(content)
	if err != nil {
		logs.Error("fail to SendMsg2Friends，err：%v", err)
		return fmt.Sprintf("给好友：%v发消息失败，失败原因为:%v", friendName, err.Error()), err
	}
	return "消息发送成功", err
}

func (b *BotOperateService) GetAllFriendDetail() ([]*openwechat.Friend, error) {
	friends, err := botSelf.Friends()
	if err != nil {
		logs.Error("fail to botSelf.Friends(),err:%v", err)
		return nil, err
	}
	logs.Info("friend Info:%v", friends)
	return friends, nil
}
