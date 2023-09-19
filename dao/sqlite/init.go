package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBProxy *gorm.DB

func Init(fileName string) {
	var err error
	DBProxy, err = gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	DBProxy.AutoMigrate(&User{})
}
