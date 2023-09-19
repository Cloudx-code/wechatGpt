package sqlite

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	UserID       string `gorm:"column:user_id;type:bigint;not null;default:0" json:"user_id"`
	Username     string `gorm:"column:user_name;type:varchar(128);not null" json:"user_name"`
	RegisteredAt time.Time
	ExpiresAt    time.Time
}

// 注册功能
func Register(db *gorm.DB, username, senderId, email string, duration time.Duration) (*User, error) {
	now := time.Now()
	user := &User{
		UserID:       senderId,
		Username:     username,
		RegisteredAt: now,
		ExpiresAt:    now.Add(duration),
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// 检查用户的服务有效性
func CanAccessService(db *gorm.DB, userID uint) (bool, error) {
	var user User
	if err := db.First(&user, userID).Error; err != nil {
		return false, err
	}

	if time.Now().After(user.ExpiresAt) {
		return false, nil
	}

	return true, nil
}
