package repos

import (
	"github.com/arrebole/culaccino/src/model"
)

// ArticlesRepository ...
type ArticlesRepository struct{}

// Find ...
func (p ArticlesRepository) Find(name string) *model.FullArticle {
	var article = model.Article{}
	db.Where("name = ?", name).First(&article)
	return article.Full(p.findSectionLinks(&article))
}

// Finds ...
func (p ArticlesRepository) Finds(offset int, limit int) *model.FullArticles {

	var articles []model.Article
	db.Offset(offset).Limit(limit).Find(&articles)

	var fullArticles = []model.FullArticle{}
	for _, v := range articles {
		fullArticles = append(fullArticles, *v.Full(p.findSectionLinks(&v)))
	}

	var total = 0
	db.Model(&model.Article{}).Count(&total)

	pagination := &model.Pagination{
		TotalSize:  total,
		RemainSize: total - offset - len(articles),
	}

	result := &model.FullArticles{
		Articles:   fullArticles,
		Pagination: pagination,
	}

	return result
}

// Create ...
func (p ArticlesRepository) Create(article *model.Article) (*model.FullArticle, error) {
	article.URL = "/api/aeticles/" + article.Name
	if err := db.Omit("id").Create(article).Error; err != nil {
		return nil, err
	}
	return article.Full([]model.Link{}), nil
}

// Update ...
func (p ArticlesRepository) Update(article *model.Article) *model.FullArticle {
	db.Model(article).Where("name = ?", article.Name).Omit("name", "url").Updates(article)
	return article.Full(p.findSectionLinks(article))
}

// FindSectionLinks for article
func (p ArticlesRepository) findSectionLinks(article *model.Article) []model.Link {
	// 找到section列表
	var sections []model.Section
	db.Where("article_id = ?", article.ID).Find(&sections)

	var result = []model.Link{}
	for _, v := range sections {
		result = append(result, v.Link())
	}
	return result
}
