package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", os.Getenv("MYSQL_URI"))
	if err != nil {
		log.Fatal(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}