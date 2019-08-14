package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Contents 具体内容
func Contents(ctx *gin.Context) {
	if id, err := pareser.New(ctx).ParamsID(); err == nil {
		data := sql.New().Query(id)
		ctx.JSON(200, module.RepData(data, nil))
		return
	}
	ctx.JSON(200, module.Fail())
}
