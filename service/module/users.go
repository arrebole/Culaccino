package module

import (
	"github.com/jinzhu/gorm"
)

// User 用户数据库
type User struct {
	gorm.Model
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Permission   int    `json:"permission"`
	Status       int    `json:"status"`
	Avatar       string `json:"avatar"`
	Popular      int    `json:"popular"`
	Star         int    `json:"star"`
	Hate         int    `json:"hate"`
	ArchiveCount int    `json:"archive_count"`
	ArchiveIDs   string `json:"archive_ids"`
	Options      string `json:"options"`
}
