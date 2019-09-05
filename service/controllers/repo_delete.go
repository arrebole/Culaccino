package controllers

import (
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// RepoDelete 删除内容api
func RepoDelete(ctx *gin.Context) {
	domain, symbol := ctx.Param("domain"), ctx.Param("symbol")

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

	sql.New().DelRepo(domain, symbol)
	ctx.JSON(200, module.ResponseSuccess())
	return
}
