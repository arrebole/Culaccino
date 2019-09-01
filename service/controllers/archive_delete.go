package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveDelete 删除内容api
func ArchiveDelete(ctx *gin.Context) {
	if id, err := pareser.New(ctx).ParamsID(); err == nil {
		sql.New().ArchiveDelete(id)
		ctx.JSON(200, module.ResponseSuccess())
		return
	}
	ctx.JSON(200, module.ResponseFail())
}
