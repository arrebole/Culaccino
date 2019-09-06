package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	UserBaseInfo(username string, password string) *session.User
	NewRepo(*module.PostArchive, *session.User)
	GetRepo(domain string, symbol string) *module.Archive
	CountRepos() *module.Count
	DelRepo(domain string, symbol string) bool
	GetRepos(domain string) ([]module.Archive, *module.Count)
	CommitRepo(domain string, symbol string, repo *module.Archive) bool
	Dashboard(int, int) ([]module.Archive, *module.Count)
}

var _instance *client

func init() {
	_instance = &client{
		db: initDB(connSQL()),
	}
}

type client struct {
	db *gorm.DB
}

// New 创建一个数据库客户端
func New() SQL {
	return _instance
}

//-------------------------------------------------------------------------
