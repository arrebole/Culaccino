package controllers

import (
	"fmt"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/middleware"
	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// StaticUpload 处理静态文件
func StaticUpload(ctx *gin.Context) {
	file, err := middleware.Parsers(ctx).FileFromBody()
	if err != nil {
		fmt.Println(err.Error())
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
