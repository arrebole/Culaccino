package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// StaticUpload 处理静态文件
func StaticUpload(ctx *gin.Context) {
	file, err := getFileFromBody(ctx)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, module.ResponseFail())
		return
	}

	file.SetRoot("./wwwroot/static")
	err = file.SaveFile()
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(400, module.ResponseFail())
		return
	}
	ctx.JSON(200, module.ResponseData(file))
}

func getFileFromBody(ctx *gin.Context) (*module.File, error) {

	// 读取上传的文件
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return nil, errors.New("FormFile() fail")
	}

	//匹配后缀名
	reg := regexp.MustCompile(`\.(\w+)$`)
	fileSuffix := reg.FindString(fileHeader.Filename)

	//打开文件
	file, err := fileHeader.Open()
	if err != nil {
		return nil, errors.New("open file fail")
	}

	// 将文件读取到应用层
	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("read file fail")
	}
	//重命名文件名称
	result := &module.File{
		Suffix: fileSuffix,
		Data:   fileByte,
	}
	return result, nil
}
