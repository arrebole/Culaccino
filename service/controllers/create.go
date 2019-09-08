package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Create 创建一个新仓库
func Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repoNew(ctx)
	}
}

// repoNew 添加内容api
func repoNew(ctx *gin.Context) {
	repo := &module.Repo{}
	if err := ctx.BindJSON(repo); err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

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
	repo.Symbol = aSession.Secret + ":" + repo.Name
	sql.Set(repo)
	ctx.JSON(200, module.ResponseSuccess())
}
