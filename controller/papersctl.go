package controller

import (
	"encoding/json"
	"net/http"

	"github.com/arrebole/culaccino/controller/transfer"
	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// Papers 首页文章索引
func Papers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		papersRead(w, r)
	case "POST":
		papersCreate(w, r)
	}
}

// papersReader( 读取papers列表
func papersRead(w http.ResponseWriter, r *http.Request) {
	transfer.Send(w, model.Papers(service.New().Table(10)))
}

// papersCreator 创建一篇文章
func papersCreate(w http.ResponseWriter, r *http.Request) {

	// 文章已存在
	id := urlLastParser(r.RequestURI)
	if service.New().Exists(id) {
		transfer.Send(w, model.Conflict)
		return
	}

	// 处理post提交的数据
	receive, err := bodyparser(r)
	if err != nil {
		transfer.Send(w, model.UnsupportedMediaType)
		return
	}

	// 数据库写入错误
	paper, err := service.New().Create(receive)
	if err != nil {
		transfer.Send(w, model.DBFail)
		return
	}
	transfer.Send(w, model.CreateSucesss(paper))
}

// bodyparser 解析post内容
func bodyparser(r *http.Request) (*model.Paper, error) {
	var result = &model.Paper{}
	err := json.NewDecoder(r.Body).Decode(result)
	return result, err
}
