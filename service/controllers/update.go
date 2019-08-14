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

	postArticle, err := pares.BodyArticle()
	if err != nil {
		ctx.JSON(200, module.Fail())
		return
	}

	sql.New().Update(id, module.ToArticle(postArticle))
	ctx.JSON(200, module.Success())
}
