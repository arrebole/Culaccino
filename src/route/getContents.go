package route

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/sql"
	"github.com/arrebole/culaccino/src/util"
	"github.com/gin-gonic/gin"
)

// Contents 目录索引api
func Contents(ctx *gin.Context) {
	if id, err := util.ParesID(ctx); err == nil {
		data := sql.New().Query(id)
		ctx.JSON(200, module.RepData(data, nil))
		return
	}
	ctx.JSON(200, module.Fail())
}
