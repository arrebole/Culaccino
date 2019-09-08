package middleware

import (
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// FileFromBody ...
func FileFromBody(ctx *gin.Context) (*module.File, error) {

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
