package sql

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

func initDB(db *gorm.DB) *gorm.DB {
	if !db.HasTable(&module.Archive{}) {
		db.CreateTable(&module.Archive{})
	}
	if !db.HasTable(&module.User{}) {
		db.CreateTable(&module.User{})

		// 添加默认用户
		db.Create(&module.User{
			Username: "root",
			Password: defaultPassWord(),
		})
	}
	return db
}

// 生成数据库链接对象
func connSQL() *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DBAddr)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
