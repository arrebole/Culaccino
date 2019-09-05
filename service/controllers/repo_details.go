package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// RepoDetails 具体内容
func RepoDetails(ctx *gin.Context) {
	domain, symbol := ctx.Param("domain"), ctx.Param("symbol")

	ctx.JSON(200, module.ResponseSuccess(sql.New().GetRepo(domain, symbol)))
}
