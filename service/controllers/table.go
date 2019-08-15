package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

const limit = 5

// Table 目录索引api
func Table(ctx *gin.Context) {
	if page, err := pareser.New(ctx).ParamsPage(); err == nil {
		data := sql.New().QueryDir(limit, page)
		count := sql.New().Count()
		ctx.JSON(200, module.RepData(data, count))
		return
	}

	ctx.JSON(200, module.Fail())
}

// TableAll 所有目录
func TableAll(ctx *gin.Context) {
	data := sql.New().QueryDirAll()
	count := sql.New().Count()
	ctx.JSON(200, module.RepData(data, count))
}
