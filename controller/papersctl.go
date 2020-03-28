package controller

import (
	"encoding/json"
	"net/http"

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
	w.Write(model.CreateResponse(0, "success", service.New().Table(10)).ToBytes())
}

// papersCreator 创建一篇文章
func papersCreate(w http.ResponseWriter, r *http.Request) {

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

// bodyparser 解析post内容
func bodyparser(r *http.Request) (*model.Paper, error) {
	var result = &model.Paper{}
	err := json.NewDecoder(r.Body).Decode(result)
	return result, err
}
