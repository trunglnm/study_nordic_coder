package main

import (
	"fmt"
	"sync"

	"./models"
	"./services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@/demo_note?charset=utf8&parseTime=True&loc=localhost")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&models.Note{}, &models.User{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer func() {
			recover()
		}()
		defer wg.Done()
		note, err := services.GetNoteById(db, 1)
		fmt.Println(note, err.Error())

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		user, err := services.GetUserById(db, 1)
		fmt.Println(user, err)

	}()
	wg.Wait()
	// time.Sleep(5 * time.Second)
}
