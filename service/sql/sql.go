package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	ArchiveCount() *module.Count
	ArchiveDelete(uint)
	ArchiveCreate(*module.PostArchive)
	ArchiveUpdate(uint, *module.Archive)
	ArchiveQuery(uint) *module.Archive
	ArchiveQueryDir(int, int) ([]module.Archive, *module.Count)
	UserExist(username string, password string) bool
}

var clientInstance *client

func init() {
	clientInstance = &client{
		db: initDB(connSQL()),
	}
}

type client struct {
	db *gorm.DB
}

// New 创建一个数据库客户端
func New() SQL {
	return clientInstance
}
