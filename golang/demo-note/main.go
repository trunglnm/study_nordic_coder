package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/demo_note?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

}
