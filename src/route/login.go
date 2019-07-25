package route

import (
	"github.com/arrebole/culaccino/src/config"
	"github.com/arrebole/culaccino/src/module"
	"github.com/arrebole/culaccino/src/util"
	"github.com/gin-gonic/gin"
)

// Login 登录验证
func Login(ctx *gin.Context) {
	userName, password := util.QueryUser(ctx)
	for _, usr := range config.Cofig.Admin {
		if userName == usr.UserName && password == usr.Password {
			ctx.JSON(200, module.LoginOk())
			return
		}
	}
	ctx.JSON(200, module.Fail())
}
