package services

import (
	"../models"
	"github.com/jinzhu/gorm"
)

func GetNoteById(db *gorm.DB, id int) (*models.Note, error) {
	note := &models.Note{}
	err := db.Where("id = ?", id).First(note).Error
	return note, err
}
