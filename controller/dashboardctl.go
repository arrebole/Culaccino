package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// Dashboard 首页文章索引
func Dashboard() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		result := &model.Response{}
		w.Write(result.Build(0, "success", service.New().Dashboard(10)))
	}
}
