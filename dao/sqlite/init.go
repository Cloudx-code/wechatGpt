package sqlite

import (
	"fmt"

	"wechatGpt/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBProxy *gorm.DB

func Init(fileName string) {
	var err error
	DBProxy, err = gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database:%v", err))
	}

	// Migrate the schema
	DBProxy.AutoMigrate(&model.User{})
}
