package orm

import (
	"os"

	"github.com/arrebole/culaccino/src/model"
	"github.com/jinzhu/gorm"

	// sqlite3 驱动
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// dbName 从环境变量中获取 DB_NAME, 没有则返回默认值 culaccino.db；
func dbName() string {
	if len(os.Getenv("DB_NAME")) == 0 {
		return "culaccino.db"
	}
	return os.Getenv("DB_NAME")
}

// Connect gorm 连接数据库
func Connect() *gorm.DB {
	result, err := gorm.Open("sqlite3", dbName())
	if err != nil {
		panic(err.Error())
	}

	return initTable(result)
}

// initTable  If it does not exist it will create the table
func initTable(db *gorm.DB) *gorm.DB {
	// 初始化 articles 表
	if !db.HasTable(&model.Article{}) {
		db.CreateTable(&model.Article{})
	}

	// 初始化 users 表
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
		db.Create(&model.User{Username: "root", Password: "123456"})
	}

	return db
}
