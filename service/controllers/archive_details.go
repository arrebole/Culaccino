package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchivesDetails 具体内容
func ArchivesDetails(ctx *gin.Context) {
	if id, err := middleware.Parsers(ctx).ParamsID(); err == nil {
		ctx.JSON(200, module.ResponseSuccess(sql.New().ArchiveQueryByID(id)))
		return
	}
	ctx.JSON(200, module.ResponseFail())
}
