package service

import (
	"time"

	"github.com/arrebole/culaccino/model"
)

// PaperSrv 执行者
type PaperSrv struct{}

// Get 获取文章
func (p PaperSrv) Get(id string) *model.Paper {
	if p.Exists(id) {
		if data, err := client.HGetAll(id).Result(); err == nil {
			return model.PaperBuilder(data, []string{})
		}
	}
	return nil
}

// Del 删除文章
func (p PaperSrv) Del(id string) error {
	if p.Exists(id) {
		return client.Del(id).Err()
	}
	return nil
}

// Exists 判断文章是否存在
func (p PaperSrv) Exists(id string) bool {
	if count, _ := client.Exists(id).Result(); count == 1 {
		return true
	}
	return false
}

func timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Create 创建一个新的paper
func (p PaperSrv) Create(paper *model.Paper) (*model.Paper, error) {
	paper.CreateAt = timeNow()
	paper.UpdateAt = paper.CreateAt
	return paper, client.HMSet(paper.ID, paper.ToMap()).Err()
}

// Update 更新已有的paper
func (p PaperSrv) Update(id string, paper *model.Paper) (*model.Paper, error) {
	var oldPaper = p.Get(paper.ID)
	paper.CreateAt = oldPaper.CreateAt
	paper.UpdateAt = timeNow()

	// 空字段使用旧数据替换
	paper.Content = fillerStr(paper.Content, oldPaper.Content)
	paper.Summary = fillerStr(paper.Summary, oldPaper.Summary)
	paper.Cover = fillerStr(paper.Cover, oldPaper.Cover)
	paper.Title = fillerStr(paper.Title, oldPaper.Title)
	paper.Type = fillerStr(paper.Type, oldPaper.Type)

	return paper, client.HMSet(paper.ID, paper.ToMap()).Err()
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
