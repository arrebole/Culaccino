package controllers

import (
	"fmt"
	"strconv"

	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)


// GetRepo 具体内容
func GetRepo(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")
	if storage == "" || repo == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	result := sql.New().GetRepo(fmt.Sprintf("%s:%s", storage, repo))
	ctx.JSON(200, model.ResponseSuccess(result))
}

// GetChapter ...
func GetChapter(ctx *gin.Context) {
	var (
		storage = ctx.Query("storage")
		repo    = ctx.Query("repo")
		chapter = ctx.Query("chapter")
	)
	if storage == "" || repo == "" || chapter == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}
	result := sql.New().GetChapter(fmt.Sprintf("%s:%s:%s", storage, repo, chapter))
	ctx.JSON(200, model.ResponseSuccess(result))
}

// GetStorage 所有者所有用户信息
func GetStorage(ctx *gin.Context) {
	storage := ctx.Query("storage")
	if storage == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	result := sql.New().GetStorage(storage)
	ctx.JSON(200, model.ResponseSuccess(result))
}

// GetDashboard 目录索引api
func GetDashboard(ctx *gin.Context) {
	var (
		pageString = ctx.DefaultQuery("page", "0")
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

// GetReposOfStorage ...
func GetReposOfStorage(ctx *gin.Context) {
	var storage = ctx.Query("storage")

	if storage == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}
	data := sql.New().GetRepos(storage)
	ctx.JSON(200, model.ResponseSuccess(data))
}
