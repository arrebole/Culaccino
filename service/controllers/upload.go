package controllers

import (
	"fmt"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Upload 创建一个新仓库
func Upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		staticUpload(ctx)
	}
}

// StaticUpload 处理静态文件
func staticUpload(ctx *gin.Context) {
	file, err := middleware.Parsers(ctx).FileFromBody()
	if err != nil {
		ctx.JSON(400, module.ResponseFail())
		return
	}

	file.SetRoot(config.UploadDir)
	err = file.SaveFile()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, module.ResponseFail())
		return
	}
	ctx.JSON(200, module.ResponseSuccess(file))
}
