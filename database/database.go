package database

import (
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
}
