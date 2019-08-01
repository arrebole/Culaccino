package module

import (
	"github.com/jinzhu/gorm"
)

// Article ...
type Article struct {
	gorm.Model
	PostArticle
	Views int `json:"views"`
}

// PostArticle 用户上传的数据
type PostArticle struct {
	Table
	Contents string `json:"contents" gorm:"type:text"`
}

// Table ...
type Table struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Target  string `json:"target"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
}
