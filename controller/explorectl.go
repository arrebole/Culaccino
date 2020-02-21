package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/model"
	"github.com/arrebole/culaccino/service"
)

// Explore 首页文章索引
func Explore() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		result := &model.Response{}
		w.Write(result.Build(0, "success", service.New().Explore(10)))
	}
}
