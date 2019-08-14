package controllers

import (
	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/pareser"
	"github.com/gin-gonic/gin"
)

// Login 登录验证
func Login(ctx *gin.Context) {
	userName, password := pareser.New(ctx).QueryUser()
	for _, usr := range config.Cofig.Admin {
		if userName == usr.UserName && password == usr.Password {
			ctx.JSON(200, module.LoginOk())
			return
		}
	}
	ctx.JSON(200, module.Fail())
}
