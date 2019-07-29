package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/arrebole/culaccino/service/util"
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
