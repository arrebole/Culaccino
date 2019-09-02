package middleware

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"

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

// QueryToken 解析url中的token
func (p *Parser) QueryToken() (string, error) {
	token := p.ctx.Query("token")
	if token == "" {
		return "", errors.New("[Pareser]: no token")
	}
	return token, nil
}

// ParamsID 解析参数中的id
func (p *Parser) ParamsID() (uint, error) {
	id, err := strconv.ParseUint(p.ctx.Param("id"), 10, 0)
	if err != nil {
		return 0, errors.New("[Pareser]: params id fail")
	}
	return uint(id), nil
}

// ParamsPage ...
func (p *Parser) ParamsPage() (int, error) {
	id, err := strconv.ParseInt(p.ctx.Param("page"), 10, 0)
	if err != nil {
		return 0, errors.New("[Pareser]: params page fail")
	}
	return int(id), nil
}

// QueryUser 解析login
func (p *Parser) QueryUser() (name string, pwd string) {
	userName := p.ctx.Query("userName")
	password := p.ctx.Query("password")
	return userName, password
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
