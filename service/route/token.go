package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arreble/culaccino/service/pareser"
	"github.com/gin-gonic/gin"
)

// Token 验证cookie
func Token(ctx *gin.Context) {
	cookie, err := pareser.New().QueryToken()
	if err == nil && cookie == module.NewToken() {
		ctx.JSON(200, module.Power())
		return
	}
	ctx.JSON(200, module.Fail())
}
