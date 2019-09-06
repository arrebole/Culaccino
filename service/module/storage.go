package module

import (
	"github.com/jinzhu/gorm"
)

// Archive ...
type Archive struct {
	gorm.Model
	PostArchive
	Author       string `json:"author"`
	AuthorID     uint   `json:"author_id"`
	Weights      int    `json:"weigths"`
	DiscussCount int    `json:"discuss_count"`
	DiscussIDs   string `json:"discuss_ids"`
	Views        int    `json:"views"`
}

// Table ...
type Table struct {
	Title   string `json:"title"`
	Area    string `json:"area"`
	Tag     string `json:"tag"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
}

// PostArchive 用户上传的数据
type PostArchive struct {
	Table
	Hide     string `json:"hide"`
	Contents string `json:"contents" gorm:"type:text"`
}

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

// Storage 主存储单元
type Storage struct {
	Symbol string
}

// Repo 仓库
type Repo struct {
	Symbol string
}

// Chapter Repo的一个章节
type Chapter struct {
	Symbol string
}
