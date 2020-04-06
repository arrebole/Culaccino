package service

import (
	"os"

	"github.com/jinzhu/gorm"

	// sqlite3 驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func dbName() string {
	if len(os.Getenv("DBName")) == 0 {
		return "culaccino.db"
	}
	return os.Getenv("DBName")
}

func instance() *gorm.DB {
	if db == nil {
		db = connect()
	}
	return db
}

func connect() *gorm.DB {
	result, err := gorm.Open("sqlite3", dbName())
	if err != nil {
		panic("failed to connect database")
	}

	return result
}
