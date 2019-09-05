package controllers

import (
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/session"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// ArchiveDelete 删除内容api
func ArchiveDelete(ctx *gin.Context) {
	id, err := middleware.Parsers(ctx).ParamsID()
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

	sql.New().ArchiveDelete(id)
	ctx.JSON(200, module.ResponseSuccess())
	return
}
