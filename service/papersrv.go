package service

import (
	"time"

	"github.com/arrebole/culaccino/model"
)

// PaperSrv 执行者
type PaperSrv struct{}

// Get 获取文章
func (p PaperSrv) Get(title string) *model.Paper {
	if p.Exists(title) {
		if data, err := client.HGetAll(title).Result(); err == nil {
			return model.PaperBuilder(data, []string{})
		}
	}
	return nil
}

// Set 如果已存在则修改文章内容，如果不存在则创建
func (p PaperSrv) Set(paper *model.Paper) error {
	if p.Exists(paper.Title) {
		return p.update(paper)
	}
	return p.create(paper)
}

// Del 删除文章
func (p PaperSrv) Del(paperTitle string) error {
	if p.Exists(paperTitle) {
		return client.Del(paperTitle).Err()
	}
	return nil
}

// Exists 判断文章是否存在
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

	// 空字段使用旧数据替换
	paper.Content = fillerStr(paper.Content, oldPaper.Content)
	paper.Summary = fillerStr(paper.Summary, oldPaper.Summary)
	paper.Cover = fillerStr(paper.Cover, oldPaper.Cover)
	paper.Title = fillerStr(paper.Title, oldPaper.Title)
	paper.Type = fillerStr(paper.Type, oldPaper.Type)

	return client.HMSet(paper.Title, paper.ToMap()).Err()
}

// fillerStr 返回一个长度不为0,的字符产,
// sample f( "", "", "samp" ) -> "samp"
func fillerStr(strs ...string) string {
	for _, v := range strs {
		if len(v) > 0 {
			return v
		}
	}
	return strs[len(strs)-1]
}
