package service

import (
	"github.com/arrebole/culaccino/src/model"
)

// Interface 对外接口
type Interface interface {
	// FindArticle 通过article title 查询第一个记录
	FindArticle(string) *model.ArticleResponse

	// FindArticles 查询所有的文章记录列表, 偏移和限制数量
	FindArticles(uint, uint) *model.ArticlesResponse

	// CreateArticle 创建文章
	CreateArticle(*model.Article) *model.ArticleResponse

	// UpdateArticle 更新文章
	UpdateArticle(*model.Article) *model.ArticleResponse

	// FindSection 通过section id 查询具体的Section
	FindSection(string, string) *model.SectionResponse

	// CreateSection 创建session
	CreateSection(*model.Section) *model.SectionResponse

	// UpdateSection 更新section
	UpdateSection(*model.Section) *model.SectionResponse
}
