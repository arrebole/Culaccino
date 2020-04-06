package model

import "github.com/jinzhu/gorm"

// Section 文章中的一个篇幅
type Section struct {
	gorm.Model
	ArticleID uint    `josn:"article_id"`
	Article   Article `gorm:"foreignkey:ArticleID"`
	Name      string  `json:"name"`
	Content   string  `json:""content`
}

// URL ...
func (p Section) URL() string {
	return "/api/articles/" + p.Article.Name + "/" + p.Name
}
