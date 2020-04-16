package model

import "time"

// Article 文章
type Article struct {
	ID uint `json:"id" gorm:"primary_key"`

	Name    string `json:"name" gorm:"unique"`
	Cover   string `json:"cover"`
	Summary string `json:"summary"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// Link ...
func (p Article) Link() Link {
	return Link{
		Name: p.Name,
		URL:  "/api/articles/" + p.Name,
	}
}

// Full ...
func (p Article) Full(links []Link) *FullArticle {
	return &FullArticle{
		Article:  p,
		Sections: links,
	}
}
