package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/gin-gonic/gin"
)

// Token 验证cookie
func Token(ctx *gin.Context) {
	cookie, err := pareser.New(ctx).QueryToken()
	if err == nil && cookie == module.NewToken() {
		ctx.JSON(200, module.Power())
		return
	}
	ctx.JSON(200, module.Fail())
}
