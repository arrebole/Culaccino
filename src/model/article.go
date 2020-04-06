package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章
type Article struct {
	gorm.Model
	Name    string `json:"name" gorm:"unique"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Tag     string `json:"tag"`
}

// URL ...
func (p Article) URL() string {
	return "/api/article/" + p.Name
}
