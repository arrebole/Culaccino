package sql

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	Count() int
	Delete(uint)
	Add(*module.PostArticle)
	Update(uint, *module.Article)
	Query(uint) *module.Article
	QueryDir(int, int) []module.Article
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
func (p *client) QueryDir(limit int, page int) []module.Article {
	var dir []module.Article
	var searchItems = "id, title, author, target, cover, summary, views, updated_at, created_at"
	p.db.Limit(limit).Offset(limit * page).Order("id desc").Select(searchItems).Find(&dir)
	return dir
}

// Count 查询文章总量
func (p *client) Count() int {
	var result int
	p.db.Model(&module.Article{}).Count(&result)
	return result
}

// 在QueryDir中用于计算剩余的数目
func (p *client) remaining(limit int, page int) int {
	remain := p.Count() - limit*(page+1)
	return max(remain, 0)
}

// QueryDirAll 查询目录
func (p *client) QueryDirAll() []module.Article {
	var dir []module.Article
	p.db.Select("id, title, author, target, cover, summary, views,updated_at, created_at").Find(&dir)
	return dir
}

// Delete 删除文章
func (p *client) Delete(id uint) {
	p.db.Where("id = ?", id).Delete(&module.Article{})
}

// Add 添加文章
func (p *client) Add(article *module.PostArticle) {
	aNewArticle := module.ToArticle(article)
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
