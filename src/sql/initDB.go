package sql

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/jinzhu/gorm"
)

func initDB(db *gorm.DB) *gorm.DB {
	if !db.HasTable(&module.Article{}) {
		db.CreateTable(&module.Article{})
	}
	return db
}
