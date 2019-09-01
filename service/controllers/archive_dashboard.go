package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

const perPage = 5

// ArchiveDashboard 目录索引api
func ArchiveDashboard(ctx *gin.Context) {
	if page, err := pareser.New(ctx).ParamsPage(); err == nil {
		ctx.JSON(200, module.ResponseData(sql.New().ArchiveQueryDir(page, perPage)))
		return
	}

	ctx.JSON(200, module.ResponseFail())
}
