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
	domain, repo := ctx.Query("domain"), ctx.Query("repo")

	cookie, err := ctx.Cookie("user_session")
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	aSession, err := session.New().Get(cookie)
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	data := sql.New().GetRepo(domain, repo)
	if data.Author == "" || aSession.UID != data.AuthorID {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.New().DelRepo(domain, repo)
	ctx.JSON(200, module.ResponseSuccess())
	return
}
