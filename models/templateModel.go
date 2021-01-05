package models

import (
	"gorm.io/gorm"
)

type TestModel struct {
	gorm.Model
}

func AddModel() {
	db.Create(&TestModel{})
}
