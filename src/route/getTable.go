package route

import (
	"github.com/arrebole/culaccino/src/sql"
	"github.com/gin-gonic/gin"
)

var db = sql.New()

// Table 目录索引api
func Table(ctx *gin.Context) {
	ctx.String(200, "hello world")
}
