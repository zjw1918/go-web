package db

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Init()  {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

