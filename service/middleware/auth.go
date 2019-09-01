package middleware

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Auth 认证中间件
func Auth(ctx *gin.Context) {
	token, err := ctx.Cookie("user_session")
	if err != nil || token != module.NewSession().Token {
		ctx.JSON(200, module.ResponseFail())
		ctx.Abort()
		return
	}
}
