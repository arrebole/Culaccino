package controllers

import (
	"fmt"

	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Export 输出内容
func Export() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		switch ctx.Query("tag") {
		case "dashboard":
			dashboard(ctx)
		case "storage":
			getStorage(ctx)
		case "repo":
			getRepo(ctx)
		case "chapter":
			getChapter(ctx)
		default:
			ctx.JSON(200, module.ResponseFail())
		}
	}
}

// repoDetails 具体内容
func getRepo(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")
	if storage == "" || repo == "" {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	var result = &module.Repo{}
	sql.Get(result, storage+":"+repo)
	ctx.JSON(200, module.ResponseSuccess(result))
}

func getChapter(ctx *gin.Context) {
	var (
		storage = ctx.Query("storage")
		repo    = ctx.Query("repo")
		chapter = ctx.Query("chapter")
	)
	if storage == "" || repo == "" || chapter == "" {
		ctx.JSON(200, module.ResponseFail())
		return
	}
	var result = &module.Chapter{}
	sql.Get(result, fmt.Sprintf("%s:%s:%s", storage, repo, chapter))
	ctx.JSON(200, module.ResponseSuccess(result))
}

// repoStorage 所有者所有
func getStorage(ctx *gin.Context) {
	cookie, err := ctx.Cookie("user_session")
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	aSession, err := session.NewStore().Get(cookie)
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	var result = &module.Storage{}
	sql.Get(result, aSession.Secret)

	ctx.JSON(200, module.ResponseSuccess(result))
}

// dashboard 目录索引api
func dashboard(ctx *gin.Context) {
	///page := ctx.DefaultQuery("page", "0")
	//p, err := strconv.Atoi(page)
	// if err != nil {
	// 	ctx.JSON(200, module.ResponseFail())
	// 	return
	// }

	var result []module.Repo
	for _, v := range sql.Explore() {
		var item = module.Repo{}
		sql.Get(item, v)
		result = append(result, item)
	}
	ctx.JSON(200, module.ResponseSuccess(result))
}
