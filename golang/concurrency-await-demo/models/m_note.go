package models

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title   string `gorm:"type:varchar(50)"`
	OwnerId int
}
