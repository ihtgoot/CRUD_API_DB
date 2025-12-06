package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "myuser:strong_password@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	return db
}

func GetDB() *gorm.DB {
	return db
}

// package-level singleton + accessor pattern.
// global DB connection
// one shared pointer
// 		one per process
// 		owned by the config package
// 		visible only through exported functions
// “How do I share one database connection across the app without passing it everywhere?”
// 		"Store it in a package and expose a getter.”
// This pattern trades explicitness and testability for simplicity and speed.
//
