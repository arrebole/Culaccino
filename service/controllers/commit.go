package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Commit 提交内容
func Commit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repoCommit(ctx)
	}
}

// RepoCommit 更新数据api
func repoCommit(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")
	if storage == "" || repo == "" {
		ctx.JSON(200, module.ResponseFail())
	}

	data := module.Repo{}
	if err := ctx.BindJSON(&data); err != nil {
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

	oldData := module.Repo{}
	sql.Get(&oldData, storage+":"+repo)
	if aSession.Secret != oldData.Parents() {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.Set(&data)
	ctx.JSON(200, module.ResponseSuccess())
}
