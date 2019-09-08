package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// Delete 删除仓库
func Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		repoDelete(ctx)
	}
}

// RepoDelete 删除内容api
func repoDelete(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")

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

	var result = module.Repo{}
	sql.Get(&result, storage+":"+repo)
	if aSession.Secret != result.Parents() {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.Delete(&result)
	ctx.JSON(200, module.ResponseSuccess())
	return
}
