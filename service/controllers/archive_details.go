package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchivesDetails 具体内容
func ArchivesDetails(ctx *gin.Context) {
	if id, err := pareser.New(ctx).ParamsID(); err == nil {
		ctx.JSON(200, module.ResponseData(sql.New().Query(id)))
		return
	}
	ctx.JSON(200, module.ResponseFail())
}
