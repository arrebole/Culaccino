package test

import (
	"strconv"
	"testing"

	"github.com/arrebole/culaccino/src/model"
	"github.com/arrebole/culaccino/src/service"
)

var store = service.CreateStore()

func createArticleModel(n int) []*model.Article {
	var result []*model.Article
	for i := 0; i < n; i++ {
		result = append(result, &model.Article{
			Name:    "test-" + strconv.Itoa(i),
			Cover:   "/img/test.png",
			Summary: "summary",
			Tag:     "",
		})
	}
	return result
}

func TestInsertArticle(T *testing.T) {
	var data = createArticleModel(100)
	for _, v := range data {
		store.CreateArticle(v)
	}
}
