package controller

import (
	"encoding/json"
	"net/http"

	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// New 创建一篇文章
func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		response := &model.Response{}
		paper, err := getPostData(r)

		// 提交数据错误
		if err != nil {
			w.Write(response.Build(-1, "fail post", nil))
			return
		}
		// 文章已存在
		if service.New().Exists(paper.Title) {
			w.Write(response.Build(-1, "paper Exists", nil))
			return
		}
		// 数据库写入错误
		if err = service.New().Set(paper); err != nil {
			w.Write(response.Build(-1, err.Error(), nil))
			return
		}
		w.Write(response.Build(0, "success", nil))
	}
}

// Update 更新一篇文章
func Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		response := &model.Response{}
		paper, err := getPostData(r)

		// 提交数据错误
		if err != nil {
			w.Write(response.Build(-1, "fail post", nil))
			return
		}
		// 文章不存在
		if !service.New().Exists(paper.Title) {
			w.Write(response.Build(-1, "not find", nil))
			return
		}
		// 数据库写入错误
		if err = service.New().Set(paper); err != nil {
			w.Write(response.Build(-1, err.Error(), nil))
			return
		}
		w.Write(response.Build(0, "success", nil))
	}
}

// Get 查看文章详细内容
func Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		response, title := &model.Response{}, r.FormValue("key")
		if title == "" {
			w.Write(response.Build(-1, "miss query key", nil))
			return
		}

		// 文章不存在
		if !service.New().Exists(title) {
			w.Write(response.Build(-1, "not find key", nil))
			return
		}
		w.Write(response.Build(0, "success", service.New().Get(title)))
	}
}

// Del ...
func Del() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		response, key := &model.Response{}, r.FormValue("key")
		if key == "" {
			w.Write(response.Build(-1, "miss query key", nil))
			return
		}

		// 文章不存在
		if !service.New().Exists(key) {
			w.Write(response.Build(-1, "not find key", nil))
			return
		}
		w.Write(response.Build(0, "success", service.New().Del(key)))
	}
}
func getPostData(r *http.Request) (*model.Paper, error) {
	var result = &model.Paper{}
	err := json.NewDecoder(r.Body).Decode(result)
	return result, err
}
