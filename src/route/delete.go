package route

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/sql"
	"github.com/arrebole/culaccino/src/util"
	"github.com/gin-gonic/gin"
)

// Delete 删除内容api
func Delete(ctx *gin.Context) {
	if id, err := util.ParesID(ctx); err == nil {
		sql.New().Delete(id)
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
