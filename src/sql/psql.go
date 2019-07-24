package sql

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/jinzhu/gorm"
)

// SQL ...
type SQL interface {
	Delete(uint)
	Add(*module.PostArticle)
	Update(*module.Article)
	Query(uint) *module.Article
}

type client struct {
	db *gorm.DB
}

// Query 查询
func (p *client) Query(id uint) *module.Article {
	var result = &module.Article{}
	p.db.First(result, id)
	return result
}

// Delete 删除文章
func (p *client) Delete(id uint) {
	p.db.Where("id = ?", id).Delete(&module.Article{})
}

// Add 添加文章
func (p *client) Add(article *module.PostArticle) {
	p.db.Create(module.ToArticle(article))
}

// Update修改文章,只更新修改的字段
func (p *client) Update(article *module.Article) {
	p.db.Model(&module.Article{}).Updates(article)
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
