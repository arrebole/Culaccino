package sql

import (
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
