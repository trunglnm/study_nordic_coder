package main

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100)"`
	Completed bool
}

type User struct {
	gorm.Model
	Name string
}
