package pareser

import (
	"errors"
	"strconv"

	"github.com/arrebole/culaccino/service/module"
	"github.com/gin-gonic/gin"
)

// Pareser 解析器
type Pareser struct {
	ctx *gin.Context
}

// New 解析器构造器
func New(ctx *gin.Context) *Pareser {
	return &Pareser{ctx: ctx}
}

// BodyArchive ...
func (p *Pareser) BodyArchive() (*module.PostArchive, error) {
	var article = &module.PostArchive{}
	if err := p.ctx.BindJSON(article); err != nil {
		return nil, errors.New("[Pareser]: pares post article fail")
	}
	return article, nil
}

// QueryToken 解析url中的token
func (p *Pareser) QueryToken() (string, error) {
	token := p.ctx.Query("token")
	if token == "" {
		return "", errors.New("[Pareser]: no token")
	}
	return token, nil
}

// ParamsID 解析参数中的id
func (p *Pareser) ParamsID() (uint, error) {
	id, err := strconv.ParseUint(p.ctx.Param("id"), 10, 0)
	if err != nil {
		return 0, errors.New("[Pareser]: params id fail")
	}
	return uint(id), nil
}

// ParamsPage ...
func (p *Pareser) ParamsPage() (int, error) {
	id, err := strconv.ParseInt(p.ctx.Param("page"), 10, 0)
	if err != nil {
		return 0, errors.New("[Pareser]: params page fail")
	}
	return int(id), nil
}

// QueryUser 解析login
func (p *Pareser) QueryUser() (name string, pwd string) {
	userName := p.ctx.Query("userName")
	password := p.ctx.Query("password")
	return userName, password
}
