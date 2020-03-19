package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// Papers 首页文章索引
func Papers() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(model.CreateResponse(0, "success", service.New().Table(10)).ToBytes())
		return
	}
}
