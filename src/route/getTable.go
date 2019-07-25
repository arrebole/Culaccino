package route

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/sql"
	"github.com/arrebole/culaccino/src/util"
	"github.com/gin-gonic/gin"
)

var db = sql.New()

// Table 目录索引api
func Table(ctx *gin.Context) {
	if id, err := util.ParesID(ctx); err == nil {
		data := sql.New().QueryDir(id)
		ctx.JSON(200, module.RepData(nil, data))
		return
	}

	ctx.JSON(200, module.Fail())
}
