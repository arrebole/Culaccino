package controllers

import (
	"fmt"

	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Search ...
func Search(ctx *gin.Context) {
	switch ctx.Query("type") {
	case "reop":
		SearchRepo(ctx)
	case "repos":
		SearchRepos(ctx)
	case "storage":
		SearchStorage(ctx)
	case "chapter":
		SearchChapter(ctx)
	default:
		ctx.JSON(200, model.ResponseFail())
	}
}

// SearchRepo 具体内容
func SearchRepo(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")
	if storage == "" || repo == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	result := sql.New().GetRepo(fmt.Sprintf("%s:%s", storage, repo))
	ctx.JSON(200, model.ResponseSuccess(result))
}

// SearchChapter ...
func SearchChapter(ctx *gin.Context) {
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

// SearchStorage 所有者所有用户信息
func SearchStorage(ctx *gin.Context) {
	storage := ctx.Query("storage")
	if storage == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	result := sql.New().GetStorage(storage)
	ctx.JSON(200, model.ResponseSuccess(result))
}

// SearchRepos 获取一个Storage中的所有repo
func SearchRepos(ctx *gin.Context) {
	var storage = ctx.Query("storage")

	if storage == "" {
		ctx.JSON(200, model.ResponseFail())
		return
	}
	data := sql.New().GetRepos(storage)
	ctx.JSON(200, model.ResponseSuccess(data))
}
