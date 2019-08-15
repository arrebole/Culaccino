package controllers

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/arrebole/culaccino/service/config"
	"github.com/arrebole/culaccino/service/module"
	"github.com/arrebole/culaccino/service/sql"
	"github.com/gin-gonic/gin"
)

func dbFilePath() string {
	path, _ := filepath.Abs(config.Cofig.SQL.DBName)
	return path
}

func dbPath() string {
	path, _ := filepath.Abs("./data")
	return path
}

func dbName() string {
	pathArray := strings.Split(config.Cofig.SQL.DBName, "/")
	return pathArray[len(pathArray)-1]
}

// DBDownLoad ...
func DBDownLoad(ctx *gin.Context) {
	ctx.FileAttachment(dbFilePath(), dbName())
}

// DBUpload ...
func DBUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(200, module.Fail())
	}
	// 重新连接
	sql.New().Close()
	defer sql.ConnSQL()

	fileName := path.Join(dbPath(), file.Filename)
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		ctx.JSON(200, module.Fail())
	}

	ctx.JSON(200, module.Success())
}
