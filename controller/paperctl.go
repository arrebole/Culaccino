package controller

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// Paper 处理单个paper请求
func Paper(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		paperRead(w, r)
	case "PATCH":
		paperUpdate(w, r)
	case "DELETE":
		paperDel(w, r)
	}
}

// Update 更新一篇文章
func paperUpdate(w http.ResponseWriter, r *http.Request) {
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

// Get 查看文章详细内容
func paperRead(w http.ResponseWriter, r *http.Request) {
	title := urlLastParser(r.RequestURI)
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

// Del ...
func paperDel(w http.ResponseWriter, r *http.Request) {
	key := urlLastParser(r.RequestURI)
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

// urlLastParser /api/papers/abc -> abc
func urlLastParser(requestURL string) string {
	decodeurl, err := url.QueryUnescape(requestURL)
	if err != nil {
		return ""
	}

	params := strings.Split(decodeurl, "/")
	if len(params) <= 0 {
		return ""
	}
	return params[len(params)-1]
}
