package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/src/service"
)

// GetSection ...
func GetSection(w http.ResponseWriter, r *http.Request) {

	// 想要定位一个文章章节，需要文章名称和章节名称
	articleName := Params(r.Context(), "article")
	sectionName := Params(r.Context(), "section")

	response := service.SectionsRepo.Find(articleName, sectionName)

	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(response))
}

// UpdateSection ...
func UpdateSection(w http.ResponseWriter, r *http.Request) {

	section, err := BodyPaserSection(r.Body)
	if err != nil || section.Name == "" {
		ErrorHandle(w, r)
		return
	}

	// 找到文章的数据，将新建的章节与文章关联
	articleName := Params(r.Context(), "article")
	article := service.ArticlesRepo.Find(articleName)

	// 关联文章与章节
	section.ArticleID = article.ID
	section.Article = article.Article

	response := service.SectionsRepo.Update(section)

	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(response))
}

// CreateSection ...
func CreateSection(w http.ResponseWriter, r *http.Request) {

	section, err := BodyPaserSection(r.Body)
	if err != nil || section.Name == "" {
		ErrorHandle(w, r)
		return
	}

	// 找到文章的数据，将新建的章节与文章关联
	articleName := Params(r.Context(), "article")
	article := service.ArticlesRepo.Find(articleName)

	// 关联文章与章节
	section.ArticleID = article.ID
	section.Article = article.Article

	response, err := service.SectionsRepo.Create(section)
	if err != nil {
		ErrorHandle(w, r)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(response))
}
