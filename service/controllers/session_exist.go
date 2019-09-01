package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/gin-gonic/gin"
)

// SessionExist 验证session是否存在
func SessionExist(ctx *gin.Context) {
	cookie, err := pareser.New(ctx).QueryToken()
	if err == nil && cookie == module.NewSession().Token {
		ctx.JSON(200, module.ResponseData(module.HadLogin))
		return
	}
	ctx.JSON(200, module.ResponseData(module.NoLogin))
}
