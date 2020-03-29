package controller

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/arrebole/culaccino/controller/transfer"
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
	// 判断文章是否存在
	id := urlLastParser(r.RequestURI)
	if !service.New().Exists(id) {
		transfer.Send(w, model.Conflict)
		return
	}

	receive, err := bodyparser(r)
	if err != nil {
		transfer.Send(w, model.UnsupportedMediaType)
		return
	}

	// 数据库写入错误
	paper, err := service.New().Update(id, receive)
	if err != nil {
		transfer.Send(w, model.DBFail)
	}

	transfer.Send(w, model.OnePaper(paper))
}

// Get 查看文章详细内容
func paperRead(w http.ResponseWriter, r *http.Request) {
	id := urlLastParser(r.RequestURI)
	if id == "" {
		transfer.Send(w, model.NotFound)
		return
	}

	// 文章不存在
	if !service.New().Exists(id) {
		transfer.Send(w, model.NotFound)
		return
	}
	transfer.Send(w, model.OnePaper(service.New().Get(id)))

}

// Del ...
func paperDel(w http.ResponseWriter, r *http.Request) {
	id := urlLastParser(r.RequestURI)
	if id == "" {
		transfer.Send(w, model.NotFound)
		return
	}
	// 文章不存在
	if !service.New().Exists(id) {
		transfer.Send(w, model.NotFound)
		return
	}
	if err := service.New().Del(id); err != nil {
		transfer.Send(w, model.DBFail)
		return
	}

	transfer.Send(w, model.Sucess())

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
