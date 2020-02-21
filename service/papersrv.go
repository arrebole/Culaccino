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
	paper.CreateAt = timeNow()
	paper.UpdateAt = paper.CreateAt
	// 如果已经存在 则把创建日期改为原始日期
	if p.Exists(paper.Title) {
		if createAt, err := client.HGet(paper.Title, "create_at").Result(); err == nil {
			paper.CreateAt = createAt
		}
	}

	return client.HMSet(paper.Title, paper.ToMap()).Err()
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
