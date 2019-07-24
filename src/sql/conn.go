package sql

import (
	"fmt"

	"github.com/arrebole/culaccino/src/config"
	"github.com/jinzhu/gorm"
)

func sqlConfig() string {
	var confg = config.Cofig.SQL
	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		confg.Host,
		confg.User,
		confg.DBName,
		confg.Password,
	)
}

// 生成数据库链接对象
func connSQL() *gorm.DB {
	db, err := gorm.Open("postgres", sqlConfig())
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
