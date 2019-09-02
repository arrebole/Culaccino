package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	ArchiveCount() *module.Count
	ArchiveDelete(uint)
	ArchiveCreate(*module.PostArchive, *session.Session)
	ArchiveUpdate(uint, *module.Archive)
	ArchiveDir(int, int) ([]module.Archive, *module.Count)
	ArchiveQueryByID(uint) *module.Archive
	ArchiveQueryByAuthorID(uint) ([]module.Archive, *module.Count)
	UserBaseInfo(username string, password string) *session.UserInfo
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
