package controllers

import (
	"strconv"

	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// SearchDashboard 目录索引api
func SearchDashboard(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "0")
	p, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	ctx.JSON(200, module.ResponseSuccess(sql.New().Dashboard(p, 5)))
}
