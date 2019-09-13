package model

import (
	"fmt"
	"time"
)

// Chapter Repo的一个章节
type Chapter struct {
	Storage  string    `json:"storage"`
	Repo     string    `json:"repo"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Contents string    `json:"contents"`
}

// Symbol 指向父名
func (p *Chapter) Symbol() string {
	return fmt.Sprintf("%s:%s:%s", p.Storage, p.Repo, p.Name)
}

// Parents 指向父名 dev:a
func (p *Chapter) Parents() string {
	return fmt.Sprintf("%s:%s", p.Storage, p.Repo)
}

// NewChapter 唯一标识
func NewChapter(arg ...string) *Chapter {
	if len(arg) != 3 {
		panic("new chapter arg need 3")
	}
	return &Chapter{
		Storage: arg[0],
		Repo:    arg[1],
		Name:    arg[2],
	}
}
