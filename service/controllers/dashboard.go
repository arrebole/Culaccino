package controllers

import (
	"strconv"

	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// GetDashboard 获取首页甲板内容
func GetDashboard(ctx *gin.Context) {
	var (
		pageString    = ctx.DefaultQuery("page", "0")
		perpageString = ctx.DefaultQuery("per_page", "0")
	)

	page, err := strconv.ParseInt(pageString, 10, 64)
	if err != nil {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	perpage, err := strconv.ParseInt(perpageString, 10, 64)
	if err != nil {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	data := sql.New().GetRepos(sql.New().Explore(page, perpage)...)
	ctx.JSON(200, model.ResponseSuccess(data))
}
