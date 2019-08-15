package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

func initDB(db *gorm.DB) *gorm.DB {
	if !db.HasTable(&module.Archive{}) {
		db.CreateTable(&module.Archive{})
	}
	return db
}
