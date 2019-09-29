package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/share"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// SessionExists 验证session是否存在
func SessionExists(ctx *gin.Context) {
	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}
	ctx.JSON(200, model.ResponseSuccess(aSession))
}

// SessionLogin 登录验证
// 💀存在多次登录漏洞
func SessionLogin(ctx *gin.Context) {
	userName, password := ctx.Query("userName"), ctx.Query("passWord")
	if userName == "" || password == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	var user = sql.New().GetStorage(userName)
	if user.Name != userName || user.Password != share.PassWord(password) {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	sessionBody := &session.Body{
		Secret:     user.Name,
		Permission: user.Permission,
	}

	aSession := session.NewStore().Set(sessionBody)
	ctx.JSON(200, model.ResponseSuccess(aSession))
}
