package service

import "github.com/arrebole/culaccino/model"

// Interface 数据库操作接口
type Interface interface {
	Get(string) *model.Paper
	Create(*model.Paper) (*model.Paper, error)
	Update(string, *model.Paper) (*model.Paper, error)
	Del(string) error
	Exists(string) bool
	Table(int64) []*model.Paper
}

// New 创建数据库操作对象
func New() Interface {
	return PaperSrv{}
}
