package models

import (
	"Gin-Template/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.DatabaseName,
		setting.DatabaseSetting.Charset)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Fail to connect database: %v", err)
	}

	err = db.AutoMigrate(&TestModel{})
	if err != nil {
		log.Fatalf("Fail to migrate database: %v", err)
	}
}
