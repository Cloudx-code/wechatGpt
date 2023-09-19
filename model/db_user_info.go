package model

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement:true" json:"id"`
	UserID       string `gorm:"column:user_id;type:varchar(128);not null;" json:"user_id"`
	Username     string `gorm:"column:user_name;type:varchar(128);not null" json:"user_name"`
	Mode         int64  `gorm:"column:mode;type:bigint;not null,default:3" json:"mode"` // mode管理，通过bitMap管理权限，默认二进制：11，表示具备文心和weTab权限
	RegisteredAt time.Time
	ExpiresAt    time.Time
}
