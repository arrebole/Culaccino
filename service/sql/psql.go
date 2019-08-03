package sql

import (
	"time"

	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	Delete(uint)
	Add(*module.PostArticle)
	Update(uint, *module.Article)
	Query(uint) *module.Article
	QueryDir(uint) []module.Article
	QueryDirAll() []module.Article
}

type client struct {
	db *gorm.DB
}

// Query 查询
func (p *client) Query(id uint) *module.Article {
	var result = &module.Article{}
	p.db.First(result, id)
	p.increaseAccess(result)
	return result
}

// // QueryDir 查询目录
func (p *client) QueryDir(page uint) []module.Article {
	const limit uint = 3
	var dir []module.Article
	p.db.Limit(limit).Offset(limit * page).Select("id, title, author, target, cover, summary,views").Find(&dir)
	return dir
}

// // QueryDirAll 查询目录
func (p *client) QueryDirAll() []module.Article {
	var dir []module.Article
	p.db.Select("id, title, author, target, cover, summary,views").Find(&dir)
	return dir
}

// Delete 删除文章
func (p *client) Delete(id uint) {
	p.db.Where("id = ?", id).Delete(&module.Article{}).Update("DeletedAt", time.Now())
}

// Add 添加文章
func (p *client) Add(article *module.PostArticle) {
	aNewArticle := module.ToArticle(article)
	aNewArticle.CreatedAt = time.Now()
	p.db.Create(aNewArticle)
}

// Update修改文章,只更新修改的字段
func (p *client) Update(id uint, article *module.Article) {
	p.db.Model(&module.Article{}).Where("id = ?", id).Updates(article)
}

// 增加文章浏览量
func (p *client) increaseAccess(article *module.Article) {
	if article.Title != "" {
		article.Views++
		p.db.Model(article).UpdateColumn("views", article.Views)
	}
}

var clientInstance *client

// New 创建一个数据库客户端
func New() SQL {
	if clientInstance == nil {
		clientInstance = &client{
			db: initDB(connSQL()),
		}
	}
	return clientInstance
}
