package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Session 处理session相关api
func Session() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Query("tag") {
		case "exist":
			sessionExist(ctx)
		case "login":
			sessionLogin(ctx)
		default:
			ctx.JSON(200, module.ResponseFail())
		}
	}
}

// SessionExist 验证session是否存在
func sessionExist(ctx *gin.Context) {
	cookie, err := ctx.Cookie("user_session")
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	aSession, err := session.NewStore().Get(cookie)
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}
	ctx.JSON(200, module.ResponseSuccess(aSession))
}

// SessionLogin 登录验证
func sessionLogin(ctx *gin.Context) {
	userName, password := ctx.Query("userName"), ctx.Query("passWord")
	if userName == "" || password == "" {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	var user = &module.Storage{}
	sql.Get(user, userName)
	if user.Name != userName {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	aSession := session.NewStore().Set(&session.Body{Secret: userName})
	ctx.JSON(200, module.ResponseSuccess(aSession))
}
