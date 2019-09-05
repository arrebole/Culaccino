package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// RepoCommit 更新数据api
func RepoCommit(ctx *gin.Context) {
	domain, symbol := ctx.Param("domain"), ctx.Param("symbol")

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

	repo := sql.New().GetRepo(domain, symbol)
	if repo.Author == "" || aSession.UID != repo.AuthorID {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.New().CommitRepo(domain, symbol, module.ToArchive(postArchive))
	ctx.JSON(200, module.ResponseSuccess())
}
