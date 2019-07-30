package util

import (
	"errors"
	"strconv"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Pares 解析article json
func Pares(ctx *gin.Context) (*module.PostArticle, error) {
	var article = &module.PostArticle{}
	if err := ctx.BindJSON(article); err != nil {
		return nil, errors.New("pares article fail")
	}

	return article, nil
}

// ParesID 解析参数中的id
func ParesID(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 0)
	if err != nil {
		return 0, errors.New("paramsID fail")
	}
	return uint(id), nil
}

// QueryUser 解析login
func QueryUser(ctx *gin.Context) (name string, pwd string) {
	userName := ctx.Query("userName")
	password := ctx.Query("password")
	return userName, password
}
