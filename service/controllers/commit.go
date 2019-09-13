package controllers

import (
	"time"

	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// RepoCommit 更新数据api
func CommitRepo(ctx *gin.Context) {
	data := model.Repo{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(200, model.ResponseFail())
		return
	}

	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}

	data.Storage = aSession.Secret
	oldData := sql.New().GetRepo(data.Symbol())
	if aSession.Secret != oldData.Storage {
		ctx.JSON(200, model.ResponseFail("名称不匹配"))
		return
	}

	sql.New().SetRepo(&data)
	ctx.JSON(200, model.ResponseSuccess())
}

func CommitChapter(ctx *gin.Context) {

	data := model.Chapter{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(200, model.ResponseFail("格式错误"))
		return
	}

	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}

	oldData := sql.New().GetChapter(data.Symbol())
	if aSession.Secret != oldData.Storage {
		ctx.JSON(200, model.ResponseFail("名称不匹配"))
		return
	}

	data.UpdateAt = time.Now()
	
	sql.New().SetChapter(&data)
	ctx.JSON(200, model.ResponseSuccess())
}
