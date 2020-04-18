package model

import "time"

// Article 文章
type Article struct {
	ID uint `json:"id" gorm:"primary_key"`

	Name     string `json:"name" gorm:"unique"`
	Cover    string `json:"cover"`
	Summary  string `json:"summary"`
	Tag      string `json:"tag"`
	URL      string `json:"url"`
	Contents string `json:"contents" gorm:"type:TEXT"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}
