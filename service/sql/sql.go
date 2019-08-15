package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	Close()
	Count() int
	Delete(uint)
	Add(*module.PostArchive)
	Update(uint, *module.Archive)
	Query(uint) *module.Archive
	QueryDir(int, int) []module.Archive
	QueryDirAll() []module.Archive
}

type client struct {
	db *gorm.DB
}

// Query 查询
func (p *client) Query(id uint) *module.Archive {
	var result = &module.Archive{}
	p.db.First(result, id)
	p.increaseAccess(result)
	return result
}

// // QueryDir 查询目录
func (p *client) QueryDir(limit int, page int) []module.Archive {
	var dir []module.Archive
	var searchItems = "id, title, author, target, cover, summary, views, updated_at, created_at"
	p.db.Limit(limit).Offset(limit * page).Order("id desc").Select(searchItems).Find(&dir)
	return dir
}

// Count 查询文章总量
func (p *client) Count() int {
	var result int
	p.db.Model(&module.Archive{}).Count(&result)
	return result
}

// 在QueryDir中用于计算剩余的数目
func (p *client) remaining(limit int, page int) int {
	remain := p.Count() - limit*(page+1)
	return max(remain, 0)
}

// QueryDirAll 查询目录
func (p *client) QueryDirAll() []module.Archive {
	var dir []module.Archive
	p.db.Select("id, title, author, target, cover, summary, views,updated_at, created_at").Find(&dir)
	return dir
}

// Delete 删除文章
func (p *client) Delete(id uint) {
	p.db.Where("id = ?", id).Delete(&module.Archive{})
}

// Add 添加文章
func (p *client) Add(Archive *module.PostArchive) {
	aNewArchive := module.ToArchive(Archive)
	p.db.Create(aNewArchive)
}

// Update修改文章,只更新修改的字段
func (p *client) Update(id uint, Archive *module.Archive) {
	p.db.Model(&module.Archive{}).Where("id = ?", id).Updates(Archive)
}

// 增加文章浏览量
func (p *client) increaseAccess(Archive *module.Archive) {
	if Archive.Title != "" {
		Archive.Views++
		p.db.Model(Archive).UpdateColumn("views", Archive.Views)
	}
}

// Close ...
func (p *client) Close() {
	p.db.Close()
}

var clientInstance *client

// ConnSQL 连接数据库
func ConnSQL() {
	clientInstance = &client{
		db: initDB(connSQL()),
	}
}

func init() {
	ConnSQL()
}

// New 创建一个数据库客户端
func New() SQL {
	return clientInstance
}
