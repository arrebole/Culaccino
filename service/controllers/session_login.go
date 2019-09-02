package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// SessionLogin 登录验证
func SessionLogin(ctx *gin.Context) {
	userName, password := middleware.Parsers(ctx).QueryUser()
	if userName == "" || password == "" {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	user := sql.New().UserBaseInfo(userName, password)
	if user == nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}
	ctx.JSON(200, module.ResponseSuccess(session.New().Set(user)))
}
