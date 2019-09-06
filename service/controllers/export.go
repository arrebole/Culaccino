package controllers

import (
	"strconv"

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
			repoStorage(ctx)
		case "repo":
			repoDetails(ctx)
		default:
			ctx.JSON(200, module.ResponseFail())
		}
	}
}

// repoDetails 具体内容
func repoDetails(ctx *gin.Context) {
	domain, repo := ctx.Query("storage"), ctx.Query("repo")
	ctx.JSON(200, module.ResponseSuccess(sql.New().GetRepo(domain, repo)))
}

// repoStorage 所有者所有
func repoStorage(ctx *gin.Context) {
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

	ctx.JSON(200, module.ResponseSuccess(sql.New().GetRepos(aSession.Secret)))
}

// dashboard 目录索引api
func dashboard(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "0")
	p, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	ctx.JSON(200, module.ResponseSuccess(sql.New().Dashboard(p, 5)))
}
