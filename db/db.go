package db

import (
	"github.com/jinzhu/gorm"
	. "github.com/zjw1918/go-web/model"
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

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func GetDB() *gorm.DB {
	return db
}

