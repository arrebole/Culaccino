package repos

import (
	"github.com/arrebole/culaccino/src/model"
)

// SectionsRepository ...
type SectionsRepository struct{}

// Find ...
func (p SectionsRepository) Find(articleName, sectionName string) *model.Section {
	// 找到article
	var article model.Article
	db.Where("name = ?", articleName).First(&article)

	// 找到section
	var result = &model.Section{
		Article:   article,
		ArticleID: article.ID,
	}
	db.Where("article_id = ? AND name = ?", article.ID, sectionName).First(result)

	return result
}

// Create ...
func (p SectionsRepository) Create(section *model.Section) (*model.Section, error) {
	// set section url
	section.URL = "/api/articles/" + section.Article.Name + "/" + section.Name

	if err := db.Create(section).Error; err != nil {
		return nil, err
	}
	return section, nil
}

// Update ...
func (p SectionsRepository) Update(section *model.Section) *model.Section {
	db.Model(section).Where("name = ? AND article_id = ?", section.Name, section.ArticleID).Omit("name", "article_id", "url").Updates(section)
	return section
}
