package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
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
	domain, repo := ctx.Query("domain"), ctx.Query("repo")
	if domain == "" || repo == "" {
		ctx.JSON(200, module.ResponseFail())
	}

	postArchive, err := middleware.Parsers(ctx).BodyArchive()
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

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

	oldData := sql.New().GetRepo(domain, repo)
	if oldData.Author == "" || aSession.UID != oldData.AuthorID {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.New().CommitRepo(domain, repo, module.ToArchive(postArchive))
	ctx.JSON(200, module.ResponseSuccess())
}
