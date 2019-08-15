package module

import (
	"github.com/jinzhu/gorm"
)

// Archive ...
type Archive struct {
	gorm.Model
	PostArchive
	Views int `json:"views"`
}

// PostArchive 用户上传的数据
type PostArchive struct {
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
