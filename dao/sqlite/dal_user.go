package sqlite

import (
	"time"

	"wechatGpt/common/logs"
	"wechatGpt/model"

	"gorm.io/gorm"
)

// DalUser user表相关操作
type DalUser struct {
}

func NewDalUser() *DalUser {
	return &DalUser{}
}

// Register 注册功能
func (d *DalUser) Register(username, senderId string, duration time.Duration) error {
	now := time.Now()
	user := &model.User{
		UserID:       senderId,
		Username:     username,
		RegisteredAt: now,
		Mode:         3,
		ExpiresAt:    now.Add(duration),
	}
	if err := DBProxy.Create(user).Error; err != nil {
		logs.Error("fail to Create User,userName:%v,senderId:%v", username, senderId)
		return err
	}
	return nil
}

func (d *DalUser) GetUserByName(userName string) ([]*model.User, error) {
	var userList []*model.User
	if err := DBProxy.Where("user_name=?", userName).Find(&userList).Error; err != nil {
		logs.Error("fail to GetUserByName,err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return userList, nil
}

func (d *DalUser) GetUserById(userId string) (*model.User, error) {
	var userInfo *model.User
	if err := DBProxy.Where("user_id=?", userId).First(&userInfo).Error; err != nil {
		logs.Error("fail to GetUserById,err:%v", err)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return userInfo, nil
}

// CanAccessService 检查用户的服务有效性
func (d *DalUser) CanAccessService(userID uint) (bool, error) {
	var user model.User
	if err := DBProxy.Where("user_id=?", userID).First(&user).Error; err != nil {
		return false, err
	}

	if time.Now().After(user.ExpiresAt) {
		return false, nil
	}

	return true, nil
}
