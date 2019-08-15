package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Add 添加内容api
func Add(ctx *gin.Context) {
	if article, err := pareser.New(ctx).BodyArchive(); err == nil {
		sql.New().Add(article)
		ctx.JSON(200, module.Success())
		return
	}
	ctx.JSON(200, module.Fail())
}
