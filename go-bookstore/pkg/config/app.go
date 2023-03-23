package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//  point of this file is to return a variable DB which will help other commands work with the DB

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:grindtime@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
