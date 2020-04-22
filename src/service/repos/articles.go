package repos

import (
	"github.com/arrebole/culaccino/src/model"
)

// ArticlesRepository ...
type ArticlesRepository struct{}

// Find ...
func (p ArticlesRepository) Find(name string) *model.Article {
	var article = &model.Article{}
	db.Where("name = ?", name).First(article)
	return article
}

// FindAll ...
func (p ArticlesRepository) FindAll(offset int, limit int) *model.Articles {
	var articles []model.Article

	selectData := []string{
		"id",
		"name",
		"cover",
		"summary",
		"tag",
		"url",
		"created_at",
		"updated_at",
	}
	db.Order("updated_at desc").Offset(offset).Select(selectData).Limit(limit).Find(&articles)

	var total = 0
	db.Model(&model.Article{}).Count(&total)

	pagination := &model.Pagination{
		TotalSize:  total,
		RemainSize: total - offset - len(articles),
	}

	return &model.Articles{
		Articles:   articles,
		Pagination: pagination,
	}
}

// Create ...
func (p ArticlesRepository) Create(article *model.Article) (*model.Article, error) {
	article.URL = "/api/articles/" + article.Name
	if err := db.Omit("id").Create(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

// Update ...
func (p ArticlesRepository) Update(article *model.Article) *model.Article {
	db.Model(article).Where("name = ?", article.Name).Omit("name", "url").Updates(article)
	return p.Find(article.Name)
}

// Delete ...
func (p ArticlesRepository) Delete(name string) *model.Article {
	var article = &model.Article{}
	db.Where("name = ?", name).Delete(article)
	return article
}
