package services

import (
	"github.com/jinzhu/gorm"

	"../models"
)

func GetUserById(db *gorm.DB, id int) (*models.User, error) {
	user := &models.User{}
	err := db.Where("id = ?", id).First(user).Error
	return user, err
}
