package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Update 更新数据api
func Update(ctx *gin.Context) {
	var pares = pareser.New(ctx)

	id, err := pares.ParamsID()
	if err != nil {
		ctx.JSON(200, module.Fail())
		return
	}

	postArchive, err := pares.BodyArchive()
	if err != nil {
		ctx.JSON(200, module.Fail())
		return
	}

	sql.New().Update(id, module.ToArchive(postArchive))
	ctx.JSON(200, module.Success())
}
