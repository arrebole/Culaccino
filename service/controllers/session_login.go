package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// SessionLogin 登录验证
func SessionLogin(ctx *gin.Context) {
	userName, password := pareser.New(ctx).QueryUser()
	if userName == "" || password == "" {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	authorized := sql.New().UserExist(userName, password)
	if !authorized {
		ctx.JSON(200, module.ResponseFail())
		return
	}
	ctx.JSON(200, module.ResponseData(module.NewSession()))
}
