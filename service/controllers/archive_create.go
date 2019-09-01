package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveCreate 添加内容api
func ArchiveCreate(ctx *gin.Context) {
	if article, err := pareser.New(ctx).BodyArchive(); err == nil {
		sql.New().ArchiveCreate(article)
		ctx.JSON(200, module.ResponseSuccess())
		return
	}
	ctx.JSON(200, module.ResponseFail())
}
