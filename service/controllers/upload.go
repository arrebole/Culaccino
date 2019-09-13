package controllers

import (
	"fmt"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/model"
	"github.com/gin-gonic/gin"
)


// StaticUpload 处理静态文件
func StaticUpload(ctx *gin.Context) {
	file, err := middleware.FileFromBody(ctx)
	if err != nil {
		ctx.JSON(400, model.ResponseFail())
		return
	}

	file.SetRoot(config.UploadDir)
	status, err := file.SaveFile()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, model.ResponseFail())
		return
	}
	ctx.JSON(200, model.ResponseSuccess(status))
}
