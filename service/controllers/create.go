package controllers

import (
	"time"

	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/model"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

// repoNew 添加内容api
func NewRepo(ctx *gin.Context) {
	data := model.Repo{}
	if err := ctx.BindJSON(&data); err != nil || data.Name == "" {
		ctx.JSON(200, model.ResponseFail("格式错误"))
		return
	}

	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}

	data.Storage = aSession.Secret
	if sql.New().ExistsRepo(data.Symbol()) {
		ctx.JSON(200, model.ResponseFail("已存在"))
		return
	}

	sql.New().SetRepo(&data)
	ctx.JSON(200, model.ResponseSuccess())
}

func NewChapter(ctx *gin.Context) {
	data := model.Chapter{}
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(200, model.ResponseFail("格式错误"))
		return
	}

	if data.Name == "" || data.Repo == "" {
		ctx.JSON(200, model.ResponseFail("信息不完全"))
		return
	}

	aSession, err := middleware.Session(ctx)
	if err != nil {
		ctx.JSON(200, model.ResponseFail("没有登录"))
		return
	}

	data.Storage = aSession.Secret
	if sql.New().ExistsChapter(data.Symbol()) {
		ctx.JSON(200, model.ResponseFail("已存在"))
		return
	}

	data.CreateAt = time.Now()
	data.UpdateAt = data.CreateAt

	sql.New().SetChapter(&data)
	ctx.JSON(200, model.ResponseSuccess())
}
