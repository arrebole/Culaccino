package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// RepoNew 添加内容api
func RepoNew(ctx *gin.Context) {
	article, err := middleware.Parsers(ctx).BodyArchive()
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

	sql.New().NewRepo(article, &aSession)
	ctx.JSON(200, module.ResponseSuccess())
}