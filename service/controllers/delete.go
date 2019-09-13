package controllers

import (
	"fmt"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)


// DelRepo 删除内容api
func DelRepo(ctx *gin.Context) {
	storage, repo := ctx.Query("storage"), ctx.Query("repo")

	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}

	if aSession.Secret != storage {
		ctx.JSON(200, model.ResponseFail("权限不足"))
		return
	}

	sql.New().DelRepo(fmt.Sprintf("%s:%s", storage, repo))
	ctx.JSON(200, model.ResponseSuccess())
	return
}
