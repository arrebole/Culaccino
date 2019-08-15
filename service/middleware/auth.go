package middleware

import (
	"os"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Auth 认证中间件
func Auth(ctx *gin.Context) {
	token, err := ctx.Cookie("user_session")
	if os.Getenv("mode") != "debug" && (err != nil || token != module.NewToken()) {
		ctx.JSON(200, module.Fail())
		ctx.Abort()
		return
	}
}
