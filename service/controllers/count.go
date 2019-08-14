package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Count 获取文章总数目
func Count(ctx *gin.Context) {
	count := sql.New().Count()
	ctx.JSON(200, module.Count(count))
}
