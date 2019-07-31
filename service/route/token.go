package route

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/util"
	"github.com/gin-gonic/gin"
)

// Token 验证cookie
func Token(ctx *gin.Context) {
	cookie, err := util.ParesToken(ctx)
	if err == nil && cookie == module.NewToken() {
		ctx.JSON(200, module.Power())
		return
	}
	ctx.JSON(200, module.Fail())
}
