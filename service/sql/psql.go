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
	QueryDir(int, int) ([]module.Article, int)
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
func (p *client) QueryDir(limit int, page int) ([]module.Article, int) {
	var dir []module.Article
	var count int
	p.db.Model(&module.Article{}).Count(&count)
	p.db.Limit(limit).Offset(limit * page).Select("id, title, author, target, cover, summary,views").Find(&dir)
	return dir, remaining(count, limit, page)
}

func remaining(all int, limit int, page int) int {
	var result = all - limit*(page+1)
	if result < 0 {
		result = 0
	}
	return result
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
