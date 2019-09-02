package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/gin-gonic/gin"
)

// SessionExist 验证session是否存在
func SessionExist(ctx *gin.Context) {
	cookie, err := ctx.Cookie("user_session")
	if err != nil {
		ctx.JSON(200, module.ResponseFail(module.NoLogin))
		return
	}

	_, err = session.New().Get(cookie)
	if err != nil {
		ctx.JSON(200, module.ResponseFail(module.NoLogin))
		return
	}

	ctx.JSON(200, module.ResponseSuccess(module.HadLogin))
}
