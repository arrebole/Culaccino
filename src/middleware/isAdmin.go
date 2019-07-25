package middleware

import (
	"github.com/arrebole/culaccino/src/module"
	"github.com/gin-gonic/gin"
)

// IsAdmin 中间件
func IsAdmin(ctx *gin.Context) {
	token, err := ctx.Cookie("user_session")
	if err != nil || token != module.NewToken() {
		ctx.JSON(200, module.Fail())
		ctx.Abort()
		return
	}
}
