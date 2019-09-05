package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveUpdate 更新数据api
func ArchiveUpdate(ctx *gin.Context) {
	var pares = middleware.Parsers(ctx)

	id, err := pares.ParamsID()
	if err != nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	postArchive, err := pares.BodyArchive()
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

	archive := sql.New().ArchiveQueryByID(id)
	if archive == nil {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	if aSession.UID != archive.AuthorID {
		ctx.JSON(200, module.ResponseFail())
		return
	}

	sql.New().ArchiveUpdate(id, module.ToArchive(postArchive))
	ctx.JSON(200, module.ResponseSuccess())
}
