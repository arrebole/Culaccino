package controller

import (
	"net/http"

	"github.com/arrebole/culaccino/src/service"
)

// GetArticle ...
func GetArticle(w http.ResponseWriter, r *http.Request) {

	// articleName 是从url中获取的params参数
	// 例如 articleName = f("http://xxx.com/api/articles/refact") -> refact
	articleName := Params(r.Context(), "article")

	// 数据库仓库中通过article的名称查询数据
	// 构建响应的数据包
	resopnse := service.ArticlesRepo.Find(articleName)

	// 返回响应
	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(resopnse))
}

// UpdateArticle ...
func UpdateArticle(w http.ResponseWriter, r *http.Request) {

	// 将 post 请求的的json数据解析为对象
	payload, err := BodyPaserArticle(r.Body)
	if err != nil {
		ErrorHandle(w, r)
		return
	}

	// 将URL中的article名设置为需要更新的article名，限制名称
	// 阻止用户提交的json中带有Name名，跨名称更新
	payload.Name = Params(r.Context(), "article")

	// 将提交的数据更新到对应数据库中的 article
	resopnse := service.ArticlesRepo.Update(payload)

	// 返回响应
	w.Header().Add("Content-type", "application/json")
	w.Write(JSON(resopnse))
}
