package route

import (
	"fmt"

	"github.com/arreble/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Table 目录索引api
func Table(ctx *gin.Context) {
	if page, err := pareser.New(ctx).ParamsPage(); err == nil {
		const limit = 5
		fmt.Println(page)
		data := sql.New().QueryDir(limit, page)
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
