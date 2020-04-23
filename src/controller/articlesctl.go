package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/src/service"
)

// ListArticles ...
func ListArticles(w http.ResponseWriter, r *http.Request) {

	// 获取	URL 中的offset和page_size参数，进行分页处理
	// f("api/articles?limit=80")-> { offset: 0, limit: 10 }
	offset := URLQueryInt(r.URL, "offset", 0)
	limit := URLQueryInt(r.URL, "page_size", 5)

	response := service.ArticlesRepo.FindAll(offset, limit)

	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(response))
}

// CreateArticle ...
func CreateArticle(w http.ResponseWriter, r *http.Request) {

	// 将 post 请求的的json数据解析为对象
	// 提交的 json 中，必须包含 name 数据
	payload, err := BodyPaserArticle(r.Body)
	if err != nil || payload.Name == "" {
		ErrorHandle(w, r)
		return
	}

	response, err := service.ArticlesRepo.Create(payload)
	if err != nil {
		ErrorHandle(w, r)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(response))
}
