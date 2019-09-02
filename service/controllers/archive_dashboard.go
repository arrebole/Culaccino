package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveDashboard 目录索引api
func ArchiveDashboard(ctx *gin.Context) {
	if page, err := middleware.Parsers(ctx).ParamsPage(); err == nil {
		ctx.JSON(200, module.ResponseSuccess(sql.New().ArchiveDir(page, 5)))
		return
	}

	ctx.JSON(200, module.ResponseFail())
}
