package route

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/sql"
	"github.com/gin-gonic/gin"
)

func pares(ctx *gin.Context) *module.PostArticle {
	var article = &module.PostArticle{}
	if err := ctx.BindJSON(article); err != nil {
		return nil
	}
	return article
}

// Add 添加内容api
func Add(ctx *gin.Context) {
	if article := pares(ctx); article != nil {
		sql.New().Add(article)
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
