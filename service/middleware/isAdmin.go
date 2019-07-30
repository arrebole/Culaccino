package middleware

import (
	"os"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// IsAdmin 中间件
func IsAdmin(ctx *gin.Context) {
	token, err := ctx.Cookie("user_session")
	if os.Getenv("mode") != "dev" && (err != nil || token != module.NewToken()) {
		ctx.JSON(200, module.Fail())
		ctx.Abort()
		return
	}
}
