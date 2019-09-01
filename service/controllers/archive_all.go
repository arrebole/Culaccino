package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveAll 所有目录
func ArchiveAll(ctx *gin.Context) {
	total := sql.New().ArchiveCount().Total
	ctx.JSON(200, module.ResponseData(sql.New().ArchiveQueryDir(0, total)))
}
