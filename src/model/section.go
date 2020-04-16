package model

import "time"

// Section 文章中的一个篇幅
type Section struct {
	ID uint `json:"-" gorm:"primary_key"`

	ArticleID uint `json:"id"`

	Article Article `json:"article" gorm:"foreignkey:ArticleID"`
	Name    string  `json:"name" gorm:"unique"`
	Content string  `json:"content"`
	URL     string  `json:"url"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// Link ...
func (p Section) Link() Link {
	return Link{Name: p.Name, URL: p.URL}
}
