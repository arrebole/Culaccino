package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/arrebole/culaccino/service/util"
	"github.com/gin-gonic/gin"
)

// Update 更新数据api
func Update(ctx *gin.Context) {
	var (
		id      uint
		article *module.PostArticle
	)
	id, err1 := util.ParesID(ctx)
	article, err2 := util.Pares(ctx)

	if err1 == nil && err2 == nil {
		sql.New().Update(id, module.ToArticle(article))
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
