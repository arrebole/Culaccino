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

		// 处理post提交的数据
		paper, err := bodyparser(r)
		if err != nil {
			w.Write(model.CreateResponse(-1, "fail post", nil).ToBytes())
			return
		}

		// 文章已存在
		if service.New().Exists(paper.Title) {
			w.Write(model.CreateResponse(-1, "paper Exists", nil).ToBytes())
			return
		}
		// 数据库写入错误
		if err = service.New().Set(paper); err != nil {
			w.Write(model.CreateResponse(-1, err.Error(), nil).ToBytes())
			return
		}
		w.Write(model.CreateResponse(0, "success", nil).ToBytes())
	}
}

// Update 更新一篇文章
func Update() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// 处理post提交的数据
		paper, err := bodyparser(r)
		if err != nil {
			w.Write(model.CreateResponse(-1, "fail post", nil).ToBytes())
			return
		}
		// 文章不存在
		if !service.New().Exists(paper.Title) {
			w.Write(model.CreateResponse(-1, "not find", nil).ToBytes())
			return
		}
		// 数据库写入错误
		if err = service.New().Set(paper); err != nil {
			w.Write(model.CreateResponse(-1, err.Error(), nil).ToBytes())
			return
		}
		w.Write(model.CreateResponse(0, "success", nil).ToBytes())
	}
}

// Get 查看文章详细内容
func Get() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		title := r.FormValue("key")
		if title == "" {
			w.Write(model.CreateResponse(-1, "miss query key", nil).ToBytes())
			return
		}

		// 文章不存在
		if !service.New().Exists(title) {
			w.Write(model.CreateResponse(-1, "not find key", nil).ToBytes())
			return
		}
		w.Write(model.CreateResponse(0, "success", service.New().Get(title)).ToBytes())
	}
}

// Del ...
func Del() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		key := r.FormValue("key")
		if key == "" {
			w.Write(model.CreateResponse(-1, "miss query key", nil).ToBytes())
			return
		}

		// 文章不存在
		if !service.New().Exists(key) {
			w.Write(model.CreateResponse(-1, "not find key", nil).ToBytes())
			return
		}
		w.Write(model.CreateResponse(0, "success", service.New().Del(key)).ToBytes())
	}
}

// bodyparser 解析post内容
func bodyparser(r *http.Request) (*model.Paper, error) {
	var result = &model.Paper{}
	err := json.NewDecoder(r.Body).Decode(result)
	return result, err
}
