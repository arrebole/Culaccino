package controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/arrebole/culaccino/src/model"
	"github.com/julienschmidt/httprouter"
)

// URLQueryInt 从URL中的query获取数据
// f("/api/article?limit=100", "limit", 0) -> 100
func URLQueryInt(url *url.URL, key string, defaul int) int {
	result, err := strconv.Atoi(url.Query().Get(key))
	if err != nil {
		return defaul
	}
	return result
}

// JSON 序列化实例对象为json bytes
func JSON(data interface{}) []byte {
	result, _ := json.Marshal(data)
	return result
}

// BodyPaserArticle ...
func BodyPaserArticle(body io.ReadCloser) (*model.Article, error) {
	var result = &model.Article{}
	decoder := json.NewDecoder(body)
	return result, decoder.Decode(result)
}

// ErrorHandle 处理错误的返回
func ErrorHandle(w http.ResponseWriter, r *http.Request) {
	data := model.Error{
		Code:  -1,
		Error: "Fail",
	}
	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(data))
}

// Params 从url中解析parsms
func Params(ctx context.Context, key string) string {
	return httprouter.ParamsFromContext(ctx).ByName(key)
}
