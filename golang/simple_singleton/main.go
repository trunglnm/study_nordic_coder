package main

import "sync"

type DB struct {
}

var instance *DB
var mutex sync.Mutex
var once sync.Once

func GetInstance() *DB {
	once.Do(func() {
		instance = &DB{}
	})
	return instance
}

func main() {

}
