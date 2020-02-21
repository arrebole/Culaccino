package service

import "github.com/arrebole/culaccino/model"

// Interface 数据库操作接口
type Interface interface {
	Get(string) *model.Paper
	Set(*model.Paper) error
	Del(string) error
	Exists(string) bool
	Explore(int64) []*model.Paper
}

// New 创建数据库操作对象
func New() Interface {
	return PaperSrv{}
}
