package route

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/sql"
	"github.com/arrebole/culaccino/src/util"
	"github.com/gin-gonic/gin"
)

// Add 添加内容api
func Add(ctx *gin.Context) {
	if article, err := util.Pares(ctx); err == nil {
		sql.New().Add(article)
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
