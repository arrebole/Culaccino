package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/arrebole/culaccino/service/util"
	"github.com/gin-gonic/gin"
)

// Update 更新数据api
func Update(ctx *gin.Context) {
	if article, err := util.Pares(ctx); err == nil {
		sql.New().Update(module.ToArticle(article))
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
