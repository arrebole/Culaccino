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
		const limit = 5
		data := sql.New().QueryDir(limit, int(id))
		ctx.JSON(200, module.RepData(nil, data))
		return
	}

	ctx.JSON(200, module.Fail())
}

// TableAll 所有目录
func TableAll(ctx *gin.Context) {
	data := sql.New().QueryDirAll()
	ctx.JSON(200, module.RepData(nil, data))

}
