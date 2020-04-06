package service

import (
	"github.com/arrebole/culaccino/src/model"
	"github.com/jinzhu/gorm"
)

// Store ...
type Store struct {
	db *gorm.DB
}

// FindArticle ...
func (p Store) FindArticle(name string) *model.ArticleResponse {
	var article model.Article
	p.db.Where("name = ?", name).First(&article)

	// 创建返回数据
	var result = &model.ArticleResponse{
		Article:      article,
		ArticleURL:   article.URL(),
		SectionLinks: p.FindSectionLink(article.Name),
	}
	return result
}

// FindArticles ...
func (p Store) FindArticles(offset uint, limit uint) *model.ArticlesResponse {
	var articles []model.Article
	p.db.Offset(offset).Limit(limit).Find(&articles)

	var count = 0
	p.db.Model(&model.Article{}).Count(&count)

	var result = &model.ArticlesResponse{}
	for _, v := range articles {
		result.Articles = append(result.Articles, *p.FindArticle(v.Name))
	}

	result.Pagination.Total = count
	result.Pagination.Pages = count / int(limit)
	result.Pagination.Current = int(offset / limit)
	return result
}

// CreateArticle ...
func (p Store) CreateArticle(article *model.Article) *model.ArticleResponse {
	p.db.Create(article)

	var result = &model.ArticleResponse{
		Article:      *article,
		ArticleURL:   article.URL(),
		SectionLinks: p.FindSectionLink(article.Name),
	}

	return result
}

// UpdateArticle ...
func (p Store) UpdateArticle(article *model.Article) *model.ArticleResponse {
	p.db.Save(article)

	var result = &model.ArticleResponse{
		Article:      *article,
		ArticleURL:   article.URL(),
		SectionLinks: p.FindSectionLink(article.Name),
	}
	return result
}

// FindSectionLink for articleID
func (p Store) FindSectionLink(name string) []model.HetoasSectionLink {
	// 找到article
	var article model.Article
	p.db.Where("name = ?", name).First(&article)

	// 找到section列表
	var sections []model.Section
	p.db.Where("article_id = ?", article.ID).Find(&sections)

	// 创建返回类型
	var result []model.HetoasSectionLink
	for _, section := range sections {
		result = append(result, model.HetoasSectionLink{
			Name: section.Name,
			URL:  section.URL(),
		})
	}
	return result
}

// FindSection ...
func (p Store) FindSection(articleName, sectionName string) *model.SectionResponse {
	// 找到article
	var article model.Article
	p.db.Where("name = ?", articleName).First(&article)

	// 找到section
	var section model.Section
	p.db.Where("article_id = ?", article.ID).Where("name = ?", sectionName).First(&section)

	return &model.SectionResponse{
		Section: section,
	}
}

// CreateSection ...
func (p Store) CreateSection(section *model.Section) *model.SectionResponse {
	p.db.Create(section)
	return &model.SectionResponse{
		Section: *section,
	}
}

// UpdateSection ...
func (p Store) UpdateSection(section *model.Section) *model.SectionResponse {
	p.db.Save(section)
	return &model.SectionResponse{
		Section: *section,
	}
}

// CreateStore ...
func CreateStore() Interface {
	return Store{db: instance()}
}
