package service

import (
	"time"

	"github.com/arrebole/culaccino/model"
)

// PaperSrv 执行者
type PaperSrv struct{}

// Get ...
func (p PaperSrv) Get(title string) *model.Paper {
	if p.Exists(title) {
		var result = &model.Paper{}
		if data, err := client.HGetAll(title).Result(); err == nil {
			result.BuildFromMap(data)
		}
		return result
	}
	return nil
}

// Set ...
func (p PaperSrv) Set(paper *model.Paper) error {
	if p.Exists(paper.Title) {
		return p.update(paper)
	}
	return p.create(paper)
}

// Del ...
func (p PaperSrv) Del(paperTitle string) error {
	if p.Exists(paperTitle) {
		return client.Del(paperTitle).Err()
	}
	return nil
}

// Exists ...
func (p PaperSrv) Exists(key string) bool {
	if count, _ := client.Exists(key).Result(); count == 1 {
		return true
	}
	return false
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (p PaperSrv) create(paper *model.Paper) error {
	paper.CreateAt = timeNow()
	paper.UpdateAt = paper.CreateAt
	return client.HMSet(paper.Title, paper.ToMap()).Err()
}

func (p PaperSrv) update(paper *model.Paper) error {
	var oldPaper = p.Get(paper.Title)
	paper.CreateAt = oldPaper.CreateAt
	paper.UpdateAt = timeNow()

	if paper.Content == "" {
		paper.Content = oldPaper.Content
	}
	if paper.Cover == "" {
		paper.Cover = oldPaper.Cover
	}
	if paper.Summary == "" {
		paper.Summary = oldPaper.Summary
	}
	if paper.Title == "" {
		paper.Title = oldPaper.Title
	}
	if paper.Type == "" {
		paper.Type = oldPaper.Type
	}
	return client.HMSet(paper.Title, paper.ToMap()).Err()
}
