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
	Target  string `json:"target"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
}

// PostArchive 用户上传的数据
type PostArchive struct {
	Table
	Hide     string `json:"hide"`
	Contents string `json:"contents" gorm:"type:text"`
}
