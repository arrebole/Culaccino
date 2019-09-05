package middleware

import (
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Parser 解析器
type Parser struct {
	ctx *gin.Context
}

// Parsers 解析器构造器
func Parsers(ctx *gin.Context) *Parser {
	return &Parser{ctx: ctx}
}

// BodyArchive ...
func (p *Parser) BodyArchive() (*module.PostArchive, error) {
	var article = &module.PostArchive{}
	if err := p.ctx.BindJSON(article); err != nil {
		return nil, errors.New("[Pareser]: pares post article fail")
	}
	return article, nil
}

// FileFromBody ...
func (p *Parser) FileFromBody() (*module.File, error) {

	// 读取上传的文件
	fileHeader, err := p.ctx.FormFile("file")
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
