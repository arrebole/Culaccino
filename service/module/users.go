package module

import "github.com/jinzhu/gorm"

// User 用户数据库
type User struct {
	gorm.Model
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Status       int64  `json:"status"`
	Avatar       string `json:"avatar"`
	PublishCount int64  `json:"publish_count"`
	Popular      int64  `json:"popular"`
	Star         int64  `json:"star"`
	Hate         int64  `json:"hate"`
	Archive      string `json:"archive"`
	Options      string `json:"options"`
}
