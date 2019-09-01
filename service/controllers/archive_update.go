package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveUpdate 更新数据api
func ArchiveUpdate(ctx *gin.Context) {
	var pares = pareser.New(ctx)

	id, err := pares.ParamsID()
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	postArchive, err := pares.BodyArchive()
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.New().ArchiveUpdate(id, module.ToArchive(postArchive))
	ctx.JSON(200, module.ResponseSuccess())
}
