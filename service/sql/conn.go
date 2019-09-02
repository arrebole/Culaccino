package sql

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/jinzhu/gorm"
)

// 生成数据库链接对象
func connSQL() *gorm.DB {
	db, err := gorm.Open("sqlite3", config.DBName)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
